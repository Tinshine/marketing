package main

import (
	"context"
	authConst "marketing/consts/auth"
	"marketing/manager/auth/handler/auth"
	"marketing/manager/auth/handler/resource"
	"marketing/manager/middleware"
	"marketing/manager/resource/handler/gift"
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

	registerAuth(h.Group("/auth"))
	registerResource(h.Group("/resource"))
}

func registerAuthResource(g *route.RouterGroup) {
	g.GET("/resource/query", resource.Query)
	g.DELETE("/resource/delete", resource.Delete)
	g.POST("/resource/add", resource.Add)
}

func registerAuth(g *route.RouterGroup) {
	registerAuthResource(g)
	g.GET("/query", auth.Query)
	g.DELETE("/delete", auth.Delete)
	g.POST("/add", auth.Add)
	g.GET("/check", auth.Check)
}

func registerResource(g *route.RouterGroup) {
	registerItem(g)
	registerGift(g)
}

func registerItem(g *route.RouterGroup) {
	g.GET("/item/query", middleware.AuthCheck(authConst.ResourceItem, authConst.Query), item.Query)
	g.POST("/item/add", middleware.AuthCheck(authConst.ResourceItem, authConst.Add), item.Add)
	g.DELETE("/item/delete", middleware.AuthCheck(authConst.ResourceItem, authConst.Delete), item.Delete)
	g.PUT("/item/update", middleware.AuthCheck(authConst.ResourceItem, authConst.Update), item.Update)
}

func registerGift(g *route.RouterGroup) {
	g.GET("/gift/query", middleware.AuthCheck(authConst.ResourceGift, authConst.Query), gift.Query)
	g.POST("/gift/add", middleware.AuthCheck(authConst.ResourceGift, authConst.Add), gift.Add)
	// g.DELETE("/gift/delete", middleware.AuthCheck(authConst.ResourceGift, authConst.Delete), gift.Delete)
	// g.PUT("/gift/update", middleware.AuthCheck(authConst.ResourceGift, authConst.Update), gift.Update)
}
