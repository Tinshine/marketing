package gift

import (
	"context"
	"marketing/consts/errs"
	"marketing/database/rds"
	model "marketing/manager/resource/model/gift"
	"marketing/util"

	"github.com/pkg/errors"
)

func Query(ctx context.Context, rq *model.QueryReq) (*model.QueryResp, error) {
	resp := new(model.QueryResp)
	if rq.Id != nil {
		gift, err := model.FindById(ctx, *rq.Id)
		if err != nil {
			return nil, errors.WithMessage(err, "find by id")
		}
		resp.Data = []*model.RespModel{gift.ToRespModel()}
		resp.Total = 1
		return resp, nil
	}

	if rq.AppId == 0 {
		return nil, errors.WithMessage(errs.InvalidParams, "app_id is required")
	}

	if rq.GroupId != nil {
		gifts, err := model.FindByGroupId(ctx, rq.AppId, *rq.GroupId)
		if err != nil {
			return nil, errors.WithMessage(err, "find by group_id")
		}
		resp.Data = make([]*model.RespModel, len(gifts))
		for i, gift := range gifts {
			resp.Data[i] = gift.ToRespModel()
		}
		resp.Total = len(gifts)
		return resp, nil
	}

	gifts, err := model.FindByAppId(ctx, rq.AppId)
	if err != nil {
		return nil, errors.WithMessage(err, "find by app_id")
	}
	resp.Data = make([]*model.RespModel, len(gifts))
	for i, gift := range gifts {
		resp.Data[i] = gift.ToRespModel()
	}
	resp.Total = len(gifts)
	return resp, nil
}

func Add(ctx context.Context, rq *model.AddReq) (*model.AddResp, error) {
	if err := rq.Validate(); err != nil {
		return nil, errors.WithMessage(err, "validate")
	}

	username, err := util.GetUsername(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "get username")
	}

	gift, err := rq.ToModel(username)
	if err != nil {
		return nil, errors.WithMessage(err, "to model")
	}

	if err := rds.GetDB(ctx).Create(&gift).Error; err != nil {
		return nil, errors.WithMessage(err, "create")
	}
	resp := new(model.AddResp)
	resp.Id = gift.Id
	return resp, nil
}
