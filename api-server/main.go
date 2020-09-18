package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var cfg config

func main() {
	// initialize config data.
	err := cfg.load()
	if nil != err {
		log.Fatal("unable to initalize configuration. : ", err.Error())
		return
	} else {
		log.Info("loaded config.")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

