package limiter

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

func RateLimiter(c context.Context, ctx *app.RequestContext) {
	ctx.Next(c)
}
