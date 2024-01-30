package rule

import (
	"context"
	"marketing/consts"
	"marketing/consts/engine"
	"marketing/consts/errs"
	"marketing/consts/gormtype"
	transaction "marketing/engine/transcation"
	"marketing/task"
	tModel "marketing/task/model"
	tService "marketing/task/service"

	"github.com/pkg/errors"
)

type Resp struct{}

type R struct {
	QuotaId uint                 `gorm:"column:"quota_id"`
	TaskIds gormtype.Slice[uint] `gorm:"column:"task_ids"` // transactions' ids
	Env     consts.Env
}

func (r *R) Execute(ctx context.Context) (*Resp, error) {
	err := r.checkQuota(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "get quota")
	}

	ok, err := r.checkConstraints(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "check constraints")
	}
	if !ok {
		return nil, errors.WithMessage(errs.ConstraintsNotMeet, "quota limit")
	}

	res, err := r.do(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "run")
	}

	return res, nil
}

func (r *R) checkQuota(ctx context.Context) error {
	// todo...
	return nil
}

func (r *R) checkConstraints(ctx context.Context) (bool, error) {
	// todo...
	return false, nil
}

func (r *R) getTasks(ctx context.Context) ([]task.T, error) {
	tasks, err := tService.BatchQuery(ctx, &tModel.BatchReq{
		Ids: r.TaskIds,
		Ev:  r.Env,
	})
	if err != nil {
		return nil, errors.WithMessage(err, "task service batch query")
	}
	t := make([]task.T, 0, len(tasks)+1)
	t = append(t, task.NewDeductQuota(r.QuotaId))
	for i := range tasks {
		if tasks[i].Type == engine.TaskType_HTTP {
			t = append(t, task.NewHTTPTask(tasks[i]))
			continue
		}
		if tasks[i].Type == engine.TaskType_RPC {
			t = append(t, task.GetRPCTask(tasks[i].Handler))
			continue
		}
	}
	return t, nil
}

func (r *R) do(ctx context.Context) (*Resp, error) {
	ts, err := r.getTasks(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "get tasks")
	}

	// todo.. need lock here to prevent concurrent execute
	if err := transaction.NewTr(ts, r.Env).Execute(ctx, &task.Params{}); err != nil {
		return nil, errors.WithMessage(err, "execute transaction")
	}

	return nil, nil
}
