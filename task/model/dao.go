package model

import (
	"marketing/consts/errs"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func FindById(db *gorm.DB, id int) (*Task, error) {
	var t *Task
	if err := db.First(&t, id).Error; err != nil {
		return nil, errors.WithMessage(err, "find by id")
	}
	return t, nil
}

func FindByIds(db *gorm.DB, ids []uint) ([]*Task, error) {
	var ts []*Task
	if err := db.Model(&Task{}).
		Where("id in ?", ids).Find(&ts).Error; err != nil {
		return nil, errors.WithMessage(err, "find by ids")
	}
	if len(ts) != len(ids) {
		return nil, errs.BatchQueryLenNotMatch
	}
	return ts, nil
}
