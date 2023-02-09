package handlers

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//

import (
	"context"
	"github.com/cortezaproject/corteza/server/pkg/api"
	"github.com/cortezaproject/corteza/server/system/rest/request"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type (
	// Internal API interface
	BakeryAPI interface {
		Bake(context.Context, *request.BakeryBake) (interface{}, error)
	}

	// HTTP API interface
	Bakery struct {
		Bake func(http.ResponseWriter, *http.Request)
	}
)

func NewBakery(h BakeryAPI) *Bakery {
	return &Bakery{
		Bake: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewBakeryBake()
			if err := params.Fill(r); err != nil {
				api.Send(w, r, err)
				return
			}

			value, err := h.Bake(r.Context(), params)
			if err != nil {
				api.Send(w, r, err)
				return
			}

			api.Send(w, r, value)
		},
	}
}

func (h Bakery) MountRoutes(r chi.Router, middlewares ...func(http.Handler) http.Handler) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares...)
		r.Post("/bakery/bake", h.Bake)
	})
}
