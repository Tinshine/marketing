package item

import (
	"context"
	"marketing/consts/errs"
	model "marketing/manager/resource/model/item"
	"marketing/util"
	"time"

	"github.com/pkg/errors"
)

func Query(ctx context.Context, rq *model.QueryReq) (*model.QueryResp, error) {
	resp := &model.QueryResp{}
	resp.Data = make([]*model.RespModel, 0, 1)

	dao := model.InitDAO()
	if rq.Id != nil {
		item, err := dao.FindById(ctx, *rq.Id)
		if err != nil {
			return nil, errors.WithMessage(err, "model FindById")
		}
		resp.Data = append(resp.Data, item.ToRespModel())
		resp.Total = 1
		return resp, nil
	}

	if rq.AppId == 0 {
		return nil, errors.WithMessage(errs.InvalidParams, "app_id is required")
	}
	if rq.ItemId == nil {
		item, err := dao.FindByItemId(ctx, rq.AppId, *rq.ItemId)
		if err != nil {
			return nil, errors.WithMessage(err, "model FindByItemId")
		}
		resp.Data = append(resp.Data, item.ToRespModel())
		resp.Total = 1
		return resp, nil
	}

	items, err := dao.FindByAppId(ctx, rq.AppId)
	if err != nil {
		return nil, errors.WithMessage(err, "model FindByAppId")
	}
	for _, item := range items {
		resp.Data = append(resp.Data, item.ToRespModel())
	}
	resp.Total = len(items)
	return resp, nil
}

func Add(ctx context.Context, rq *model.AddReq) (*model.AddResp, error) {
	if err := rq.Validate(); err != nil {
		return nil, errors.WithMessage(err, "model Validate")
	}
	username, err := util.GetUsername(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "util GetUsername")
	}

	item, err := model.InitDAO().FirstOrInit(ctx, rq, username)
	if err != nil {
		return nil, errors.WithMessage(err, "model FirstOrInit")
	}
	resp := new(model.AddResp)
	resp.Id = item.Id
	return resp, nil
}

func Delete(ctx context.Context, rq *model.DeleteReq) error {
	if err := rq.Validate(); err != nil {
		return errors.WithMessage(err, "model Validate")
	}

	err := model.InitDAO().DeleteById(ctx, rq.Id)
	if err != nil {
		return errors.WithMessage(err, "model DeleteById")
	}
	return nil
}

func Update(ctx context.Context, rq *model.UpdateReq) error {
	if err := rq.Validate(); err != nil {
		return errors.WithMessage(err, "model Validate")
	}
	username, err := util.GetUsername(ctx)
	if err != nil {
		return errors.WithMessage(err, "util GetUsername")
	}

	if err = model.InitDAO().UpdateById(ctx, rq.Id, map[string]interface{}{
		"descr":      rq.Descr,
		"updated_by": username,
		"updated_at": time.Now().Unix(),
	}); err != nil {
		return errors.WithMessage(err, "model UpdateById")
	}

	return nil
}
