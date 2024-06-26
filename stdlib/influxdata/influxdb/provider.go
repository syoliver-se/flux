package influxdb

import (
	"context"

	"github.com/syoliver-se/flux/dependencies/influxdb"
)

type (
	Dependency            = influxdb.Dependency
	Provider              = influxdb.Provider
	UnimplementedProvider = influxdb.UnimplementedProvider
)

func GetProvider(ctx context.Context) Provider {
	return influxdb.GetProvider(ctx)
}
