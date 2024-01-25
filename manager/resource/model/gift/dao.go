package gift

import (
	"context"
	"marketing/database/rds"

	"github.com/pkg/errors"
)

func FindById(ctx context.Context, id int) (*Gift, error) {
	var gift Gift
	if err := rds.GetDB(ctx).First(&gift, id).Error; err != nil {
		return nil, errors.WithMessage(err, "db first")
	}
	return &gift, nil
}

func FindByAppId(ctx context.Context, appId uint) ([]*Gift, error) {
	var gifts []*Gift
	if err := rds.GetDB(ctx).Where("app_id = ?", appId).
		Find(&gifts).Error; err != nil {
		return nil, errors.WithMessage(err, "db find")
	}
	return gifts, nil
}

func FindByGroupId(ctx context.Context, appId, groupId uint) ([]*Gift, error) {
	var gifts []*Gift
	if err := rds.GetDB(ctx).Where("app_id = ? and group_id = ?", appId, groupId).
		Find(&gifts).Error; err != nil {
		return nil, errors.WithMessage(err, "db find")
	}
	return gifts, nil
}
