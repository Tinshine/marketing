package model

import (
	"errors"
	"marketing/consts"
	"marketing/consts/engine"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string        `gorm:"column:"name"`
	Type        engine.TrType `gorm:"column:"type"`
	Domain      string        `gorm:"column:"domain"`
	TryPath     string        `gorm:"column:"try_path"`
	CancelPath  string        `gorm:"column:"cancel_path"`
	ConfirmPath string        `gorm:"column:"confirm_path"`
	Handler     string        `gorm:"column:"handler"` // rpc handler name
}

func (t *Task) TableName() string {
	return "task"
}

type QueryReq struct {
	Id uint       `json:"id"`
	Ev consts.Env `json:"ev"`
}

func (q *QueryReq) Validate() error {
	if q.Ev != consts.Test && q.Ev != consts.Prod {
		return errors.New("env is invalid")
	}
	return nil
}

type QueryResp struct {
	Task
}

type BatchReq struct {
	Ids []uint     `json:"ids"`
	Ev  consts.Env `json:"ev"`
}

func (q *BatchReq) Validate() error {
	for _, id := range q.Ids {
		if id == 0 {
			return errors.New("id is invalid")
		}
	}
	if q.Ev != consts.Test && q.Ev != consts.Prod {
		return errors.New("env is invalid")
	}
	return nil
}
