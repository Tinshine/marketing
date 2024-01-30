package service

import (
	"context"
	"marketing/database/rds"
	"marketing/task/model"

	"github.com/pkg/errors"
)

func Query(ctx context.Context, rq *model.QueryReq) (*model.QueryResp, error) {
	if err := rq.Validate(); err != nil {
		return nil, errors.WithMessage(err, "validate")
	}

	if rq.Id != 0 {
		t, err := model.FindById(rds.DB(ctx, rq.Ev), int(rq.Id))
		if err != nil {
			return nil, errors.WithMessage(err, "find by id")
		}
		return &model.QueryResp{Task: *t}, nil
	}

	return nil, nil
}

func BatchQuery(ctx context.Context, rq *model.BatchReq) ([]*model.Task, error) {
	if err := rq.Validate(); err != nil {
		return nil, errors.WithMessage(err, "validate")
	}

	if len(rq.Ids) == 0 {
		return nil, nil
	}

	return model.FindByIds(rds.DB(ctx, rq.Ev), rq.Ids)
}
