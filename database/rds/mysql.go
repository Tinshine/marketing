package rds

import (
	"context"
	"fmt"
	consts "marketing/consts/conf"
	"marketing/util/conf"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var Init = sync.OnceFunc(setDB)

func setDB() {
	usr, err := conf.GetConf(consts.MySQLConfUserKey)
	if err != nil {
		panic(err)
	}
	pswd, err := conf.GetConf(consts.MySQLConfPswdKey)
	if err != nil {
		panic(err)
	}
	ip, err := conf.GetConf(consts.MySQLConfIPKey)
	if err != nil {
		panic(err)
	}
	port, err := conf.GetConf(consts.MySQLConfPortKey)
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/dbname?charset=utf8mb4&parseTime=True&loc=Local",
		usr, pswd, ip, port,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetDB(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
