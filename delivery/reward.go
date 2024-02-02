package delivery

import (
	"context"
	"marketing/delivery/model"
	trM "marketing/engine/transcation/model"
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

func (r *reward) Try(ctx context.Context, params *trM.Params) (*trM.Resp, error) {
	if err := checkParams(ctx, params); err != nil {
		return nil, err
	}

	// todo...
	model.FirstOrInitOrder(ctx, params.Reward["group_id"].(int64), params.User.GetId(), r.TrId, params.Ev)

	return nil, nil
}
func (r *reward) Cancel(ctx context.Context, params *trM.Params) (*trM.Resp, error) {
	// todo...
	return nil, nil
}
func (r *reward) Confirm(ctx context.Context, params *trM.Params) (*trM.Resp, error) {
	// todo...
	return nil, nil
}
