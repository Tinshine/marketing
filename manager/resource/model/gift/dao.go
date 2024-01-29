package gift

import (
	"context"
	"marketing/consts/resource"
	"marketing/database/rds"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func FindById(db *gorm.DB, id int) (*Gift, error) {
	var gift Gift
	if err := db.First(&gift, id).Error; err != nil {
		return nil, errors.WithMessage(err, "db first")
	}
	return &gift, nil
}

func FindGiftTypeById(id int) (resource.GiftType, error) {
	var gift Gift
	if err := rds.TestDB(context.TODO()).Select("gift_type").First(&gift, id).Error; err != nil {
		return -1, errors.WithMessage(err, "db first")
	}
	giftType := gift.GiftType
	return giftType, nil
}

func FindByAppId(db *gorm.DB, appId uint) ([]*Gift, error) {
	var gifts []*Gift
	if err := db.Where("app_id = ?", appId).
		Find(&gifts).Error; err != nil {
		return nil, errors.WithMessage(err, "db find")
	}
	return gifts, nil
}

func FindByGroupId(db *gorm.DB, appId, groupId uint) ([]*Gift, error) {
	var gifts []*Gift
	if err := db.Where("app_id = ? and group_id = ?", appId, groupId).
		Find(&gifts).Error; err != nil {
		return nil, errors.WithMessage(err, "db find")
	}
	return gifts, nil
}

func UpdateById(db *gorm.DB, id int, fields map[string]interface{}) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		var gift Gift
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&gift, id).Error; err != nil {
			return errors.WithMessage(err, "db find")
		}
		return tx.Model(&gift).Updates(fields).Error
	})
	if err != nil {
		return errors.WithMessage(err, "db transaction")
	}
	return nil
}
