package util

import (
	"context"
	"marketing/consts"
	"marketing/consts/errs"
)

func GetUsername(ctx context.Context) (string, error) {
	v := ctx.Value(consts.CtxUsernameKey)
	if v == nil {
		return "", errs.EmptyUsername
	}
	return v.(string), nil
}
