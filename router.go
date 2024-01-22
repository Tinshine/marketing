package main

import (
	"context"
	"marketing/manager/auth/handler/auth"
	"marketing/manager/auth/handler/resource"
	"marketing/manager/resource/handler/item"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func register(h *server.Hertz) {
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "pong")
	})

	registerAuthResource(h)
	registerAuth(h)
	registerItem(h)
}

func registerItem(h *server.Hertz) {
	h.GET("/resource/item/query", item.Query)
}

func registerAuthResource(h *server.Hertz) {
	h.GET("/auth/resource/query", resource.Query)
	h.DELETE("/auth/resource/delete", resource.Delete)
	h.POST("/auth/resource/add", resource.Add)
}

func registerAuth(h *server.Hertz) {
	h.GET("/auth/query", auth.Query)
	h.DELETE("/auth/delete", auth.Delete)
	h.POST("/auth/add", auth.Add)
	h.GET("/auth/check", auth.Check)
}
