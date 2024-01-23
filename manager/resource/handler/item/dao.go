package item

import (
	"context"
	"marketing/database/rds"

	"github.com/pkg/errors"
)

func FindById(ctx context.Context, id int) (*Item, error) {
	var item Item
	if err := rds.GetDB(ctx).First(&item, id).Error; err != nil {
		return nil, errors.WithMessage(err, "db first")
	}
	return &item, nil
}

func FindByItemId(ctx context.Context, appId, itemId int) ([]*Item, error) {
	var items []*Item
	if err := rds.GetDB(ctx).Where("app_id = ? and item_id = ?", appId, itemId).
		Find(&items).Error; err != nil {
		return nil, errors.WithMessage(err, "db find")
	}
	return items, nil
}

func FindByAppId(ctx context.Context, appId int) ([]*Item, error) {
	var items []*Item
	if err := rds.GetDB(ctx).Where("app_id = ?", appId).
		Find(&items).Error; err != nil {
		return nil, errors.WithMessage(err, "db find")
	}
	return items, nil
}
