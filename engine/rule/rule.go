package rule

import (
	"context"
	"marketing/consts/errs"
	qt "marketing/manager/quota/model"

	"github.com/pkg/errors"
)

type R struct {
}

func (r *R) Execute(ctx context.Context) (*Resp, error) {
	q, err := r.getQuota(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "get quota")
	}
	if !q.Avaiable() {
		return nil, errors.WithMessage(errs.QuotaLimit, "quota limit")
	}

	ok, err := r.checkConstraints(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "check constraints")
	}
	if !ok {
		return nil, errors.WithMessage(errs.ConstraintsNotMeet, "quota limit")
	}

	res, err := r.run(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "run")
	}

	return res, nil
}

func (r *R) getQuota(ctx context.Context) (*qt.Quota, error) {
	// todo...
	return nil, nil
}

func (r *R) checkConstraints(ctx context.Context) (bool, error) {
	// todo...
	return false, nil
}

func (r *R) run(ctx context.Context) (*Resp, error) {
	// todo...
	return nil, nil
}

type Resp struct{}
