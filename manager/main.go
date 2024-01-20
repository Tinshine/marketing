package main

import (
	"marketing/item_manager/util/conf"
	"marketing/item_manager/util/log"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	conf.InitConf()
	log.SetLogger()

	h := server.Default()
	register(h)
	h.Spin()
}
