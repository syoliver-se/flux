package repl

import (
	"context"

	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/internal/operation"
	"github.com/syoliver-se/flux/lang"
	"github.com/syoliver-se/flux/plan"
)

// CompilerType specific to the Flux REPL
const CompilerType = "REPL"

// Compiler specific to the Flux REPL
type Compiler struct {
	Spec *operation.Spec `json:"spec"`
}

func (c Compiler) Compile(ctx context.Context, runtime flux.Runtime) (flux.Program, error) {
	planner := plan.PlannerBuilder{}.Build()
	ps, err := planner.Plan(ctx, c.Spec)
	if err != nil {
		return nil, err
	}

	return &lang.Program{
		PlanSpec: ps,
	}, err
}

func (c Compiler) CompilerType() flux.CompilerType {
	return CompilerType
}

func AddCompilerToMappings(mappings flux.CompilerMappings) error {
	return mappings.Add(CompilerType, func() flux.Compiler {
		return new(Compiler)
	})
}
