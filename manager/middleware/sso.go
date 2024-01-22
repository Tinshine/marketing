package middleware

import (
	"context"
	"marketing/consts"

	"github.com/cloudwego/hertz/pkg/app"
)

func SSO(ctx context.Context, c *app.RequestContext) {
	c.Next(context.WithValue(ctx, consts.CtxUsernameKey, "tinshine"))
	// todo...
}
