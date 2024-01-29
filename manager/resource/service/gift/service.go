package gift

import (
	"context"
	"marketing/consts"
	"marketing/consts/errs"
	rConst "marketing/consts/resource"
	"marketing/database/rds"
	"marketing/database/redis"
	model "marketing/manager/resource/model/gift"
	"marketing/util"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Query(ctx context.Context, rq *model.QueryReq) (*model.QueryResp, error) {
	if err := rq.Validate(); err != nil {
		return nil, errors.WithMessage(err, "validate")
	}

	resp := new(model.QueryResp)
	if rq.Id != nil {
		gift, err := model.FindById(rds.DB(ctx, rq.Env), *rq.Id)
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
		gifts, err := model.FindByGroupId(rds.DB(ctx, rq.Env), rq.AppId, *rq.GroupId)
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

	gifts, err := model.FindByAppId(rds.DB(ctx, rq.Env), rq.AppId)
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

	if err := rds.TestDB(ctx).Create(&gift).Error; err != nil {
		return nil, errors.WithMessage(err, "create")
	}
	resp := new(model.AddResp)
	resp.Id = gift.Id
	return resp, nil
}

func Update(ctx context.Context, rq *model.UpdateReq) error {
	if err := rq.Validate(); err != nil {
		return errors.WithMessage(err, "validate")
	}

	username, err := util.GetUsername(ctx)
	if err != nil {
		return errors.WithMessage(err, "get username")
	}

	fields := map[string]interface{}{
		"updated_by": username,
	}
	if rq.Emails != nil {
		fields["emails"] = rq.Emails
	}
	if rq.GiftName != nil {
		fields["gift_name"] = rq.GiftName
	}
	if rq.Items != nil {
		fields["items"] = rq.Items
	}
	if rq.LotteryRate != nil {
		fields["lottery_rate"] = rq.LotteryRate
	}

	if err := model.UpdateById(rds.TestDB(ctx), rq.Id, fields); err != nil {
		return errors.WithMessage(err, "update by id")
	}
	return nil
}

func Sync(ctx context.Context, rq *model.SyncReq) error {
	if err := rq.Validate(); err != nil {
		return errors.WithMessage(err, "validate")
	}
	username, err := util.GetUsername(ctx)
	if err != nil {
		return errors.WithMessage(err, "get username")
	}

	key := util.MakeKey(rConst.RedisPrefixSync, rq.Id)
	locked, lockV, err := redis.Lock(ctx, key, time.Second*3)
	if err != nil {
		return errors.WithMessage(err, "redis lock")
	}
	defer redis.Unlock(ctx, key, locked, lockV)
	if !locked {
		return errors.WithMessage(errs.TooManyRequests, "redis lock failed")
	}

	gift, err := model.FindById(rds.TestDB(ctx), int(rq.Id))
	if err != nil {
		return errors.WithMessage(err, "find by id")
	}

	newGift := &model.Gift{
		Id:          gift.Id,
		AppId:       gift.AppId,
		GiftType:    gift.GiftType,
		GiftName:    gift.GiftName,
		LotteryRate: gift.LotteryRate,
		GroupId:     gift.GroupId,
		Items:       gift.Items,
		Emails:      gift.Emails,
		State:       consts.StateCreated,
		CreatedBy:   username,
		CreatedAt:   time.Now().Unix(),
	}
	if err := rds.ProdDB(ctx).Create(&newGift).Error; err != nil {
		return errors.WithMessage(err, "create")
	}

	return nil
}

func Delete(ctx context.Context, rq *model.DeleteReq) error {
	if err := rq.Validate(); err != nil {
		return errors.WithMessage(err, "validate")
	}

	if rq.Env == consts.Test {
		_, err := model.FindById(rds.ProdDB(ctx), int(rq.Id))
		if err != nil && err != gorm.ErrRecordNotFound {
			return errors.WithMessage(err, "find prod by id")
		}
		if err == nil {
			return errors.WithMessage(errs.DeleteNotAllowed,
				"can't delete test conf before prod conf deleted")
		}
		// it's ok that prod record not found
	}

	gift, err := model.FindById(rds.DB(ctx, rq.Env), int(rq.Id))
	if err != nil {
		return errors.WithMessage(err, "find by id")
	}
	if gift.State == consts.StateCreated || gift.State == consts.StateOffline {
		if err := rds.DB(ctx, rq.Env).Delete(&gift).Error; err != nil {
			return errors.WithMessage(err, "delete")
		}
		return nil
	}

	return errors.WithMessage(errs.DeleteNotAllowed, "conf in such state can't be deleted")
}

func Release(ctx context.Context, rq *model.ReleaseReq) error {
	if err := rq.Validate(); err != nil {
		return errors.WithMessage(err, "validate")
	}

	username, err := util.GetUsername(ctx)
	if err != nil {
		return errors.WithMessage(err, "get username")
	}

	return rds.DB(ctx, rq.Env).Transaction(func(tx *gorm.DB) error {
		var gift model.Gift
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&gift, rq.Id).Error; err != nil {
			return errors.WithMessage(err, "db find")
		}
		if gift.State != consts.StateCreated {
			return errs.InvalidState
		}
		if err := tx.Model(&gift).Updates(map[string]interface{}{
			"updated_by": username,
			"updated_at": time.Now().Unix(),
			"state":      consts.StateOnline,
		}).Error; err != nil {
			return errors.WithMessage(err, "update by id")
		}
		return nil
	})
}
