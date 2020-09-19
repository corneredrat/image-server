package api
import (
	"net/http"
	"io"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"crypto/md5"
	log "github.com/sirupsen/logrus"
)

func logAndSetResponse(message string, statusCode int,c *gin.Context) {
	log.Error(message)
	c.JSON(statusCode, gin.H{
		"message": message,
	})
}

func AddImageToAlbum(c *gin.Context) {
	
	// essential variables required to store variables in database/ file system
	//var albumName 	string
	//var imageName	string
	var imageHash	string
	var imageData	[]byte

	//albumName	= c.Param("albumname")
	// Read file from request.
	f, uploadedFile, err := c.Request.FormFile("file")
	if nil != err {
		message := "unable read image file. : "+err.Error()
		logAndSetResponse(message, http.StatusInternalServerError, c)
	}
	// compute hash of the file
	_,err 	= f.Read(imageData)
	if nil != err {
		message := "unable read image file. : "+err.Error()
		logAndSetResponse(message, http.StatusInternalServerError, c)
	}
	imageHash		= fmt.Sprintf("%x",md5.Sum(imageData))
	log.Info("imageHash: ",imageHash)
	localFile, _ 	:= os.Create(uploadedFile.Filename)
	
	_, err = io.Copy(localFile,f)
	if nil != err {
		message := "unable write image file. : "+err.Error()
		logAndSetResponse(message, http.StatusInternalServerError, c)
	}
}
