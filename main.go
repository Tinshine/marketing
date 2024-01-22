package main

import (
	"marketing/database/rds"
	"marketing/util/conf"
	"marketing/util/log"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	conf.Init()
	log.Init()
	rds.Init()

	h := server.Default()
	register(h)
	h.Spin()
}
