package store

import "github.com/kong/inc-kubernetes-controller/internal/koko/test/util"

func init() {
	util.RegisterSchemasFromFS()
}