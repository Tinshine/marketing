package tr

import (
	"context"
	"marketing/engine/transcation/model"
)

type deductQuota struct {
	TrId uint
}

func (d *deductQuota) SetTrId(trId uint) {
	d.TrId = trId
}

func (d *deductQuota) GetTrId() uint {
	return d.TrId
}
func (d *deductQuota) Try(ctx context.Context, p *model.Params) (*model.Resp, error) {
	// todo...
	return &model.Resp{}, nil
}

func (d *deductQuota) Cancel(ctx context.Context, p *model.Params) (*model.Resp, error) {
	// todo...
	return &model.Resp{}, nil
}

func (d *deductQuota) Confirm(ctx context.Context, p *model.Params) (*model.Resp, error) {
	// todo...
	return &model.Resp{}, nil
}
