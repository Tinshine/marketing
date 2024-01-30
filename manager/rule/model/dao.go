package model

import (
	"marketing/consts/errs"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func FindByRuleId(db *gorm.DB, ruleId string) (*Rule, error) {
	var rule []Rule
	if err := db.Model(&Rule{}).
		Where("rule_id =?", ruleId).Find(&rule).Error; err != nil {
		return nil, errors.WithMessage(err, "db find")
	}
	if len(rule) > 1 {
		return nil, errs.DuplicateRule
	}
	if len(rule) == 0 {
		return nil, errs.ConfNotFound
	}
	return &rule[0], nil
}
