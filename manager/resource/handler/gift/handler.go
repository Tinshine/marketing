package gift

import (
	"context"
	"marketing/consts/errs"
	model "marketing/manager/resource/model/gift"

	service "marketing/manager/resource/service/gift"
	"marketing/util"
	"marketing/util/log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/pkg/errors"
)

func Query(ctx context.Context, c *app.RequestContext) {
	req := new(model.QueryReq)
	if err := c.Bind(req); err != nil {
		log.Error("Query.Bind", err)
		util.Error(c, errors.WithMessage(errs.Bind, err.Error()))
		return
	}

	resp, err := service.Query(ctx, req)
	if err != nil {
		log.Error("Query.service.Query", err, "req", req)
		util.Error(c, errors.WithMessage(errs.Internal, err.Error()))
		return
	}

	util.JSON(c, resp)
}

func Add(ctx context.Context, c *app.RequestContext) {
	req := new(model.AddReq)
	if err := c.Bind(req); err != nil {
		log.Error("Add.Bind", err)
		util.Error(c, errors.WithMessage(errs.Bind, err.Error()))
		return
	}

	resp, err := service.Add(ctx, req)
	if err != nil {
		log.Error("Add.service.Add", err, "req", req)
		util.Error(c, errors.WithMessage(errs.Internal, err.Error()))
		return
	}

	util.JSON(c, resp)
}

func Update(ctx context.Context, c *app.RequestContext) {
	req := new(model.UpdateReq)
	if err := c.Bind(req); err != nil {
		log.Error("Update.Bind", err)
		util.Error(c, errors.WithMessage(errs.Bind, err.Error()))
		return
	}

	err := service.Update(ctx, req)
	if err != nil {
		log.Error("Update.service.Update", err, "req", req)
		util.Error(c, errors.WithMessage(errs.Internal, err.Error()))
		return
	}

	util.JSON(c, nil)
}

func Sync(ctx context.Context, c *app.RequestContext) {
	req := new(model.SyncReq)
	if err := c.Bind(req); err != nil {
		log.Error("Sync.Bind", err)
		util.Error(c, errors.WithMessage(errs.Bind, err.Error()))
		return
	}

	resp, err := service.Sync(ctx, req)
	if err != nil {
		log.Error("Sync.service.Sync", err, "req", req)
		util.Error(c, errors.WithMessage(errs.Internal, err.Error()))
		return
	}

	util.JSON(c, resp)
}
