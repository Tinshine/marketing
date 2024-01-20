package main

import (
	"context"
	"marketing/item_manager/handler/resource/item"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func register(h *server.Hertz) {
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "pong")
	})

	registerItem(h)
}

func registerItem(h *server.Hertz) {
	h.GET("/item/query", item.Query)
}
