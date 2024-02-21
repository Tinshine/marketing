package reward

import (
	"context"
	trM "marketing/engine/transcation/model"

	"github.com/pkg/errors"
)

type reward struct {
	TxId string
}

func NewReward(txId string) *reward {
	return &reward{txId}
}

func (r *reward) GetTxId() string { return r.TxId }

func (r *reward) Try(ctx context.Context, params *trM.Params) (*trM.Resp, error) {
	req, err := parseParams(r.TxId, params)
	if err != nil {
		return nil, errors.WithMessage(err, "parse params")
	}
	resp := new(trM.Resp)

	if err := tryOrder(ctx, req); err != nil {
		return nil, errors.WithMessage(err, "try order")
	}
	return resp, nil
}
func (r *reward) Cancel(ctx context.Context, params *trM.Params) (*trM.Resp, error) {
	req, err := parseParams(r.TxId, params)
	if err != nil {
		return nil, errors.WithMessage(err, "parse params")
	}
	resp := new(trM.Resp)

	if err := cancelOrder(ctx, req); err != nil {
		return nil, errors.WithMessage(err, "cancel order")
	}
	return resp, nil
}
func (r *reward) Confirm(ctx context.Context, params *trM.Params) (*trM.Resp, error) {
	req, err := parseParams(r.TxId, params)
	if err != nil {
		return nil, errors.WithMessage(err, "parse params")
	}
	resp := new(trM.Resp)

	if err := confirmOrder(ctx, req); err != nil {
		return nil, errors.WithMessage(err, "confirm order")
	}

	if err := delivery(ctx, req); err != nil {
		return nil, errors.WithMessage(err, "delivery")
	}
	return resp, nil
}
