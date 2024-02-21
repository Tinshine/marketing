package model

import (
	"marketing/consts"
	"time"
)

type Order struct {
	Id        uint           `gorm:"id"`
	AppId     uint           `gorm:"app_id"`
	OrderId   string         `gorm:"order_id"`
	TxId      string         `gorm:"tx_id"`
	GroupId   uint           `gorm:"group_id"` // gift group_id
	UserId    string         `gorm:"user_id"`
	TxState   consts.TxState `gorm:"tx_state"`
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at"`
}

func (o *Order) TableName() string {
	return "order"
}

type RewardReq struct {
	Ev      consts.Env
	UserId  string
	QuotaId uint
	TxId    string
	AppId   uint
	GroupId uint
}
