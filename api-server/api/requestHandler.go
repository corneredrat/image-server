package api
import (
	"net/http"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func GetSingleAlbum (c *gin.Context) {
	
	albumName	:= make([]string, 0)
	name		:= c.Param("albumname")
	albumName	= append(albumName, name)
	res, err 	:= getAlbums(albumName)

	if err != nil {
		message := fmt.Sprintf("unable to fetch albums from db: %v", err.Error())
		logAndSetResponse(message, http.StatusInternalServerError,  c)
	}
	c.JSON(http.StatusOK, res)
}

func GetAlbum (c *gin.Context) {
	var queryStringParameters 	map[string][]string
	albumNames					:= make([]string,0)
	albumsRequested 			:= make([]string,0)
	queryStringParameters 		= c.Request.URL.Query()
	
	albumsRequested = queryStringParameters["name"]
	if len(albumsRequested) == 0 {
		res , err := getAlbums(albumNames)
		if err != nil {
			message := fmt.Sprintf("unable to fetch albums from db: %v", err.Error())
			logAndSetResponse(message, http.StatusInternalServerError,  c)
		}
		c.JSON(http.StatusOK, res)
	
		} else {
			for _, query := range queryStringParameters["name"] {
				log.Info("processing [get] query for album : ",query)
				albumNames = append(albumNames, query)
			}
			res, err := getAlbums(albumNames)
			if err != nil {
				message := fmt.Sprintf("unable to fetch albums from db: %v", err.Error())
				logAndSetResponse(message, http.StatusInternalServerError,  c)
			}
			c.JSON(http.StatusOK, res)
		}
}

func AddAlbum(c *gin.Context) {

	var inputData	map[string]string
	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		message := fmt.Sprintf("unable to process inputData, check formatting {name: albumname} :%v",err.Error())
		logAndSetResponse(message, http.StatusBadRequest, c)
		return
	}
	err = addAlbumToDB(inputData["name"])
	if err != nil {
		if err.Error() == "AlreadyExsists" {
			message := fmt.Sprintf("Album with the name %v already exsists.",inputData["name"])
			logAndSetResponse(message, http.StatusNotAcceptable, c)
			return
		} else {
			message := fmt.Sprintf("unable to process request: ",err.Error())
			logAndSetResponse(message, http.StatusInternalServerError, c)
			return
		}
	}
}

func AddImage(c *gin.Context) {
	
	// essential variables required to store variables in database/ file system
	var albumName 		string
	var imageName		string
	var imageHash		string
	var imageData		[]byte
	var imageLocation 	string
	//var al			album
	//var img 		image		

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
	imageName			= uploadedFile.Filename
	imageHash			= computeHash(imageData)
	imageLocation, err	= saveFile(imageHash, imageName, f)
	if err != nil {
		if nil != err {
			if err.Error() != "FileExsists" {
				message := "error while saving file : "+err.Error()
				logAndSetResponse(message, http.StatusInternalServerError, c)
				return
			}
			
		}
	}
	albumName = c.Param("albumname")
	err = addImage(
		albumName,
		imageName,
		imageLocation,
		imageHash)
	if err != nil {
		if err.Error() == "AlbumDoesntExsist" {
			msg := fmt.Sprintf("%v album does not exsist.",albumName)
			logAndSetResponse(msg, http.StatusNotAcceptable, c)
			return 
		} else if err.Error() == "ImageExistsInAlbum" {
			msg := fmt.Sprintf("%v image already exsists in %v album.",imageName, albumName)
			logAndSetResponse(msg, http.StatusNotAcceptable, c)
			return
		} else {
			msg := fmt.Sprintf("unable to add %v image to %v album : %v",imageName, albumName, err.Error())
			logAndSetResponse(msg, http.StatusInternalServerError, c)
			return
		}
	}		
	return
}
