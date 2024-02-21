package tr

import (
	"context"
	"marketing/engine/transcation/model"
	tskM "marketing/task/model"
)

type httpTr struct {
	TrId string
	*tskM.Task
}

func (h *httpTr) GetTxId() string {
	return h.TrId
}

func (h *httpTr) Try(ctx context.Context, p *model.Params) (*model.Resp, error) {
	// todo...
	return &model.Resp{}, nil
}

func (h *httpTr) Cancel(ctx context.Context, p *model.Params) (*model.Resp, error) {
	// todo...
	return &model.Resp{}, nil
}

func (h *httpTr) Confirm(ctx context.Context, p *model.Params) (*model.Resp, error) {
	// todo...
	return &model.Resp{}, nil
}
