package plugin

import (
	"context"

	"github.com/cortezaproject/corteza/server/automation/types"
	"github.com/cortezaproject/corteza/server/pkg/expr"
	"github.com/cortezaproject/corteza/server/pkg/plugin/automation"
)

type (
	cp struct {
		af automation.AutomationFunction
	}
)

func (c *cp) automationFunction() *types.Function {
	var (
		f = c.af.Meta()
	)

	f.Handler = func(ctx context.Context, in *expr.Vars) (o *expr.Vars, err error) {
		out, err := c.af.Exec(in)

		if err != nil {
			return nil, err
		}

		o = &expr.Vars{}
		err = o.Assign(out)

		return
	}

	return f
}
