package transaction

import (
	"context"
	"marketing/consts"
	"marketing/database/rds"
	"marketing/task"
	"marketing/util/idgen"
	"marketing/util/log"

	"github.com/pkg/errors"
)

type tr struct {
	Tasks    []task.T
	TryResps map[uint]*task.Resp
	Ev       consts.Env
	Txs      map[uint]*Transaction
}

func NewTr(tasks []task.T, ev consts.Env) *tr {
	return &tr{
		Tasks: tasks,
		Ev:    ev,
	}
}

func (t *tr) newTransaction(ctx context.Context) error {
	txId, err := idgen.Gen(ctx)
	if err != nil {
		return errors.WithMessage(err, "idgen gen")
	}
	var trans []*Transaction
	for _, tsk := range t.Tasks {
		trans = append(trans, &Transaction{
			State:  consts.StateTry,
			TaskId: tsk.GetId(),
			TxId:   txId,
		})
	}
	if err := rds.DB(ctx, t.Ev).Create(&trans).Error; err != nil {
		return errors.WithMessage(err, "create")
	}
	txs := make(map[uint]*Transaction, len(trans))
	for i := range trans {
		txs[trans[i].TaskId] = trans[i]
	}
	t.Txs = txs
	return nil
}

func (t *tr) Execute(ctx context.Context, params *task.Params) error {
	var err error
	defer func() {
		if err != nil {
			t.Rollback(ctx, params)
		} else {
			t.Commit(ctx, params)
		}
	}()

	if err = t.newTransaction(ctx); err != nil {
		return errors.WithMessage(err, "new transaction")
	}

	if err = t.Try(ctx, params); err != nil {
		return errors.WithMessage(err, "try transaction")
	}
	return nil
}

func (t *tr) Commit(ctx context.Context, params *task.Params) {
	// todo.. parallel confirm
}

func (t *tr) Rollback(ctx context.Context, params *task.Params) {
	if len(t.TryResps) == 0 {
		return
	}
	for i := range t.Tasks {
		i := i
		go func() {
			tsk := t.Tasks[i]
			tx := t.Txs[tsk.GetId()]
			var err error
			defer func() {
				if err != nil {
					retryCancel <- tx
				}
			}()
			if err = cancelTransaction(ctx, tx, t.Ev); err != nil {
				log.Error("Rollback.cancelTransaction", err, "task", tsk.GetId())
				return
			}
			if _, ok := t.TryResps[tsk.GetId()]; ok {
				var resp *task.Resp
				resp, err = tsk.Cancel(ctx, params)
				if err != nil {
					log.Error("Rollback.Cancel.Error", err, "tx", tx)
					return
				}
				log.Info("Rollback.Cancel.Success", "tx", tx, "resp", resp)
			}
			// no try, no need to cancel
		}()

	}
}

func (t *tr) Try(ctx context.Context, params *task.Params) error {
	resps := make(map[uint]*task.Resp, len(t.Tasks))
	defer func() {
		t.TryResps = resps
	}()

	for _, task := range t.Tasks {
		resp, err := task.Try(ctx, params)
		if err != nil {
			return err
		}
		resps[task.GetId()] = resp
	}
	return nil
}

// todo...
var retryCancel chan *Transaction
