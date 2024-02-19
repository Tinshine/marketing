package rds

import (
	"context"
	"fmt"
	"marketing/consts"
	confConst "marketing/consts/conf"
	"marketing/util/conf"
	"marketing/util/log"
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
	usr, err := conf.GetConf(consts.Test, confConst.ConfKeyMysqlUser)
	if err != nil {
		panic(err)
	}
	pswd, err := conf.GetConf(consts.Test, confConst.ConfKeyMysqlPswd)
	if err != nil {
		panic(err)
	}
	ip, err := conf.GetConf(consts.Test, confConst.ConfKeyMysqlIP)
	if err != nil {
		panic(err)
	}
	port, err := conf.GetConf(consts.Test, confConst.ConfKeyMysqlPort)
	if err != nil {
		panic(err)
	}
	dbName, err := conf.GetConf(consts.Test, confConst.ConfKeyMysqlDBName)
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		usr, pswd, ip, port, dbName,
	)
	log.Info("setTestDB.conf", "usr", usr, "pswd", pswd, "ip", ip, "port", port, "dsn", dsn)
	testDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func setProdDB() {
	usr, err := conf.GetConf(consts.Prod, confConst.ConfKeyMysqlUser)
	if err != nil {
		panic(err)
	}
	pswd, err := conf.GetConf(consts.Prod, confConst.ConfKeyMysqlPswd)
	if err != nil {
		panic(err)
	}
	ip, err := conf.GetConf(consts.Prod, confConst.ConfKeyMysqlIP)
	if err != nil {
		panic(err)
	}
	port, err := conf.GetConf(consts.Prod, confConst.ConfKeyMysqlPort)
	if err != nil {
		panic(err)
	}
	dbName, err := conf.GetConf(consts.Prod, confConst.ConfKeyMysqlDBName)
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		usr, pswd, ip, port, dbName,
	)
	log.Info("setProdDB.conf", "usr", usr, "pswd", pswd, "ip", ip, "port", port, "dsn", dsn)
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
