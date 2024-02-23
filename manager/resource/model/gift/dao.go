package gift

import (
	"context"
	"marketing/consts"
	"marketing/database/rds"
	"marketing/util"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DAO interface {
	SetEnv(env consts.Env) DAO
	FindById(ctx context.Context, id int) (*Gift, error)
	FindByAppId(ctx context.Context, appId uint) ([]*Gift, error)
	FindByGroupId(ctx context.Context, appId, groupId uint) ([]*Gift, error)
	UpdateById(ctx context.Context, id int, fields map[string]interface{}) error
}

func InitDAO() DAO {
	if util.IsUnitTest() {
		return &mockDAO{}
	}
	return &rdsDAO{}
}

type rdsDAO struct {
	env consts.Env
}

func (dao *rdsDAO) SetEnv(env consts.Env) DAO {
	dao.env = env
	return dao
}

func (dao *rdsDAO) FindById(ctx context.Context, id int) (*Gift, error) {
	var gift Gift
	if err := rds.DB(ctx, dao.env).First(&gift, id).Error; err != nil {
		return nil, errors.WithMessage(err, "db first")
	}
	return &gift, nil
}

func (dao *rdsDAO) FindByAppId(ctx context.Context, appId uint) ([]*Gift, error) {
	var gifts []*Gift
	if err := rds.DB(ctx, dao.env).Where("app_id = ?", appId).
		Find(&gifts).Error; err != nil {
		return nil, errors.WithMessage(err, "db find")
	}
	return gifts, nil
}

func (dao *rdsDAO) FindByGroupId(ctx context.Context, appId, groupId uint) ([]*Gift, error) {
	var gifts []*Gift
	if err := rds.DB(ctx, dao.env).Where("app_id = ? and group_id = ?", appId, groupId).
		Find(&gifts).Error; err != nil {
		return nil, errors.WithMessage(err, "db find")
	}
	return gifts, nil
}

func (dao *rdsDAO) UpdateById(ctx context.Context, id int, fields map[string]interface{}) error {
	err := rds.DB(ctx, dao.env).Transaction(func(tx *gorm.DB) error {
		var gift Gift
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&gift, id).Error; err != nil {
			return errors.WithMessage(err, "db first")
		}
		return tx.Model(&gift).Updates(fields).Error
	})
	if err != nil {
		return errors.WithMessage(err, "db transaction")
	}
	return nil
}
