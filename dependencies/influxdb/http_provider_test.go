package influxdb

import (
	"context"
	nethttp "net/http"
	"testing"

	"github.com/syoliver-se/flux"
	"github.com/syoliver-se/flux/codes"
	"github.com/syoliver-se/flux/dependencies/http"
	fluxurl "github.com/syoliver-se/flux/dependencies/url"
	"github.com/syoliver-se/flux/dependency"
	"github.com/syoliver-se/flux/internal/errors"
)

type mockClient struct{}

var _ http.Client = mockClient{}

func (c mockClient) Do(req *nethttp.Request) (*nethttp.Response, error) {
	return nil, errors.New(codes.Internal, "mock error")
}

func TestPrivateClient(t *testing.T) {
	h := HttpProvider{
		DefaultConfig: Config{
			Host: "http://myhost.com:8085",
		},
	}
	deps := flux.Deps{Deps: flux.WrappedDeps{
		HTTPClient:   mockClient{},
		URLValidator: fluxurl.PrivateIPValidator{},
	}}
	ctx, _ := dependency.Inject(context.Background(), deps)
	c, err := h.clientFor(ctx, Config{})
	if err != nil {
		t.Error(err)
	}
	_, err = c.Client.Do(&nethttp.Request{})
	if err != nil {
		if err.Error() != "an internal error has occurred: mock error" {
			t.Fatalf("got unexpected error: %s", err)
		}
	} else {
		t.Fatal("expected error but got none")
	}
}
