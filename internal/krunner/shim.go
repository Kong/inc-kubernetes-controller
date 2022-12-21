package krunner

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"go.uber.org/zap"

	"github.com/kong/inc-kubernetes-controller/internal/koko/config"
	"github.com/kong/inc-kubernetes-controller/internal/koko/db"
	genJSONSchema "github.com/kong/inc-kubernetes-controller/internal/koko/gen/jsonschema"
	"github.com/kong/inc-kubernetes-controller/internal/koko/model/json/schema"
	"github.com/kong/inc-kubernetes-controller/internal/koko/persistence"
	"github.com/kong/inc-kubernetes-controller/internal/koko/plugin"
	"github.com/kong/inc-kubernetes-controller/internal/koko/plugin/validators"
	serverUtil "github.com/kong/inc-kubernetes-controller/internal/koko/server/util"
	"github.com/kong/inc-kubernetes-controller/internal/koko/store"
)

type DPAuthMode int

const (
	DPAuthSharedMTLS = iota
	DPAuthPKIMTLS
)

var (
	demoCert = []byte(`-----BEGIN CERTIFICATE-----
MIIBWTCB3wIBADAKBggqhkjOPQQDAjAaMRgwFgYDVQQDDA9rb25nX2NsdXN0ZXJp
bmcwHhcNMjIxMjIxMjMwMTQ1WhcNMjUxMjIwMjMwMTQ1WjAaMRgwFgYDVQQDDA9r
b25nX2NsdXN0ZXJpbmcwdjAQBgcqhkjOPQIBBgUrgQQAIgNiAAS1bNY1+AXwval0
OWpg4b9TnhVJDH6XCs+ISJknt0AZVVGgMh3oIsoD7r8G7YWBn4sop0xPcg7gTkzP
m85yDePSDi13VaACSEFMskCwYZOST8sWXZgJMIIXn+4INTtKizcwCgYIKoZIzj0E
AwIDaQAwZgIxAMKIOu9IWTSRrGZzrh6CClWeMzIbWkqNvY/+lrrThtwqsL/ZnCOG
zIQnVqg4bzTXiwIxAPQt8brYcaKZblGaRHv6RjPJehSdl3esK2wAQZyLpm9xJksK
/0kxKw/bNl6q68ABcg==
-----END CERTIFICATE-----`)
	demoKey = []byte(`-----BEGIN PRIVATE KEY-----
MIG2AgEAMBAGByqGSM49AgEGBSuBBAAiBIGeMIGbAgEBBDAW9BT3Pv/SL/j/3qyY
/MTo67S60CDDiq4HG9PSMlz/exh6/e5eBoVatamcI0RRZ9OhZANiAAS1bNY1+AXw
val0OWpg4b9TnhVJDH6XCs+ISJknt0AZVVGgMh3oIsoD7r8G7YWBn4sop0xPcg7g
TkzPm85yDePSDi13VaACSEFMskCwYZOST8sWXZgJMIIXn+4INTtKizc=
-----END PRIVATE KEY-----`)
)

type ServerConfig struct {
	DPAuthMode    DPAuthMode
	DPAuthCert    tls.Certificate
	DPAuthCACerts []*x509.Certificate

	KongCPCert tls.Certificate

	Logger                  *zap.Logger
	Metrics                 config.Metrics
	Database                config.Database
	DisableAnonymousReports bool
}

type KokoRunnable struct {
	config      ServerConfig
	store       *store.ObjectStore
	storeLoader serverUtil.StoreLoader
	validator   *validators.LuaValidator
}

func InitKokoRunnable(logger *zap.Logger) (*KokoRunnable, error) {
	var k KokoRunnable
	var err error
	k.config, err = getServerConfig(logger)
	if err != nil {
		return nil, err
	}
	k.store, k.storeLoader, err = getStore(&k.config)
	if err != nil {
		return nil, err
	}
	k.validator, err = validators.NewLuaValidator(validators.Opts{
		Logger:      logger,
		StoreLoader: k.storeLoader,
	})
	if err != nil {
		return nil, err
	}
	err = k.validator.LoadSchemasFromEmbed(plugin.Schemas, "schemas")
	if err != nil {
		return nil, err
	}
	return &k, err
}

func (k *KokoRunnable) GetStore() *store.ObjectStore {
	return k.store
}

func (k *KokoRunnable) GetStoreLoader() serverUtil.StoreLoader {
	return k.storeLoader
}

func (k *KokoRunnable) GetValidator() *validators.LuaValidator {
	return k.validator
}

func (k *KokoRunnable) Start(ctx context.Context) error {
	if k.store == nil {
		return fmt.Errorf("KokoRunnable is not initialized, cannot start")
	}
	return k.Run(ctx)
}

func (k *KokoRunnable) NeedLeaderElection() bool {
	return true
}

// getServerConfig returns a static Koko configuration
func getServerConfig(logger *zap.Logger) (ServerConfig, error) {
	cert, err := tls.X509KeyPair(demoCert, demoKey)
	if err != nil {
		return ServerConfig{}, err
	}
	serverConfig := ServerConfig{
		DPAuthMode: DPAuthSharedMTLS,
		DPAuthCert: cert,
		KongCPCert: cert,
		Logger:     logger,
		Metrics: config.Metrics{
			ClientType: "noop",
			Prometheus: config.PrometheusMetrics{
				Enable: false,
			},
		},
		Database: config.Database{
			Dialect:      "sqlite3",
			QueryTimeout: "5s",
			SQLite: config.SQLite{
				InMemory: true,
			},
		},
		DisableAnonymousReports: true,
	}
	return serverConfig, nil
}

func getStore(config *ServerConfig) (*store.ObjectStore, *serverUtil.DefaultStoreLoader, error) {
	schema.RegisterSchemasFromFS(&genJSONSchema.KongSchemas)
	persister, err := setupDB(config.Logger, config.Database)
	if err != nil {
		return nil, nil, err
	}

	store := store.New(persister, config.Logger.With(zap.String("component", "store"))).ForCluster(store.DefaultCluster)
	storeLoader := serverUtil.DefaultStoreLoader{Store: store}

	return store, &storeLoader, nil
}

func setupDB(logger *zap.Logger, configDB config.Database) (persistence.Persister, error) {
	config, err := config.ToDBConfig(configDB, logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	config.Logger = logger
	m, err := db.NewMigrator(config)
	if err != nil {
		return nil, err
	}
	c, l, err := m.Status()
	if err != nil {
		return nil, err
	}
	logger.Sugar().Debugf("migration status: current: %d, latest: %d", c, l)

	if c != l {
		if configDB.Dialect == db.DialectSQLite3 {
			logger.Sugar().Info("migration out of date")
			logger.Sugar().Info("running migration in-process as sqlite" +
				" database detected")
			err := runMigrations(m)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("database schema out of date, " +
				"please run 'koko db migrate-up' to migrate the schema to" +
				" latest version")
		}
	}

	// setup data store
	return db.NewPersister(config)
}

func runMigrations(m *db.Migrator) error {
	if err := m.Up(); err != nil {
		return fmt.Errorf("migrating database: %v", err)
	}
	return nil
}
