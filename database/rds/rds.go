package rds

import (
	"context"
	"marketing/consts"

	"gorm.io/gorm"
)

// RDS represents for relationship database
type RDS interface {
	Init()
	ProdDB(ctx context.Context) *gorm.DB
	TestDB(ctx context.Context) *gorm.DB
	DB(ctx context.Context, ev consts.Env) *gorm.DB
}
