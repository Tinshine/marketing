package resource

import (
	"context"
	"marketing/consts/auth"
	"marketing/database/rds"
	model "marketing/manager/auth/model/resource"
	"marketing/util"

	"github.com/pkg/errors"
)

func Query(ctx context.Context, rq *model.QueryReq) (*model.QueryResp, error) {
	if err := rq.Validate(); err != nil {
		return nil, errors.WithMessage(err, "validate")
	}

	db := rds.TestDB(ctx).Model(model.AuthRes{})
	if rq.ResType != nil {
		db = db.Where("res_type = ?", *rq.ResType)
	}
	if rq.ResId != nil {
		db = db.Where("res_id = ?", *rq.ResId)
	}
	if rq.AuthTypes != nil {
		db = db.Where("auth_type in ?", *rq.AuthTypes)
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
	if err := rds.TestDB(ctx).Delete(model.AuthRes{}, rq.Id).Error; err != nil {
		return errors.WithMessage(err, "db delete")
	}

	return nil
}

func Add(ctx context.Context, rq *model.AddReq) error {
	if err := rq.Validate(); err != nil {
		return errors.WithMessage(err, "validate")
	}
	username, err := util.GetUsername(ctx)
	if err != nil {
		return errors.WithMessage(err, "get username")
	}

	// remove duplicated auth_type and add admin type
	authTypes := map[auth.AuthType]struct{}{}
	for _, a := range rq.AuthTypes {
		authTypes[a] = struct{}{}
	}
	authTypes[auth.Admin] = struct{}{}

	authRes := make([]model.AuthRes, 0, len(authTypes))
	for a := range authTypes {
		authRes = append(authRes, model.AuthRes{
			AppId:     rq.AppId,
			ResType:   rq.ResType,
			ResId:     rq.ResId,
			AuthType:  a,
			CreatedBy: username,
		})
	}

	if err := rds.TestDB(ctx).Create(&authRes).Error; err != nil {
		return errors.WithMessage(err, "db create")
	}
	return nil
}
