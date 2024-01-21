package main

import (
	"marketing/database/rds"
	"marketing/util/conf"
	"marketing/util/log"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	conf.InitConf()
	log.SetLogger()
	rds.InitMySQL()

	h := server.Default()
	register(h)
	h.Spin()
}
