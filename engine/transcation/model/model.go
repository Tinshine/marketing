package model

import (
	"marketing/consts"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	State  consts.TransactionState `gorm:"column:state"`
	TaskId uint                    `gorm:"column:task_id"`
	TxId   string                  `gorm:"column:tx_id"`
}

func (t *Transaction) TableName() string {
	return "transaction"
}
