package idgen

import (
	"context"
	"marketing/consts/errs"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func Gen(ctx context.Context) (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", errors.WithMessage(errs.Internal, err.Error())
	}
	return u.String(), nil
}
