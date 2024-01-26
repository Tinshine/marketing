package auth

import (
	"context"
	"marketing/consts/auth"
	"marketing/consts/errs"
	"marketing/database/rds"
	model "marketing/manager/auth/model/auth"
	"marketing/util"
	"marketing/util/log"
	"time"

	"github.com/pkg/errors"
)

func Query(ctx context.Context, rq *model.QueryReq) (*model.QueryResp, error) {
	if err := rq.Validate(); err != nil {
		return nil, errors.WithMessage(err, "validate")
	}

	username, err := util.GetUsername(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "get username")
	}

	db := rds.TestDB(ctx).Model(&model.Auth{})
	db = db.Where("app_id = ? and username = ?", rq.AppId, username)
	if rq.AuthResId != 0 {
		db = db.Where("auth_res_id = ?", rq.AuthResId)
	}
	var records []*model.Auth
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

	if err := rds.TestDB(ctx).Delete(&model.Auth{}, rq.Id).Error; err != nil {
		return errors.WithMessage(err, "db delete")
	}

	return nil
}

func Add(ctx context.Context, rq *model.AddReq) (*model.AddResp, error) {
	if err := rq.Validate(); err != nil {
		return nil, errors.WithMessage(err, "validate")
	}

	grantor, err := util.GetUsername(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "get username")
	}

	record := &model.Auth{
		AppId:     rq.AppId,
		Username:  rq.Username,
		AuthResId: rq.AuthResId,
		CreatedBy: grantor,
		ExpiredAt: rq.ExpiredAt,
	}
	if err := rds.TestDB(ctx).Create(&record).Error; err != nil {
		return nil, errors.WithMessage(err, "db create")
	}

	resp := new(model.AddResp)
	resp.Id = record.Id
	return resp, nil
}

func Check(ctx context.Context, rq *model.CheckReq) (*model.CheckResp, error) {
	if err := rq.Validate(); err != nil {
		return nil, errors.WithMessage(err, "validate")
	}

	var records []*model.Auth
	if err := rds.TestDB(ctx).Model(&model.Auth{}).
		Where("app_id = ? and username = ? and auth_res_id = ?",
			rq.AppId, rq.Username, rq.AuthResId).Find(&records).Error; err != nil {
		return nil, errors.WithMessage(err, "db find")
	}
	if len(records) > 1 {
		err := errors.New("records num bigger than 1")
		log.Error("Check.records.num", err, "rq", rq, "records", records)
		return nil, errs.Internal
	}
	resp := new(model.CheckResp)
	if len(records) == 0 {
		resp.Reason = auth.NotGranted
		return resp, nil
	}
	expiredAt := records[0].ExpiredAt
	if expiredAt != -1 && time.Now().Unix() > expiredAt {
		resp.Reason = auth.Expired
		return resp, nil
	}
	resp.Pass = true
	return resp, nil
}
