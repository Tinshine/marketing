package item

import (
	"context"
	"marketing/util"
	"marketing/util/log"

	model "marketing/manager/resource/model/item"
	service "marketing/manager/resource/service/item"

	"github.com/cloudwego/hertz/pkg/app"
)

func Query(c context.Context, ctx *app.RequestContext) {
	req := new(model.QueryReq)
	if err := ctx.Bind(req); err != nil {
		log.Error("Query.Bind", err)
		util.Error(ctx, err)
		return
	}

	resp, err := service.Query(c, req)
	if err != nil {
		log.Error("Query.service.Query", err, "req", req)
		util.Error(ctx, err)
		return
	}

	util.JSON(ctx, resp)
}

func Add(c context.Context, ctx *app.RequestContext) {
	req := new(model.AddReq)
	if err := ctx.Bind(req); err != nil {
		log.Error("Add.Bind", err)
		util.Error(ctx, err)
		return
	}

	resp, err := service.Add(c, req)
	if err != nil {
		log.Error("Add.service.Add", err, "req", req)
		util.Error(ctx, err)
		return
	}

	util.JSON(ctx, resp)
}

func Delete(c context.Context, ctx *app.RequestContext) {
	req := new(model.DeleteReq)
	if err := ctx.Bind(req); err != nil {
		log.Error("Delete.Bind", err)
		util.Error(ctx, err)
		return
	}

	err := service.Delete(c, req)
	if err != nil {
		log.Error("Delete.service.Delete", err, "req", req)
		util.Error(ctx, err)
		return
	}

	util.JSON(ctx, nil)
}

func Update(c context.Context, ctx *app.RequestContext) {
	req := new(model.UpdateReq)
	if err := ctx.Bind(req); err != nil {
		log.Error("Update.Bind", err)
		util.Error(ctx, err)
		return
	}

	err := service.Update(c, req)
	if err != nil {
		log.Error("Update.service.Update", err, "req", req)
		util.Error(ctx, err)
		return
	}

	util.JSON(ctx, nil)
}
