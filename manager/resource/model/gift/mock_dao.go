package gift

import (
	"context"
	"marketing/consts"
	"marketing/consts/resource"
	"reflect"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type mockDAO struct {
	gifts map[uint]*Gift
}

func initTestData() map[uint]*Gift {
	ods := make(map[uint]*Gift)
	ods[1] = &Gift{
		Id:       1,
		AppId:    1,
		GiftType: resource.Normal,
		GiftName: "test-normal",
		GroupId:  1,
		Items: `[
			{
				"item_id": 1,
				"count": 1,
				"role_limit": 1,
				"game_limit": 1
			}
		]`,
		Emails:    "[]",
		CreatedAt: time.Now().Unix(),
	}
	ods[2] = &Gift{
		Id:       2,
		AppId:    1,
		GiftType: resource.Lottery,
		GiftName: "test-lottery",
		GroupId:  2,
		Items: `[
			{
				"item_id": 2,
				"count": 1,
				"role_limit": 1,
				"game_limit": 1
			}
		]`,
		Emails:    "[]",
		CreatedAt: time.Now().Unix(),
	}
	return ods
}

func (dao *mockDAO) SetEnv(env consts.Env) DAO {
	if dao.gifts == nil {
		dao.gifts = initTestData()
	}
	return dao
}

func (dao *mockDAO) FindById(ctx context.Context, id int) (*Gift, error) {
	for i := range dao.gifts {
		if i == uint(id) {
			return dao.gifts[i], nil
		}
	}
	return nil, errors.WithMessage(gorm.ErrRecordNotFound, "db first")
}

func (dao *mockDAO) FindByAppId(ctx context.Context, appId uint) ([]*Gift, error) {
	var gifts []*Gift
	for i := range dao.gifts {
		if i == appId {
			gifts = append(gifts, dao.gifts[i])
		}
	}
	return gifts, nil
}

func (dao *mockDAO) FindByGroupId(ctx context.Context, appId, groupId uint) ([]*Gift, error) {
	var gifts []*Gift
	for i := range dao.gifts {
		if i == appId && i == groupId {
			gifts = append(gifts, dao.gifts[i])
		}
	}
	return gifts, nil
}

func (dao *mockDAO) UpdateById(ctx context.Context, id int, fields map[string]interface{}) error {
	var g *Gift
	for i := range dao.gifts {
		if i == uint(id) {
			g = dao.gifts[i]
		}
	}
	if g == nil {
		return errors.WithMessage(gorm.ErrRecordNotFound, "db first")
	}
	v := reflect.ValueOf(g).Elem()
	for key, value := range fields {
		v.FieldByName(key).Set(reflect.ValueOf(value))
	}
	return nil
}
