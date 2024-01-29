package rds

import (
	"context"
	"fmt"
	"marketing/consts"
	confConst "marketing/consts/conf"
	"marketing/util/conf"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	testDB *gorm.DB
	prodDB *gorm.DB

	Init = sync.OnceFunc(setDB)
)

func setDB() {
	setTestDB()
	setProdDB()
}

func setTestDB() {
	usr, err := conf.GetConf(consts.Test, confConst.MySQLConfUserKey)
	if err != nil {
		panic(err)
	}
	pswd, err := conf.GetConf(consts.Test, confConst.MySQLConfPswdKey)
	if err != nil {
		panic(err)
	}
	ip, err := conf.GetConf(consts.Test, confConst.MySQLConfIPKey)
	if err != nil {
		panic(err)
	}
	port, err := conf.GetConf(consts.Test, confConst.MySQLConfPortKey)
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/dbname?charset=utf8mb4&parseTime=True&loc=Local",
		usr, pswd, ip, port,
	)
	testDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func setProdDB() {
	usr, err := conf.GetConf(consts.Prod, confConst.MySQLConfUserKey)
	if err != nil {
		panic(err)
	}
	pswd, err := conf.GetConf(consts.Prod, confConst.MySQLConfPswdKey)
	if err != nil {
		panic(err)
	}
	ip, err := conf.GetConf(consts.Prod, confConst.MySQLConfIPKey)
	if err != nil {
		panic(err)
	}
	port, err := conf.GetConf(consts.Prod, confConst.MySQLConfPortKey)
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/dbname?charset=utf8mb4&parseTime=True&loc=Local",
		usr, pswd, ip, port,
	)
	prodDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func TestDB(ctx context.Context) *gorm.DB {
	return testDB.WithContext(ctx)
}

func ProdDB(ctx context.Context) *gorm.DB {
	return prodDB.WithContext(ctx)
}

func DB(ctx context.Context, ev consts.Env) *gorm.DB {
	if ev == consts.Test {
		return TestDB(ctx)
	}
	if ev == consts.Prod {
		return ProdDB(ctx)
	}
	return nil
}
