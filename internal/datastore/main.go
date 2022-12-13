package datastore

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/kong/inc-kubernetes-controller/internal/koko/persistence"
	"github.com/kong/inc-kubernetes-controller/internal/koko/persistence/sqlite"
)

type Store struct {
	Database persistence.Persister
}

// Setup creates a datastore
func (s *Store) Setup(logger *zap.Logger) error {
	// KOKO Migration code removed since DB is always in memory

	// KOKO Stripped down from internal/config/db.go Config type
	// Dialect is always SQLite
	// SQLite is always InMemory
	// QueryTimeout is hardcoded
	store, err := sqlite.New(sqlite.Opts{InMemory: true}, time.Minute, logger)
	if err != nil {
		return err
	}
	s.Database = store
	return nil
}

// Start implements the Start function for controller-runtime Runnable interface
func (s *Store) Start(ctx context.Context) error {
	return nil
}

// NeedLeaderElection implements the NeedLeaderElection function for controller-runtime LeaderElectionRunnable interface
func (s *Store) NeedLeaderElection() bool {
	return true
}
