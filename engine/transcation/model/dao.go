package model

import (
	"context"
	"marketing/consts"
	"marketing/consts/errs"
	"marketing/database/rds"
	"marketing/util/log"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CancelTx(ctx context.Context, trId uint, ev consts.Env) error {
	return rds.DB(ctx, ev).Transaction(func(tx *gorm.DB) error {
		var r Transaction
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&r, trId).Error; err != nil {
			return errors.WithMessage(err, "db find")
		}
		if r.State == consts.StateCancel {
			log.Warn("cancelTransaction.duplicate", "record", r)
			return nil
		}
		if r.State == consts.StateConfirm {
			err := errs.TransactionConfirmed
			log.Error("cancelTransaction.confirmed", err, "record", r)
			return err
		}
		return tx.Model(&r).UpdateColumn("state", consts.StateCancel).Error
	})
}

func ConfirmTx(ctx context.Context, trId uint, ev consts.Env) error {
	return rds.DB(ctx, ev).Transaction(func(tx *gorm.DB) error {
		var r Transaction
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&r, trId).Error; err != nil {
			return errors.WithMessage(err, "db find")
		}
		if r.State == consts.StateConfirm {
			log.Warn("confirmTransaction.duplicate", "record", r)
			return nil
		}
		if r.State == consts.StateCancel {
			err := errs.TransactionCanceled
			log.Error("confirmTransaction.canceled", err, "record", r)
			return err
		}
		return tx.Model(&r).UpdateColumn("state", consts.StateConfirm).Error
	})
}
