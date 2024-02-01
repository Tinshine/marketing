package transaction

import (
	"context"
	"marketing/consts"
	"marketing/database/rds"
	"marketing/engine/transcation/model"
	"marketing/engine/transcation/tr"
	tsM "marketing/task/model"
	"marketing/util/idgen"
	"marketing/util/log"

	"github.com/pkg/errors"
)

type tx struct {
	TryResps map[uint]*model.Resp
	Ev       consts.Env
	Trs      []model.T
	TxId     string
}

func NewTx(ev consts.Env) *tx {
	return &tx{
		Ev: ev,
	}
}

func (t *tx) newTransaction(ctx context.Context, tasks []*tsM.Task) error {
	// make map: task_id -> Tr
	trMap := make(map[uint]model.T, len(tasks))
	for _, tsk := range tasks {
		trMap[tsk.ID] = tr.NewTr(tsk)
	}

	txId, err := idgen.Gen(ctx)
	if err != nil {
		return errors.WithMessage(err, "idgen gen")
	}
	var ts []*model.Transaction
	for _, tsk := range tasks {
		ts = append(ts, &model.Transaction{
			State:  consts.StateTry,
			TaskId: tsk.ID,
			TxId:   txId,
		})
	}
	if err := rds.DB(ctx, t.Ev).Create(&ts).Error; err != nil {
		return errors.WithMessage(err, "create")
	}

	t.Trs = make([]model.T, 0, len(trMap))
	// traverse map to invoke each Tr's SetTrId func
	// ti.ID is the tr_id corresponding to ti.TaskId
	for _, ti := range ts {
		tr := trMap[ti.TaskId]
		tr.SetTrId(ti.ID)
		t.Trs = append(t.Trs, tr)
	}
	t.TxId = txId
	return nil
}

func (t *tx) Execute(ctx context.Context, tasks []*tsM.Task, params *model.Params) error {
	var err error
	defer func() {
		if err != nil {
			t.Rollback(ctx, params)
		} else {
			t.Commit(ctx, params)
		}
	}()

	if err = t.newTransaction(ctx, tasks); err != nil {
		return errors.WithMessage(err, "new transaction")
	}

	if err = t.Try(ctx, params); err != nil {
		return errors.WithMessage(err, "try transaction")
	}
	return nil
}

func (t *tx) Commit(ctx context.Context, params *model.Params) {
	for i := range t.Trs {
		i := i
		go func() {
			var err error
			defer func() {
				if err != nil {
					retryConfirm <- t.Trs[i]
				}
			}()
			var resp *model.Resp
			resp, err = t.Trs[i].Confirm(ctx, params)
			if err != nil {
				log.Error("Commit.confirm.Error", err, "trId", t.Trs[i].GetTrId())
				return
			}
			log.Info("Commit.confirm.success", "trId", t.Trs[i].GetTrId(), "resp", resp)
			// after all have been confirmed, update db,
			// if confirm success, but update failed, retry will be ok.
			if err := model.ConfirmTx(ctx, t.Trs[i].GetTrId(), t.Ev); err != nil {
				log.Error("Commit.confirmTx", err, "trId", t.Trs[i].GetTrId())
				return
			}
		}()
	}
}

func (t *tx) Rollback(ctx context.Context, params *model.Params) {
	if len(t.TryResps) == 0 {
		return
	}
	for i := range t.Trs {
		i := i
		if _, ok := t.TryResps[t.Trs[i].GetTrId()]; !ok {
			continue
		}
		go func() {
			var err error
			defer func() {
				if err != nil {
					retryCancel <- t.Trs[i]
				}
			}()
			var resp *model.Resp
			resp, err = t.Trs[i].Cancel(ctx, params)
			if err != nil {
				log.Error("Rollback.Cancel.Error", err, "tx", t)
				return
			}
			log.Info("Rollback.Cancel.Success", "tx", t, "resp", resp)
			// after all have been cancelled success, update db,
			// if cancellation success, but update failed, retry with be ok.
			if err = model.CancelTx(ctx, t.Trs[i].GetTrId(), t.Ev); err != nil {
				log.Error("Rollback.cancelTx", err, "trId", t.Trs[i].GetTrId())
				return
			}
		}()

	}
}

func (t *tx) Try(ctx context.Context, params *model.Params) error {
	resps := make(map[uint]*model.Resp, len(t.Trs))
	defer func() {
		t.TryResps = resps
	}()

	for _, tr := range t.Trs {
		resp, err := tr.Try(ctx, params)
		if err != nil {
			log.Error("Try.Try.Error", err, "tx", tr)
			return err
		}
		log.Info("Try.Try.Success", "tx", tr, "resp", resp)
		resps[tr.GetTrId()] = resp
	}

	log.Info("Try.All.Success", "tx", t, "resp")
	return nil
}

// todo...
var retryCancel chan model.T
var retryConfirm chan model.T
