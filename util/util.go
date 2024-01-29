package util

import (
	"context"
	"fmt"
	"marketing/consts"
	"marketing/consts/errs"
	"strings"
)

func GetUsername(ctx context.Context) (string, error) {
	v := ctx.Value(consts.CtxUsernameKey)
	if v == nil {
		return "", errs.EmptyUsername
	}
	return v.(string), nil
}

func MakeKey(items ...interface{}) string {
	if len(items) == 0 {
		return ""
	}
	var sb strings.Builder
	for _, item := range items {
		sb.WriteString(fmt.Sprintf("%v", item))
		sb.WriteString("_")
	}
	key := sb.String()
	return key[:len(key)-1]
}
