package table_test

import (
	"testing"
	"time"

	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/execute"
	"github.com/syoliver-se/flux/plan"
	"github.com/syoliver-se/flux/plan/plantest"
	"github.com/syoliver-se/flux/stdlib/experimental/table"
	"github.com/syoliver-se/flux/stdlib/influxdata/influxdb"
	"github.com/syoliver-se/flux/stdlib/universe"
)

func TestIdempotentTableFill(t *testing.T) {
	windowSpec := universe.WindowProcedureSpec{
		Window: plan.WindowSpec{
			Every:  flux.ConvertDuration(time.Minute),
			Period: flux.ConvertDuration(time.Minute),
		},
		TimeColumn:  execute.DefaultTimeColLabel,
		StartColumn: execute.DefaultStartColLabel,
		StopColumn:  execute.DefaultStopColLabel,
	}

	tc := plantest.RuleTestCase{
		Rules: []plan.Rule{
			table.IdempotentTableFill{},
		},
		Before: &plantest.PlanSpec{
			Nodes: []plan.Node{
				plan.CreateLogicalNode("from", &influxdb.FromProcedureSpec{}),
				plan.CreateLogicalNode("fill0", &table.FillProcedureSpec{}),
				plan.CreateLogicalNode("fill1", &table.FillProcedureSpec{}),
				plan.CreateLogicalNode("window", &windowSpec),
			},
			Edges: [][2]int{
				{0, 1},
				{1, 2},
				{2, 3},
			},
		},
		After: &plantest.PlanSpec{
			Nodes: []plan.Node{
				plan.CreateLogicalNode("from", &influxdb.FromProcedureSpec{}),
				plan.CreateLogicalNode("fill0", &table.FillProcedureSpec{}),
				plan.CreateLogicalNode("window", &windowSpec),
			},
			Edges: [][2]int{
				{0, 1},
				{1, 2},
			},
		},
	}
	plantest.LogicalRuleTestHelper(t, &tc)
}
