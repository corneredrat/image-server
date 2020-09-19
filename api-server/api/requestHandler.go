package api
import (
	"net/http"
	"io"
	"os"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)



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
	imageHash		= computeHash(imageData)
	localFile, _ 	:= os.Create(uploadedFile.Filename)
	
	_, err = io.Copy(localFile,f)
	if nil != err {
		message := "unable write image file. : "+err.Error()
		logAndSetResponse(message, http.StatusInternalServerError, c)
	}
}
