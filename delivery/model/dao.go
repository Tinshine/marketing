package model

import (
	"context"
	"marketing/consts"
	"marketing/consts/errs"
	"marketing/database/rds"
	"marketing/util"
	"marketing/util/idgen"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type DAO interface {
	SetEnv(env consts.Env) DAO
	FindOrder(c context.Context, txId, userId string, groupId uint) (*Order, error)
	CreateOrder(c context.Context, txId, userId string, appId, groupId uint, state consts.TxState) (*Order, error)
	UpdateOrder(c context.Context, id uint, src, dest consts.TxState) error
}

type rdsDAO struct {
	env consts.Env
}

func InitDAO() DAO {
	if util.IsUnitTest() {
		return &mockDAO{}
	}
	return &rdsDAO{}
}

func (dao *rdsDAO) SetEnv(env consts.Env) DAO {
	dao.env = env
	return dao
}

func (dao *rdsDAO) FindOrder(c context.Context, txId, userId string, groupId uint) (*Order, error) {
	var order Order
	if err := rds.DB(c, dao.env).
		Where("tx_id = ? and user_id = ? and group_id = ?", txId, userId, groupId).
		First(&order).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, errors.WithMessage(errs.Internal, err.Error())
		}
		return nil, errs.OrderNotFound
	}
	return &order, nil
}

func (dao *rdsDAO) CreateOrder(c context.Context,
	txId, userId string, appId, groupId uint, state consts.TxState) (*Order, error) {
	order, err := reqToModel(c, txId, userId, appId, groupId, state)
	if err != nil {
		return nil, errors.WithMessage(err, "req to model")
	}
	if err := rds.DB(c, dao.env).Create(order).Error; err != nil {
		return nil, errors.WithMessage(errs.Internal, err.Error())
	}
	return order, nil
}

func reqToModel(c context.Context, txId, userId string, appId, groupId uint, state consts.TxState) (*Order, error) {
	orderId, err := idgen.Gen(c)
	if err != nil {
		return nil, errors.WithMessage(err, "idgen gen")
	}
	order := Order{}
	order.AppId = appId
	order.GroupId = groupId
	order.UserId = userId
	order.TxId = txId
	order.OrderId = orderId
	order.TxState = consts.StateTry
	return &order, nil
}

func (dao *rdsDAO) UpdateOrder(c context.Context, id uint, src, dest consts.TxState) error {
	return rds.DB(c, dao.env).Model(&Order{}).
		Where("id = ? and tx_state = ?", id, src).
		UpdateColumn("tx_state", dest).Error
}
