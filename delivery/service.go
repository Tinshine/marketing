package delivery

import (
	"context"
	"marketing/consts/errs"
	"marketing/engine/transcation/model"

	"github.com/pkg/errors"
)

func checkParams(ctx context.Context, params *model.Params) error {
	if params == nil {
		return errors.WithMessage(errs.InvalidParams, "params is nil")
	}
	if params.User == nil {
		return errors.WithMessage(errs.InvalidParams, "params.User is nil")
	}

	return nil
}
