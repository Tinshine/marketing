package item

import (
	"context"
	"marketing/util/log"

	"github.com/cloudwego/hertz/pkg/app"
)

func Query(c context.Context, ctx *app.RequestContext) {
	req := new(QueryReq)
	if err := ctx.Bind(req); err != nil {
		log.Error("Query.Bind", err)
		ctx.Error(err)
		return
	}

	resp := new(QueryResp)
	resp.Items = make([]ItemResp, 0)
	resp.Total = 0

	ctx.JSON(200, resp)
}
