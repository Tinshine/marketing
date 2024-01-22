package resource

import (
	"context"
	"marketing/database/rds"
	model "marketing/manager/auth/model/resource"
	"marketing/util"

	"github.com/pkg/errors"
)

func Query(ctx context.Context, rq *model.QueryReq) (*model.QueryResp, error) {
	if err := rq.Validate(); err != nil {
		return nil, errors.WithMessage(err, "validate")
	}

	db := rds.GetDB(ctx).Model(model.AuthRes{})
	if rq.ResType != nil {
		db = db.Where("res_type = ?", *rq.ResType)
	}
	if rq.ResId != nil {
		db = db.Where("res_id = ?", *rq.ResId)
	}
	if rq.AuthType != nil {
		db = db.Where("auth_type = ?", *rq.AuthType)
	}
	if rq.CreatedBy != nil {
		db = db.Where("created_by = ?", *rq.CreatedBy)
	}

	var records []*model.AuthRes
	if err := db.Find(&records).Error; err != nil {
		return nil, errors.WithMessage(err, "db find")
	}

	data := make([]*model.RespModel, 0, len(records))
	for _, r := range records {
		data = append(data, r.ToRespModel())
	}

	resp := new(model.QueryResp)
	resp.Total = len(records)
	resp.Data = data
	return resp, nil
}

func Delete(ctx context.Context, rq *model.DeleteReq) error {
	if err := rq.Validate(); err != nil {
		return errors.WithMessage(err, "validate")
	}
	if err := rds.GetDB(ctx).Delete(model.AuthRes{}, rq.Id).Error; err != nil {
		return errors.WithMessage(err, "db delete")
	}

	return nil
}

func Add(ctx context.Context, rq *model.AddReq) (*model.AddResp, error) {
	if err := rq.Validate(); err != nil {
		return nil, errors.WithMessage(err, "validate")
	}
	username, err := util.GetUsername(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "get username")
	}
	authRes := new(model.AuthRes)
	authRes.AppId = rq.AppId
	authRes.ResType = rq.ResType
	authRes.ResId = rq.ResId
	authRes.AuthType = rq.AuthType
	authRes.CreatedBy = username

	if err := rds.GetDB(ctx).Create(&authRes).Error; err != nil {
		return nil, errors.WithMessage(err, "db create")
	}

	resp := new(model.AddResp)
	resp.Id = authRes.Id
	return resp, nil
}
