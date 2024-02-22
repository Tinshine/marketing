package gift

import (
	"context"
	"fmt"
	"marketing/consts"
	"marketing/consts/resource"
	"reflect"
)

type mockDAO struct {
	val any
}

func (dao *mockDAO) SetEnv(env consts.Env) DAO {
	return dao
}

func (dao *mockDAO) FindById(ctx context.Context, id int) (*Gift, error) {
	return dao.val.(*Gift), nil
}

func (dao *mockDAO) FindGiftTypeById(ctx context.Context, id int) (resource.GiftType, error) {
	return dao.val.(resource.GiftType), nil
}

func (dao *mockDAO) FindByAppId(ctx context.Context, appId uint) ([]*Gift, error) {
	return dao.val.([]*Gift), nil
}

func (dao *mockDAO) FindByGroupId(ctx context.Context, appId, groupId uint) ([]*Gift, error) {
	return dao.val.([]*Gift), nil
}

func (dao *mockDAO) UpdateById(ctx context.Context, id int, fields map[string]interface{}) error {
	g := dao.val.(*Gift)
	v := reflect.ValueOf(g).Elem()

	for k, val := range fields {
		field := v.FieldByName(k)
		if !field.IsValid() {
			return fmt.Errorf("no such field: %s in struct", k)
		}

		if !field.CanSet() {
			return fmt.Errorf("cannot set field: %s in struct", k)
		}

		fieldType := field.Type()
		valType := reflect.ValueOf(val).Type()

		if fieldType != valType {
			return fmt.Errorf("provided value type didn't match obj field type")
		}

		field.Set(reflect.ValueOf(val))
	}
	dao.val = g
	return nil
}
