package main

import (
	"context"
	authConst "marketing/consts/auth"
	"marketing/manager/auth/handler/auth"
	"marketing/manager/auth/handler/resource"
	"marketing/manager/middleware"
	"marketing/manager/resource/handler/item"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hConsts "github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/hertz/pkg/route"
)

func register(h *server.Hertz) {
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(hConsts.StatusOK, "pong")
	})

	h.Use(middleware.SSO)

	registerAuthResource(h)
	registerAuth(h)

	registerItem(h.Group("resource"))
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

func registerItem(g *route.RouterGroup) {
	g.GET("/item/query", middleware.AuthCheck(authConst.ResourceItem, authConst.Query), item.Query)
}
