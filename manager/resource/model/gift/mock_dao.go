package gift

import (
	"context"
	"marketing/consts"
	"reflect"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type mockDAO struct {
	gifts map[uint]*Gift
}

func (dao *mockDAO) SetEnv(env consts.Env) DAO {
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
