package model

import (
	"context"
	"marketing/consts"
	"marketing/consts/errs"
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

type mockDAO struct {
	ods map[uint]*Order
}

func initTestOrderData() map[uint]*Order {
	ods := make(map[uint]*Order)
	ods[1] = &Order{
		Id:        1,
		AppId:     1,
		OrderId:   "1",
		TxId:      "1",
		GroupId:   1,
		UserId:    "1",
		TxState:   consts.StateTry,
		CreatedAt: time.Now(),
	}
	ods[2] = &Order{
		Id:        2,
		AppId:     2,
		OrderId:   "2",
		GroupId:   2,
		UserId:    "2",
		TxState:   consts.StateTry,
		CreatedAt: time.Now(),
	}
	return ods
}

func (dao *mockDAO) SetEnv(env consts.Env) DAO {
	if dao.ods == nil {
		dao.ods = initTestOrderData()
	}
	return dao
}

func (dao *mockDAO) FindOrder(c context.Context, txId, userId string, groupId uint) (*Order, error) {
	for _, od := range dao.ods {
		if od.TxId == txId && od.UserId == userId && od.GroupId == groupId {
			return od, nil
		}
	}
	return nil, errs.OrderNotFound
}

func (dao *mockDAO) CreateOrder(c context.Context,
	txId, userId string, appId, groupId uint, state consts.TxState) (*Order, error) {
	order, err := reqToModel(c, txId, userId, appId, groupId, state)
	if err != nil {
		return nil, errors.WithMessage(err, "req to model")
	}
	order.Id = uint(rand.Intn(1000000) + 1000000)
	order.CreatedAt = time.Now()
	dao.ods[order.Id] = order
	return order, nil
}

func (dao *mockDAO) UpdateOrder(c context.Context, id uint, src, dest consts.TxState) error {
	for _, od := range dao.ods {
		if od.Id == id && od.TxState == src {
			od.TxState = dest
			return nil
		}
	}
	return nil
}
