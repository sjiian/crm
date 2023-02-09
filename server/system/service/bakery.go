package service

import (
	"context"
	"fmt"

	"github.com/cortezaproject/corteza/server/pkg/errors"
	"github.com/cortezaproject/corteza/server/store"
	"github.com/cortezaproject/corteza/server/system/types"
	"github.com/davecgh/go-spew/spew"
)

type (
	bakery struct {
		store store.Storer
	}

	BakeryService interface {
		Bake(ctx context.Context, qty int, code string) error
	}
)

func Bakery() *bakery {
	return &bakery{
		store: DefaultStore,
	}
}

func (svc bakery) Bake(ctx context.Context, qty int, code string) (err error) {
	spew.Dump(fmt.Sprintf("baking %d of %s", qty, code))

	return store.CreateOrder(ctx, svc.store, &types.Order{
		ID:        nextID(),
		Quantity:  qty,
		Code:      code,
		CreatedAt: *now(),
	})
}

func loadOrder(ctx context.Context, s store.Orders, ID uint64) (res *types.Order, err error) {
	if ID == 0 {
		return nil, UserErrInvalidID()
	}

	if res, err = store.LookupOrderByID(ctx, s, ID); errors.IsNotFound(err) {
		return nil, UserErrNotFound()
	}

	return
}
