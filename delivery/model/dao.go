package model

import (
	"context"
	"fmt"
	"marketing/consts"
	"marketing/consts/errs"
	"marketing/database/rds"
	"math/rand"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func FindOrder(c context.Context, r *RewardReq) (*Order, error) {
	var order Order
	if err := rds.DB(c, r.Ev).
		Where("tx_id = ? and user_id = ? and group_id = ?", r.TxId, r.UserId, r.GroupId).
		Find(&order).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, errors.WithMessage(errs.Internal, err.Error())
		}
		return nil, errs.OrderNotFound
	}
	return &order, nil
}

func CreateOrder(c context.Context, req *RewardReq, state consts.TxState) (*Order, error) {
	order := Order{}
	order.AppId = req.AppId
	order.GroupId = req.GroupId
	order.UserId = req.UserId
	order.TxId = req.TxId
	order.OrderId = genOrderId(req)
	order.TxState = consts.StateTry
	if err := rds.DB(c, req.Ev).Create(&order).Error; err != nil {
		return nil, errors.WithMessage(errs.Internal, err.Error())
	}
	return &order, nil
}

func UpdateOrder(c context.Context, id uint, ev consts.Env, src, dest consts.TxState) error {
	return rds.DB(c, ev).
		Where("id = ? and tx_state = ?", id, src).
		UpdateColumn("tx_state", dest).Error
}

func genOrderId(req *RewardReq) string {
	now := time.Now().UnixNano()
	sb := strings.Builder{}
	for i := 0; i < 4; i++ {
		n := rand.Intn(26)
		sb.WriteByte(byte('A' + n))
	}
	return fmt.Sprintf("%d-%d-%s-%s", now, req.AppId, req.UserId, sb.String())
}
