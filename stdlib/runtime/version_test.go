//go:build go1.12
// +build go1.12

package runtime_test

import (
	"context"
	"runtime/debug"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/syoliver-se/flux/codes"
	"github.com/syoliver-se/flux/dependencies/dependenciestest"
	"github.com/syoliver-se/flux/dependency"
	"github.com/syoliver-se/flux/internal/errors"
	"github.com/syoliver-se/flux/stdlib/runtime"
	"github.com/syoliver-se/flux/values"
)

func TestVersion(t *testing.T) {
	for _, tt := range []struct {
		name    string
		bi      *debug.BuildInfo
		want    values.Value
		wantErr error
	}{
		{
			name: "main module",
			bi: &debug.BuildInfo{
				Path: "github.com/syoliver-se/flux",
				Main: debug.Module{
					Path:    "github.com/syoliver-se/flux",
					Version: "v0.38.0",
				},
			},
			want: values.NewString("v0.38.0"),
		},
		{
			name: "replaced main module",
			bi: &debug.BuildInfo{
				Path: "github.com/syoliver-se/flux",
				Main: debug.Module{
					Path:    "github.com/syoliver-se/flux",
					Version: "v0.38.0",
					Replace: &debug.Module{
						Path:    "github.com/syoliver-se/flux",
						Version: "(devel)",
					},
				},
			},
			want: values.NewString("(devel)"),
		},
		{
			name: "dependency module",
			bi: &debug.BuildInfo{
				Path: "github.com/influxdata/influxdb",
				Main: debug.Module{
					Path:    "github.com/influxdata/influxdb",
					Version: "v2.0.0",
				},
				Deps: []*debug.Module{
					{
						Path:    "github.com/syoliver-se/flux",
						Version: "v0.38.0",
					},
				},
			},
			want: values.NewString("v0.38.0"),
		},
		{
			name: "replaced dependency module",
			bi: &debug.BuildInfo{
				Path: "github.com/influxdata/influxdb",
				Main: debug.Module{
					Path:    "github.com/influxdata/influxdb",
					Version: "v2.0.0",
				},
				Deps: []*debug.Module{
					{
						Path:    "github.com/syoliver-se/flux",
						Version: "v0.38.0",
						Replace: &debug.Module{
							Path: "github.com/syoliver-se/flux",
						},
					},
				},
			},
			want: values.NewString(""),
		},
		{
			name:    "build info not present",
			wantErr: errors.New(codes.NotFound, "build info is not present"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			runtime.SetBuildInfo(tt.bi)

			ctx, deps := dependency.Inject(context.Background(), dependenciestest.Default())
			defer deps.Finish()
			got, err := runtime.Version(ctx, nil)
			if err != nil {
				if tt.wantErr != nil {
					if !cmp.Equal(tt.wantErr, err) {
						t.Fatalf("unexpected error -want/+got:\n%s", cmp.Diff(tt.wantErr, err))
					}
					return
				} else {
					t.Fatalf("unexpected error: %s", err)
				}
			}

			if !got.Equal(tt.want) {
				t.Fatalf("unexpected value -want/+got:\n\t- %v\n\t+ %v", tt.want, got)
			}
		})
	}
}
