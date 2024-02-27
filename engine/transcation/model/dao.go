package model

import (
	"context"
	"marketing/consts"
	"marketing/consts/errs"
	"marketing/database/rds"
	"marketing/util"
	"marketing/util/log"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DAO interface {
	SetEnv(env consts.Env) DAO
	CreateTx(ctx context.Context, records []*Transaction) error
	CancelTx(ctx context.Context, trId string) error
	ConfirmTx(ctx context.Context, trId string) error
}

func InitDAO() DAO {
	if util.IsUnitTest() {
		return initMockDAO()
	}
	return &rdsDAO{}
}

type rdsDAO struct {
	env consts.Env
}

func (dao *rdsDAO) SetEnv(env consts.Env) DAO {
	dao.env = env
	return dao
}

func (dao *rdsDAO) CreateTx(ctx context.Context, records []*Transaction) error {
	return rds.DB(ctx, dao.env).Create(&records).Error
}

func (dao *rdsDAO) CancelTx(ctx context.Context, trId string) error {
	return rds.DB(ctx, dao.env).Transaction(func(tx *gorm.DB) error {
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

func (dao *rdsDAO) ConfirmTx(ctx context.Context, trId string) error {
	return rds.DB(ctx, dao.env).Transaction(func(tx *gorm.DB) error {
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
