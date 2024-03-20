package main

import (
	"context"
	"encoding/json"

	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/ast"
	"github.com/syoliver-se/flux/cmd/flux/cmd"
	"github.com/syoliver-se/flux/codes"
	"github.com/syoliver-se/flux/dependencies/testing"
	"github.com/syoliver-se/flux/dependency"
	"github.com/syoliver-se/flux/execute/executetest"
	"github.com/syoliver-se/flux/internal/errors"
	"github.com/syoliver-se/flux/lang"
	"github.com/syoliver-se/flux/memory"
	"github.com/syoliver-se/flux/runtime"
)

func NewTestExecutor(ctx context.Context) (cmd.TestExecutor, error) {
	return testExecutor{}, nil
}

type testExecutor struct{}

func (testExecutor) Run(pkg *ast.Package, fn cmd.TestResultFunc) error {
	jsonAST, err := json.Marshal(pkg)
	if err != nil {
		return err
	}
	c := lang.ASTCompiler{AST: jsonAST}

	ctx, span := dependency.Inject(context.Background(),
		executetest.NewTestExecuteDependencies(),
		testing.FrameworkConfig{},
	)
	defer span.Finish()
	program, err := c.Compile(ctx, runtime.Default)
	if err != nil {
		return errors.Wrap(err, codes.Invalid, "failed to compile")
	}

	alloc := &memory.ResourceAllocator{}
	query, err := program.Start(ctx, alloc)
	if err != nil {
		return errors.Wrap(err, codes.Inherit, "error while executing program")
	}

	results := flux.NewResultIteratorFromQuery(query)
	defer results.Release()

	return fn(ctx, results)
}

func (testExecutor) Close() error { return nil }
