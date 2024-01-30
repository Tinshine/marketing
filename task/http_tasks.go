package task

import (
	"context"
	"marketing/task/model"
)

type HttpTask struct {
	Id          uint
	Domain      string
	TryPath     string
	CancelPath  string
	ConfirmPath string
}

func NewHTTPTask(task *model.Task) *HttpTask {
	return &HttpTask{
		Id:          task.ID,
		Domain:      task.Domain,
		TryPath:     task.TryPath,
		CancelPath:  task.CancelPath,
		ConfirmPath: task.ConfirmPath,
	}
}

func (h *HttpTask) GetId() uint {
	return h.Id
}

func (h *HttpTask) Try(ctx context.Context, p *Params) (*Resp, error) {
	// todo...
	return &Resp{}, nil
}

func (h *HttpTask) Cancel(ctx context.Context, p *Params) (*Resp, error) {
	// todo...
	return &Resp{}, nil
}

func (h *HttpTask) Confirm(ctx context.Context, p *Params) (*Resp, error) {
	// todo...
	return &Resp{}, nil
}
