package resource

import (
	"context"
	"errors"
	"fmt"
	"strings"

	v1 "github.com/kong/inc-kubernetes-controller/internal/koko/gen/grpc/kong/admin/model/v1"
	"github.com/kong/inc-kubernetes-controller/internal/koko/model"
	"github.com/kong/inc-kubernetes-controller/internal/koko/model/json/extension"
	"github.com/kong/inc-kubernetes-controller/internal/koko/model/json/generator"
	"github.com/kong/inc-kubernetes-controller/internal/koko/model/json/validation"
	"github.com/kong/inc-kubernetes-controller/internal/koko/model/json/validation/typedefs"
)

const (
	TypeSNI model.Type = "sni"
)

type SNI struct {
	SNI *v1.SNI
}

func NewSNI() SNI {
	return SNI{
		SNI: &v1.SNI{},
	}
}

func (r SNI) ID() string {
	if r.SNI == nil {
		return ""
	}
	return r.SNI.Id
}

func (r SNI) Type() model.Type {
	return TypeSNI
}

func (r SNI) Resource() model.Resource {
	return r.SNI
}

// SetResource implements the Object.SetResource interface.
func (r SNI) SetResource(pr model.Resource) error { return model.SetResource(r, pr) }

func (r SNI) Validate(ctx context.Context) error {
	if hostnameCheck(r.SNI.Name) != typeName {
		errWrap := validation.Error{}
		errWrap.Errs = append(errWrap.Errs, &v1.ErrorDetail{
			Type:     v1.ErrorType_ERROR_TYPE_FIELD,
			Field:    "name",
			Messages: []string{"must not be an IP"},
		})
		return errWrap
	}
	if len(strings.Split(r.SNI.Name, ":")) > 1 {
		errWrap := validation.Error{}
		errWrap.Errs = append(errWrap.Errs, &v1.ErrorDetail{
			Type:  v1.ErrorType_ERROR_TYPE_FIELD,
			Field: "name",
			Messages: []string{
				"must not contain a port",
			},
		})
		return errWrap
	}
	err := validation.Validate(string(TypeSNI), r.SNI)
	if err != nil {
		return err
	}
	return nil
}

func (r SNI) ProcessDefaults(ctx context.Context) error {
	if r.SNI == nil {
		return fmt.Errorf("invalid nil resource")
	}
	defaultID(&r.SNI.Id)
	return nil
}

func (r SNI) Indexes() []model.Index {
	if r.SNI.Certificate == nil {
		panic(errors.New("Certificate can not be nil"))
	}
	res := []model.Index{
		{
			Name:      "name",
			Type:      model.IndexUnique,
			Value:     r.SNI.Name,
			FieldName: "name",
		},
	}
	res = append(res, model.Index{
		Name:        "certificate_id",
		Type:        model.IndexForeign,
		ForeignType: TypeCertificate,
		FieldName:   "certificate.id",
		Value:       r.SNI.Certificate.Id,
	})
	return res
}

func init() {
	err := model.RegisterType(TypeSNI, &v1.SNI{}, func() model.Object {
		return NewSNI()
	})
	if err != nil {
		panic(err)
	}

	sniSchema := &generator.Schema{
		Type: "object",
		Properties: map[string]*generator.Schema{
			"id":          typedefs.ID,
			"name":        typedefs.WilcardHost,
			"certificate": typedefs.ReferenceObject,
			"tags":        typedefs.Tags,
			"created_at":  typedefs.UnixEpoch,
			"updated_at":  typedefs.UnixEpoch,
		},
		Required: []string{
			"id",
			"name",
			"certificate",
		},
		XKokoConfig: &extension.Config{
			ResourceAPIPath: "snis",
		},
	}
	err = generator.DefaultRegistry.Register(string(TypeSNI), sniSchema)
	if err != nil {
		panic(err)
	}
}
