package datastore

import (
	"context"
	"embed"
	"time"

	"go.uber.org/zap"

	"github.com/kong/inc-kubernetes-controller/internal/koko/db"
	model "github.com/kong/inc-kubernetes-controller/internal/koko/gen/grpc/kong/admin/model/v1"
	"github.com/kong/inc-kubernetes-controller/internal/koko/persistence"
	"github.com/kong/inc-kubernetes-controller/internal/koko/persistence/sqlite"
	"github.com/kong/inc-kubernetes-controller/internal/koko/store"
)

const dialect = "sqlite3"

var migrations embed.FS

// KOKO Stripped down from internal/config/db.go Config type
// Dialect is always SQLite
// SQLite is always InMemory
// QueryTimeout is hardcoded
func sqliteConfig() sqlite.Opts {
	return sqlite.Opts{InMemory: true}
}

func dbConfig() db.Config {
	return db.Config{
		Dialect:      db.DialectSQLite3,
		SQLite:       sqliteConfig(),
		QueryTimeout: time.Minute,
	}
}

type StoreRunner struct {
	Database persistence.Persister
	Store    *store.ObjectStore
}

// Setup creates a datastore
func (s *StoreRunner) Setup(logger *zap.Logger) error {
	config := dbConfig()
	config.Logger = logger
	m, err := db.NewMigrator(config)
	if err != nil {
		return err
	}
	// KOKO always use sqlite, always migrate, fresh in-memory db

	err = m.Up()
	if err != nil {
		return err
	}

	db, err := sqlite.New(sqliteConfig(), time.Minute, logger)
	if err != nil {
		return err
	}
	s.Database = db
	s.Store = store.New(db, logger.With(zap.String("component", "store"))).ForCluster(store.DefaultCluster)

	return nil
}

func (s *StoreRunner) Load(_ context.Context, _ *model.RequestCluster) store.Store {
	return s.Store
}

// Start implements the Start function for controller-runtime Runnable interface
func (s *StoreRunner) Start(ctx context.Context) error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	if err := s.Setup(logger); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		return nil
	}
}

// NeedLeaderElection implements the NeedLeaderElection function for controller-runtime LeaderElectionRunnable interface
func (s *StoreRunner) NeedLeaderElection() bool {
	return true
}
