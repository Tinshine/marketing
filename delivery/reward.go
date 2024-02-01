package delivery

import (
	"context"
	"marketing/engine/transcation/model"
)

type reward struct {
	TrId uint
}

func NewReward() *reward {
	return &reward{}
}

func (r *reward) SetTrId(trId uint) {
	r.TrId = trId
}
func (r *reward) GetTrId() uint {
	return r.TrId
}
func (r *reward) Try(ctx context.Context, params *model.Params) (*model.Resp, error) {
	if err := checkParams(ctx, params); err != nil {
		return nil, err
	}

	return nil, nil
}
func (r *reward) Cancel(ctx context.Context, params *model.Params) (*model.Resp, error) {
	// todo...
	return nil, nil
}
func (r *reward) Confirm(ctx context.Context, params *model.Params) (*model.Resp, error) {
	// todo...
	return nil, nil
}
