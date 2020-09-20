package api

import (
	"github.com/gin-gonic/gin"
	"crypto/md5"
	"fmt"
	"io"
	"errors"
	"strings"
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

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

func saveFile(imageHash string, imageName string, f multipart.File) (string, error) {
	temp 		:= strings.Split(imageName, ".")
	format 		:= temp[len(temp)-1]
	fileName	:= "images/"+imageHash+"."+format
	if fileExists(fileName) {
		return fileName, errors.New("FileExsists")
	}
	localFile, _ 	:= os.Create(fileName)
	_, err 			:= io.Copy(localFile,f)
	if err != nil {
		return fileName, err
	}
	return fileName, nil
}