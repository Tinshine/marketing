package util

import (
	"context"
	"errors"
	"fmt"
	"marketing/consts"
	"marketing/consts/errs"
	"os"
	"strings"
)

func GetUsername(ctx context.Context) (string, error) {
	if ctx == nil {
		return "", errors.New("nil context")
	}
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

func IsUnitTest() bool {
	return os.Getenv("unit_test") == "1"
}

func SetUnitTestMode() {
	os.Setenv("unit_test", "1")
}

func UnsetUnitTestMode() {
	os.Unsetenv("unit_test")
}
