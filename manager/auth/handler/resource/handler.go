package resource

import (
	"context"
	"marketing/consts/errs"
	model "marketing/manager/auth/model/resource"
	service "marketing/manager/auth/service/resource"
	"marketing/util"
	"marketing/util/log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/pkg/errors"
)

func Query(c context.Context, ctx *app.RequestContext) {
	req := new(model.QueryReq)
	if err := ctx.Bind(&req); err != nil {
		util.Error(ctx, errors.WithMessage(errs.Bind, err.Error()))
		return
	}

	resp, err := service.Query(c, req)
	if err != nil {
		log.Error("Query.service.Query", err, "req", req)
		util.Error(ctx, errors.WithMessage(errs.Internal, err.Error()))
		return
	}

	util.JSON(ctx, resp)
}

func Delete(c context.Context, ctx *app.RequestContext) {
	req := new(model.DeleteReq)
	if err := ctx.Bind(&req); err != nil {
		util.Error(ctx, errors.WithMessage(errs.Bind, err.Error()))
		return
	}

	err := service.Delete(c, req)
	if err != nil {
		log.Error("Delete.service.Delete", err, "req", req)
		util.Error(ctx, errors.WithMessage(errs.Internal, err.Error()))
		return
	}

	util.JSON(ctx, nil)
}

func Add(c context.Context, ctx *app.RequestContext) {
	req := new(model.AddReq)
	if err := ctx.Bind(&req); err != nil {
		util.Error(ctx, errors.WithMessage(errs.Bind, err.Error()))
		return
	}

	err := service.Add(c, req)
	if err != nil {
		log.Error("Add.service.Add", err, "req", req)
		util.Error(ctx, errors.WithMessage(errs.Internal, err.Error()))
		return
	}

	util.JSON(ctx, nil)
}
