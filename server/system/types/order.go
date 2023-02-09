package types

import (
	"time"

	"github.com/cortezaproject/corteza/server/pkg/filter"
)

type (
	Order struct {
		ID       uint64 `json:"orderID,string"`
		Quantity int    `json:"quantity"`
		Code     string `json:"code"`
		Comment  string `json:"code"`

		CreatedAt time.Time  `json:"createdAt,omitempty"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		DeletedAt *time.Time `json:"deletedAt,omitempty"`
	}

	OrderFilter struct {
		OrderID []uint64 `json:"orderID,string"`

		Query string `json:"query"`

		Deleted   filter.State `json:"deleted"`
		Suspended filter.State `json:"suspended"`

		// Check fn is called by store backend for each resource found function can
		// modify the resource and return false if store should not return it
		//
		// Store then loads additional resources to satisfy the paging parameters
		Check func(*Order) (bool, error) `json:"-"`

		// Standard helpers for paging and sorting
		filter.Sorting
		filter.Paging
	}
)
