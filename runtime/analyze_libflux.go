package runtime

import (
	"context"

	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/libflux/go/libflux"
	"github.com/syoliver-se/flux/semantic"
)

// AnalyzeSource parses and analyzes the given Flux source,
// using libflux.
func AnalyzeSource(ctx context.Context, fluxSrc string) (*semantic.Package, error) {
	ast := libflux.ParseString(fluxSrc)
	return AnalyzePackage(ctx, ast)
}

func AnalyzePackage(ctx context.Context, astPkg flux.ASTHandle) (*semantic.Package, error) {
	hdl := astPkg.(*libflux.ASTPkg)
	defer hdl.Free()

	options := libflux.NewOptions(ctx)
	sem, err := libflux.AnalyzeWithOptions(hdl, options)
	if err != nil {
		return nil, err
	}
	defer sem.Free()
	bs, err := sem.MarshalFB()
	if err != nil {
		return nil, err
	}
	return semantic.DeserializeFromFlatBuffer(bs)
}
