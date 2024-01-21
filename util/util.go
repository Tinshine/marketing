package util

import (
	"context"
	"marketing/const/errs"
)

func GetUsername(ctx context.Context) (string, error) {
	v := ctx.Value("username")
	if v == nil {
		return "", errs.EmptyUsername
	}
	return v.(string), nil
}
