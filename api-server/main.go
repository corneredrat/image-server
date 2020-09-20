package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/corneredrat/image-server/api-server/api"
	"github.com/corneredrat/image-server/api-server/config"
	_ "github.com/corneredrat/image-server/api-server/docs"
)

// @title Image Service API
// @version 0.2
// @description Serves API requests to GET and POST images and albums
// @termsOfService http://swagger.io/terms/

// @contact.name Raghu
// @contact.email raghunandankst@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
// @x-extension-openapi {"example": "value on a json format"}

func main() {
	// initialize config data.
	err := config.Load()
	log.Info("config: ", config.Options)
	if nil != err {
		log.Fatal("unable to initalize configuration. : ", err.Error())
		return
	} else {
		log.Info("loaded config.")
	}

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	r.GET("/album",api.GetAlbum)
	r.GET("/album/:albumname",api.GetSingleAlbum)
	r.DELETE("/album", func (c *gin.Context) {})
	r.POST("/album", api.AddAlbum)
	
	r.GET("/album/:albumname/image",func (c *gin.Context) {
		albumName := c.Param("albumname")
		log.Info("recieved request for album : ", albumName)
	})
	r.DELETE("/album/:album/image",func (c *gin.Context) {} )
	r.POST("/album/:albumname/image",api.AddImage)
	
	r.Run()
}

