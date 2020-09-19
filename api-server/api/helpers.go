package api

import (
	"github.com/gin-gonic/gin"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"mime/multipart"
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

func saveFile(imageHash string, imageName string, f multipart.File) error {
	localFile, _ 	:= os.Create("images/"+imageHash+"_"+imageName)
	_, err 			:= io.Copy(localFile,f)
	if err != nil {
		return err
	}
	return nil
}