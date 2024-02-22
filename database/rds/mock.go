package rds

import (
	"context"
	"marketing/consts"

	"gorm.io/gorm"
)

var (
	mockDB RDS
)

func InitMockDB() {
	mockDB = &mock{}
	mockDB.Init()
}

type mock struct {
	db *gorm.DB
}

func (m *mock) Init() {
	m.db = &gorm.DB{}
}

func (m *mock) ProdDB(ctx context.Context) *gorm.DB {
	return m.db
}

func (m *mock) TestDB(ctx context.Context) *gorm.DB {
	return m.db
}

func (m *mock) DB(ctx context.Context, ev consts.Env) *gorm.DB {
	return m.db
}
