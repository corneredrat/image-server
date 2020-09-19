package api

import (
	"github.com/gin-gonic/gin"
	"crypto/md5"
	"fmt"
	log "github.com/sirupsen/logrus"
)


func logAndSetResponse(message string, statusCode int,c *gin.Context) {
	log.Error(message)
	c.JSON(statusCode, gin.H{
		"message": message,
	})
}

func computeHash(data []byte) string {
	return fmt.Sprintf("%x",md5.Sum(data))
}