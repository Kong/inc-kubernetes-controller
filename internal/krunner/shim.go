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
MIICejCCAiCgAwIBAgIUM/uXwOFH6RmU8DZMj85TerHjrQgwCgYIKoZIzj0EAwIw
gYQxCzAJBgNVBAYTAlVTMRQwEgYDVQQHEwtNaW5uZWFwb2xpczEUMBIGA1UEChML
WWFrIFNoYXZpbmcxDDAKBgNVBAsTA1dvbzE7MDkGA1UEAxMyWWFrIFNoYXZlcyBG
YWtlIEludGVybWVkaWF0ZSBDZXJ0aWZpY2F0ZSBBdXRob3JpdHkwHhcNMTgwMzE1
MjIwNDAwWhcNNDgwMzA3MjIwNDAwWjCBjDELMAkGA1UEBhMCVVMxFDASBgNVBAcT
C01pbm5lYXBvbGlzMRQwEgYDVQQKEwtZYWsgU2hhdmluZzEMMAoGA1UECxMDV29v
MUMwQQYDVQQDEzpZYWsgU2hhdmVzIEZha2UgSW50ZXJtZWRpYXRlIENlcnRpZmlj
YXRlIEF1dGhvcml0eSBMZXZlbCAyMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE
tBwGcExavS1hi32ovoU/VUtbQw9Ah15ypMOsUCKkvbz7waPCpRMtkiO7TqFH9gQD
ZYUz+RKvzh2XGTvvwiEBGqNmMGQwDgYDVR0PAQH/BAQDAgGGMBIGA1UdEwEB/wQI
MAYBAf8CAQQwHQYDVR0OBBYEFIo/bUwbPmcFeJKLx/yYcgFaFP/VMB8GA1UdIwQY
MBaAFNjbCjYBxylShp5O8FCPhx4MxrqUMAoGCCqGSM49BAMCA0gAMEUCIQCwsXPy
k+JSGQNPVlZyZyI74WDb0y8fnuGaZl0kTQNw2wIgANjimow/L1M5uwl22uAS0tcA
VRXF4m/aMVPkrJTGP6s=
-----END CERTIFICATE-----`)
	demoKey = []byte(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIMLjAzy9GIRvZpJ29CHaW04bWfVtpqrMFTwYPtDpmQBQoAoGCCqGSM49
AwEHoUQDQgAEtBwGcExavS1hi32ovoU/VUtbQw9Ah15ypMOsUCKkvbz7waPCpRMt
kiO7TqFH9gQDZYUz+RKvzh2XGTvvwiEBGg==
-----END EC PRIVATE KEY-----`)
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
