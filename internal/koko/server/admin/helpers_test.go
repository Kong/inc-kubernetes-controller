package admin

import (
	"net/http/httptest"
	"testing"

	"github.com/kong/inc-kubernetes-controller/internal/koko/log"
	"github.com/kong/inc-kubernetes-controller/internal/koko/plugin"
	"github.com/kong/inc-kubernetes-controller/internal/koko/plugin/validators"
	"github.com/kong/inc-kubernetes-controller/internal/koko/resource"
	serverUtil "github.com/kong/inc-kubernetes-controller/internal/koko/server/util"
	"github.com/kong/inc-kubernetes-controller/internal/koko/store"
	"github.com/kong/inc-kubernetes-controller/internal/koko/test/util"
	"github.com/stretchr/testify/require"
)

var validator plugin.Validator

func init() {
	luaValidator, err := validators.NewLuaValidator(validators.Opts{Logger: log.Logger})
	if err != nil {
		panic(err)
	}
	err = luaValidator.LoadSchemasFromEmbed(plugin.Schemas, "schemas")
	if err != nil {
		panic(err)
	}
	validator = luaValidator
	resource.SetValidator(validator)
}

func setup(t *testing.T) (*httptest.Server, func()) {
	p, err := util.GetPersister(t)
	require.Nil(t, err)
	objectStore := store.New(p, log.Logger)

	server, cleanup := setupWithDB(t, objectStore.ForCluster(store.DefaultCluster))
	return server, func() {
		cleanup()
	}
}

func setupWithDB(t *testing.T, store store.Store) (*httptest.Server, func()) {
	storeLoader := serverUtil.DefaultStoreLoader{
		Store: store,
	}
	handler, err := NewHandler(HandlerOpts{
		Logger:      log.Logger,
		StoreLoader: storeLoader,
		Validator:   validator,
	})
	if err != nil {
		t.Fatalf("creating httptest.Server: %v", err)
	}

	// Because the Validator is created before the StoreLoader for most tests the following mechanism
	// has been established to set the store loader and update the resource Validator appropriately.
	luaValidator, ok := validator.(*validators.LuaValidator)
	if ok {
		luaValidator.SetStoreLoader(storeLoader)
		resource.SetValidator(luaValidator)
	}

	h := serverUtil.HandlerWithRecovery(serverUtil.HandlerWithLogger(handler, log.Logger), log.Logger)
	s := httptest.NewServer(h)
	return s, func() {
		s.Close()
	}
}