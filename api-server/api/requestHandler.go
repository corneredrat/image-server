package api
import (
	"net/http"
	"github.com/gin-gonic/gin"
)



func AddImageToAlbum(c *gin.Context) {
	
	// essential variables required to store variables in database/ file system
	var albumName 	string
	var imageName	string
	var imageHash	string
	var imageData	[]byte
	var al			album
	var img 		image		

	// Read uploaded file from request.
	f, uploadedFile, err := c.Request.FormFile("file")
	if nil != err {
		message := "unable read image file. : "+err.Error()
		logAndSetResponse(message, http.StatusInternalServerError, c)
		return
	}
	// compute hash of the file
	_,err 	= f.Read(imageData)
	if nil != err {
		message := "unable read image file. : "+err.Error()
		logAndSetResponse(message, http.StatusInternalServerError, c)
		return
	}
	imageName		= uploadedFile.Filename
	imageHash		= computeHash(imageData)
	err				= saveFile(imageHash, imageName, f)
	if err != nil {
		if nil != err {
			message := "save. : "+err.Error()
			logAndSetResponse(message, http.StatusInternalServerError, c)
			return
		}
	}
}
