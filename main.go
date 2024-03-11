package main

import (
	"marketing/database/rds"
	"marketing/util/conf"
	"marketing/util/log"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
)

func main() {
	conf.Init()
	log.Init()
	rds.Init()

	h := server.Default()
	h.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                 // Allowed domains, need to bring schema
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"}, // Allowed request methods
		AllowHeaders:     []string{"Origin", "Content-Type"},                // Allowed request headers
		ExposeHeaders:    []string{"Content-Length"},                        // Request headers allowed in the upload_file
		AllowCredentials: true,                                              // Whether cookies are attached
		MaxAge:           12 * time.Hour,                                    // Maximum length of upload_file-side cache preflash requests (seconds)
	}))

	register(h)
	h.Spin()
}
