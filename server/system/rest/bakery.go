package rest

import (
	"context"

	"github.com/cortezaproject/corteza/server/pkg/api"
	"github.com/cortezaproject/corteza/server/system/rest/request"
	"github.com/cortezaproject/corteza/server/system/service"
	"github.com/pkg/errors"
)

var _ = errors.Wrap

type (
	Bakery struct {
		bakery bakeryService
	}

	bakeryService interface {
		Bake(context.Context, int, string) error
	}
)

func (Bakery) New() *Bakery {
	return &Bakery{
		bakery: service.DefaultBakery,
	}
}

func (ctrl Bakery) Bake(ctx context.Context, r *request.BakeryBake) (interface{}, error) {
	return api.OK(), ctrl.bakery.Bake(ctx, r.Quantity, r.Code)
}
