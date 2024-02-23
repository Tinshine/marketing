package item

import (
	"context"
	"marketing/consts/errs"
	"marketing/database/rds"
	"marketing/util"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DAO interface {
	FindById(ctx context.Context, id int) (*Item, error)
	FindByItemId(ctx context.Context, appId uint, itemId int) (*Item, error)
	FindByAppId(ctx context.Context, appId uint) ([]*Item, error)
	FirstOrInit(ctx context.Context, rq *AddReq, creator string) (*Item, error)
	DeleteById(ctx context.Context, id int) error
	UpdateById(ctx context.Context, id int, fields map[string]interface{}) error
}

func InitDAO() DAO {
	if util.IsUnitTest() {
		return initMockDAO()
	}
	return &rdsDAO{}
}

type rdsDAO struct{}

func (doa *rdsDAO) FindById(ctx context.Context, id int) (*Item, error) {
	var item Item
	if err := rds.TestDB(ctx).First(&item, id).Error; err != nil {
		return nil, errors.WithMessage(err, "db first")
	}
	return &item, nil
}

func (doa *rdsDAO) FindByItemId(ctx context.Context, appId uint, itemId int) (*Item, error) {
	var items []*Item
	if err := rds.TestDB(ctx).Where("app_id = ? and item_id = ?", appId, itemId).
		Find(&items).Error; err != nil {
		return nil, errors.WithMessage(err, "db find")
	}
	if len(items) == 0 {
		return nil, errs.ItemNotFound
	}
	if len(items) > 1 {
		return nil, errors.WithMessage(errs.Internal, "too many items for one item_id")
	}
	return items[0], nil
}

func (doa *rdsDAO) FindByAppId(ctx context.Context, appId uint) ([]*Item, error) {
	var items []*Item
	if err := rds.TestDB(ctx).Where("app_id = ?", appId).
		Find(&items).Error; err != nil {
		return nil, errors.WithMessage(err, "db find")
	}
	return items, nil
}

func (doa *rdsDAO) FirstOrInit(ctx context.Context, rq *AddReq, creator string) (*Item, error) {
	item := new(Item)
	if err := rds.TestDB(ctx).Transaction(func(tx *gorm.DB) error {
		exist, err := itemExist(tx, rq.AppId, rq.ItemId)
		if err != nil {
			return errors.WithMessage(err, "check item exist")
		}
		if exist {
			return errs.DuplicateItem
		}
		item.AppId = rq.AppId
		item.ItemType = rq.ItemType
		item.ItemId = rq.ItemId
		item.ItemName = rq.ItemName
		item.Descr = rq.Descr
		item.CreatedBy = creator
		return tx.Create(&item).Error
	}); err != nil {
		return nil, errors.WithMessage(err, "db transaction")
	}
	return item, nil
}

func itemExist(tx *gorm.DB, appId uint, itemId int) (bool, error) {
	var item *Item
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(&Item{}).
		Where("app_id = ? and item_id = ?", appId, itemId).First(&item).Error
	if err == nil {
		return true, nil
	}
	if err != gorm.ErrRecordNotFound {
		return false, errors.WithMessage(err, "db find")
	}
	return false, nil
}

func (doa *rdsDAO) DeleteById(ctx context.Context, id int) error {
	if err := rds.TestDB(ctx).Delete(&Item{}, id).Error; err != nil {
		return errors.WithMessage(err, "db delete")
	}
	return nil
}

func (doa *rdsDAO) UpdateById(ctx context.Context, id int, fields map[string]interface{}) error {
	err := rds.TestDB(ctx).Transaction(func(tx *gorm.DB) error {
		var item Item
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&item, id).Error; err != nil {
			return errors.WithMessage(err, "db first")
		}
		return tx.Model(&item).Updates(fields).Error
	})
	if err != nil {
		return errors.WithMessage(err, "db transaction")
	}
	return nil
}
