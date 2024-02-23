package item

import (
	"context"
	"marketing/consts/errs"
	"marketing/consts/resource"
	"math/rand"
	"reflect"
	"sync"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type mockDAO struct {
	items map[int]*Item
	mu    sync.Mutex
}

func initMockDAO() *mockDAO {
	return &mockDAO{
		items: initTestData(),
		mu:    sync.Mutex{},
	}
}

func initTestData() map[int]*Item {
	return map[int]*Item{
		1: {
			Id:           1,
			AppId:        1,
			ItemType:     resource.Credit,
			ItemId:       1,
			ItemName:     "test-credit-1",
			DeliveryType: resource.Sync,
			CreatedAt:    time.Now().Unix(),
		},
		2: {
			Id:           2,
			AppId:        2,
			ItemType:     resource.Credit,
			ItemId:       2,
			ItemName:     "test-credit-2",
			DeliveryType: resource.Sync,
			CreatedAt:    time.Now().Unix(),
		},
	}
}

func (dao *mockDAO) FindById(ctx context.Context, id int) (*Item, error) {
	item, ok := dao.items[id]
	if !ok {
		return nil, errors.WithMessage(gorm.ErrRecordNotFound, "db first")
	}
	return item, nil
}

func (dao *mockDAO) FindByItemId(ctx context.Context, appId uint, itemId int) (*Item, error) {
	var items []*Item
	for _, item := range dao.items {
		if item.AppId == appId && item.ItemId == itemId {
			items = append(items, item)
		}
	}
	if len(items) == 0 {
		return nil, errs.ItemNotFound
	}
	if len(items) > 1 {
		return nil, errors.WithMessage(errs.Internal, "too many items for one item_id")
	}
	return items[0], nil
}

func (dao *mockDAO) FindByAppId(ctx context.Context, appId uint) ([]*Item, error) {
	var items []*Item
	for _, item := range dao.items {
		if item.AppId == appId {
			items = append(items, item)
		}
	}
	return items, nil
}

func (dao *mockDAO) FirstOrInit(ctx context.Context, rq *AddReq, creator string) (*Item, error) {
	dao.mu.Lock()
	defer dao.mu.Unlock()
	for _, item := range dao.items {
		if item.AppId == rq.AppId && item.ItemId == rq.ItemId {
			return nil, errs.DuplicateItem
		}
	}
	item := new(Item)
	item.AppId = rq.AppId
	item.ItemType = rq.ItemType
	item.ItemId = rq.ItemId
	item.ItemName = rq.ItemName
	item.Descr = rq.Descr
	item.CreatedBy = creator
	item.Id = rand.Intn(10000) + 10000
	dao.items[item.ItemId] = item
	return item, nil
}

func (dao *mockDAO) DeleteById(ctx context.Context, id int) error {
	delete(dao.items, id)
	return nil
}

func (dao *mockDAO) UpdateById(ctx context.Context, id int, fields map[string]interface{}) error {
	item, ok := dao.items[id]
	if !ok {
		return errors.WithMessage(gorm.ErrRecordNotFound, "db first")
	}
	v := reflect.ValueOf(item).Elem()
	for key, value := range fields {
		v.FieldByName(key).Set(reflect.ValueOf(value))
	}
	dao.items[id] = item
	return nil
}
