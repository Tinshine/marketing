package reward

import (
	"context"
	"marketing/consts/errs"
	"marketing/delivery/model"
	trM "marketing/engine/transcation/model"
	"marketing/util/log"

	"github.com/pkg/errors"
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

func (r *reward) Try(ctx context.Context, params *trM.Params) (resp *trM.Resp, err error) {
	resp = new(trM.Resp)
	defer func() {
		if err == nil {
			// do...
		} else {
			// todo...
		}
	}()
	req, err := parseParams(params)
	if err != nil {
		return nil, err
	}

	if err := checkLimit(ctx, req); err != nil {
		return nil, errors.WithMessage(err, "check limit")
	}

	order, err := model.FindOrder(ctx, params.Input["group_id"].(int64), params.User.GetId(), r.TrId, params.Ev)
	if err != nil && err != errs.OrderNotFound {
		return nil, errors.WithMessage(err, "find order")
	}
	if err == nil {
		// a duplicate order exists, should ignore this in case cancel or confirm has been performed
		log.Error("Try.FindOrder.Exist", errs.DuplicatedTry, "params", params, "order", order)
		return nil, nil
	}

	// todo...

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
