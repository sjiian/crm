package system

import (
	"github.com/cortezaproject/corteza/server/codegen/schema"
)

order: {
	model: {
		attributes: {
		  id:     schema.IdField
		  quantity: {
		  	sortable: true,
		  	goType: "int"
				dal: { type: "Number" }
			}
		  code: {
		  	sortable: true,
		  	unique: true,
		  	ignoreCase: true
				dal: {}
			}
		  created_at: schema.SortableTimestampNowField
		  updated_at: schema.SortableTimestampNilField
		  deleted_at: schema.SortableTimestampNilField
		}

		indexes: {
			"primary": { attribute: "id" }
		}
	}

	features: {
		labels: false
	}

	filter: {
		struct: {
			order_id: {goType: "[]uint64", ident: "userID", storeIdent: "id"}

			deleted: {goType: "filter.State", storeIdent: "deleted_at"}
			suspended: {goType: "filter.State", storeIdent: "suspended_at"}
		}

		query: ["code"]
		byValue: []
		byNilState: ["deleted"]
	}

	rbac: {
		operations: {
			"read": description:         "Read user"
			"update": description:       "Update user"
			"delete": description:       "Delete user"
		}
	}

	store: {
		api: {
			lookups: [
				{
					fields: ["id"]
					description: """
						searches for user by ID

						It returns user even if deleted or suspended
						"""
				}, {
					fields: ["code"]
					nullConstraint: ["deleted_at"]
					constraintCheck: true
					description: """
						searches for user by email

						It returns only valid user (not deleted, not suspended)
						"""
				}
			]
		}
	}
}
