package mock

import (
	"context"
	"github.com/syoliver-se/flux/codes"
	"github.com/syoliver-se/flux/internal/errors"
)

type SecretService map[string]string

func (s SecretService) LoadSecret(ctx context.Context, k string) (string, error) {
	v, ok := s[k]
	if ok {
		return v, nil
	}
	return "", errors.Newf(codes.NotFound, "secret key %q not found", k)
}
