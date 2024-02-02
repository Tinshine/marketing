package model

import (
	"marketing/common"
	"marketing/consts"
)

type Order struct {
	common.Model
	OrderId string         `gorm:"order_id"`
	TxId    uint           `gorm:"tx_id"`
	GroupId int64          `gorm:"group_id"` // gift group_id
	UserId  string         `gorm:"user_id"`
	TrState consts.TrState `gorm:"tr_state"`
}

func (o *Order) TableName() string {
	return "order"
}

type RewardReq struct{}
