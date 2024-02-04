package model

import (
	"context"
	"marketing/consts"
	"marketing/consts/errs"
	"marketing/database/rds"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func FindOrder(c context.Context, groupId int64, userId string, txId uint, ev consts.Env) (*Order, error) {
	var order Order
	if err := rds.DB(c, ev).
		Where("tx_id = ? and user_id = ? and group_id = ?", txId, userId, groupId).
		Find(&order).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, errors.WithMessage(errs.Internal, err.Error())
		}
		return nil, errs.OrderNotFound
	}
	return &order, nil
}

func FirstOrInitOrder(c context.Context, grougId int64, userId string, txId uint, ev consts.Env) (*Order, error) {
	var order Order
	if err := rds.DB(c, ev).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("tx_id = ? and user_id = ? and group_id = ?", txId, userId, grougId).
			First(&order).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return errors.WithMessage(errs.Internal, err.Error())
			}
			if err == nil {
				return nil
			}
			order = Order{
				UserId:  userId,
				TxId:    txId,
				GroupId: grougId,
				TrState: consts.StateTry,
			}
			if err := tx.Create(&order).Error; err != nil {
				return errors.WithStack(err)
			}
		}
		return nil
	}); err != nil {
		return nil, errors.WithMessage(err, "transaction")
	}
	return &order, nil
}
