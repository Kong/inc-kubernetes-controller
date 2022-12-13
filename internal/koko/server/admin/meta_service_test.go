package admin

import (
	"context"
	"testing"

	"github.com/kong/inc-kubernetes-controller/internal/koko/log"
	"github.com/stretchr/testify/require"
)

func TestMetaService_GetVersion(t *testing.T) {
	s := MetaService{Logger: log.Logger}
	resp, err := s.GetVersion(context.Background(), nil)
	expectedVersion := "dev"
	require.Equal(t, resp.Version, expectedVersion, "unexpected version")
	require.Nil(t, err, "expected no error")
}
