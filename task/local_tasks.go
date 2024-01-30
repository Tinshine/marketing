package task

import "context"

type deductQuota struct {
	Id uint
}

func NewDeductQuota(quotaId uint) *deductQuota {
	return &deductQuota{quotaId}
}

func (d *deductQuota) GetId() uint {
	return d.Id
}
func (d *deductQuota) Try(ctx context.Context, p *Params) (*Resp, error) {
	// todo...
	return &Resp{}, nil
}

func (d *deductQuota) Cancel(ctx context.Context, p *Params) (*Resp, error) {
	// todo...
	return &Resp{}, nil
}

func (d *deductQuota) Confirm(ctx context.Context, p *Params) (*Resp, error) {
	// todo...
	return &Resp{}, nil
}
