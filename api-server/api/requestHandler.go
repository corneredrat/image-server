package api
import (
	"net/http"
	"bytes"
	"io"
	"fmt"
	_ "mime/multipart"
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/corneredrat/image-server/api-server/notifier"
)
//@tags Image
//@summary get Image 
//@Param albumname path string true "enter album name whose image is to be returned."
//@Param imagename path string true "enter image name which needs to be returned."
//@success 200 {object} multipart.Form "Status OK."
//@failure 500 {object} map[string]string "on internal server error"
//@Router /album/{albumname}/image/{imagename} [GET]
func GetImage (c *gin.Context) {
	imageName 		:= c.Param("imagename")
	albumName 		:= c.Param("albumname")
	var img 		image

	img, err 	:= getImageFromAlbum(imageName, albumName)
	if err != nil {
		if err.Error() == "ImageNotFound" {
			logAndSetResponse("Image Not Found", http.StatusNotFound, c )
			return
		} else if err.Error() == "AlbumNotFound" {
			logAndSetResponse("Album Not Found", http.StatusNotFound, c)
			return
		} else {
			msg := fmt.Sprintf("unable to fetch records: %v",err.Error())
			logAndSetResponse(msg, http.StatusInternalServerError, c)
			return
		}
	}
	c.FileAttachment(img.location, img.name)
}

//@tags Album
//@Summary get details of an album
//@description returns dictionary object of album
//@Param albumname path string true "Album Name"
//@success 200 {object} map[string]string "returns json data of the album."
//@failure 404 {object} map[string]string "if album is not present."
//@failure 500 {object} map[string]string "on internal server error"
//@Router /album/{albumname} [get]
func GetSingleAlbum (c *gin.Context) {
	
	albumName	:= make([]string, 0)
	name		:= c.Param("albumname")
	albumName	= append(albumName, name)
	res, err 	:= getAlbums(albumName)

	if err != nil {
		if err.Error() == "AlbumNotFound" {
			logAndSetResponse("Album Not Found", http.StatusNotFound, c)
			return
		} else {
			message := fmt.Sprintf("unable to fetch albums from db: %v", err.Error())
			logAndSetResponse(message, http.StatusInternalServerError,  c)
			return	
		}
	}
	c.JSON(http.StatusOK, res)
}

//@tags Album
//@summary delete album
//@Param albumname path string true "enter album name that needs to be deleted."
//@success 200 {object} map[string]string "Status OK."
//@failure 500 {object} map[string]string "on internal server error"
//@Router /album/{albumname} [DELETE]
func DeleteSingleAlbum(c *gin.Context) {
	albumName	:= make([]string, 0)
	name		:= c.Param("albumname")
	albumName	= append(albumName, name)
	res, err 	:= getAlbums(albumName)
	if err != nil {
		if err.Error() == "AlbumNotFound" {
			logAndSetResponse("Album Not Found", http.StatusNotFound, c)
			return
		} else {
			message := fmt.Sprintf("unable to fetch albums from db: %v", err.Error())
			logAndSetResponse(message, http.StatusInternalServerError,  c)
			return	
		}
	}
	// if Album has images...
	if res[albumName[0]] != nil {
		// Remove them  one by one.
		for _, img := range res[albumName[0]].([]map[string]string) {
			imageName	:= img["name"]
			err 		:= deleteImageFromAlbum(imageName, albumName[0])
			if err != nil {
				if err.Error() == "ImageNotFound" {
					logAndSetResponse("Image Not Found", http.StatusNotFound, c )
					return
				} else if err.Error() == "AlbumNotFound" {
					logAndSetResponse("Album Not Found", http.StatusNotFound, c)
					return
				} else {
					message := fmt.Sprintf("unable to delete image from db: %v", err.Error())
					logAndSetResponse(message, http.StatusInternalServerError,  c)
					return
				}
			}	
		}
	}
	
	err = deleteAlbumFromDB(albumName[0])
	if err != nil {
		msg := fmt.Sprintf("unable to delete album from db: %v",err.Error())
		logAndSetResponse(msg, http.StatusInternalServerError, c)
		return 
	}
	c.JSON(http.StatusOK, gin.H{"message":"Album Deleted"})
}
//@tags Image
//@summary get all images in album 
//@Param albumname path string true "enter album name whose images are to be returned."
//@success 200 {object} map[string]string "Status OK."
//@failure 500 {object} map[string]string "on internal server error"
//@Router /album/{albumname}/image [GET]
func GetAllImagesInAlbum(c *gin.Context) {
	albumName	:= make([]string, 0)
	name		:= c.Param("albumname")
	albumName	= append(albumName, name)
	imageURLs 	:= make(map[string]string,0)
	res, err 	:= getAlbums(albumName)
	if err != nil {
		if err.Error() == "AlbumNotFound" {
			logAndSetResponse("Album Not Found", http.StatusNotFound, c)
			return
		} else {
			message := fmt.Sprintf("unable to fetch albums from db: %v", err.Error())
			logAndSetResponse(message, http.StatusInternalServerError,  c)
			return	
		}
	}
	for _, img := range res[albumName[0]].([]map[string]string) {
		imageName				:= img["name"]
		imageURLs[imageName]	= c.Request.URL.String()+"/"+imageName
			
	}
	
	c.JSON(http.StatusOK, imageURLs)
}

//@tags Image
//@summary delete an image 
//@Param albumname path string true "enter albumname in which image is present"
//@Param imagename path string true "enter imagename which needs to be deleted"
//@success 201 {object} map[string]string "Status OK."
//@failure 500 {object} map[string]string "on internal server error"
//@Router /album/{albumname}/image/{imagename} [DELETE]
func DeleteImage(c *gin.Context) {
	albumName	:= c.Param("albumname")
	imageName	:= c.Param("imagename")
	err 		:= deleteImageFromAlbum(imageName, albumName)
	if err != nil {
		if err.Error() == "ImageNotFound" {
			logAndSetResponse("Image Not Found", http.StatusNotFound, c )
			return
		} else if err.Error() == "AlbumNotFound" {
			logAndSetResponse("Album Not Found", http.StatusNotFound, c)
			return
		} else {
			message := fmt.Sprintf("unable to delete image from db: %v", err.Error())
			logAndSetResponse(message, http.StatusInternalServerError,  c)
			return
		}
	}
	notifier.Notify(
		map[string]string{
			"entity"	:"image",
			"album"		: albumName,
			"image"		: imageName,
			"action"	: "create",
		},
		notifier.CreateOp,
	)
	c.JSON(http.StatusAccepted, gin.H{"message":"image deleted from album"})
}

//@tags Album
//@Summary list all the albums 
//@description Returns dictionary object of all albums, use query string parameters to narrow down searching.
//@description example: /album?name=myAlbum1&name=myAlbum2
//@success 200 {object} map[string]string "returns json data of the album."
//@failure 404 {object} map[string]string "if album is not present."
//@failure 500 {object} map[string]string "on internal server error"
//@Router /album [get]
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
			return
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
				return
			}
			c.JSON(http.StatusOK, res)
		}
}

//@tags Album
//@Summary add album
//@description Create album, check example for details on payload.
//@description Example payload:
//@description {
//@description	  "name":"myAlbumName"
//@description }
//@Param request body map[string]string true "refer to example"
//@success 201 {object} map[string]string "Status Created."
//@failure 500 {object} map[string]string "on internal server error"
//@Router /album [POST]
func AddAlbum(c *gin.Context) {

	var inputData	map[string]string
	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		message := fmt.Sprintf("unable to process inputData, check formatting {name: albumname} :%v",err.Error())
		logAndSetResponse(message, http.StatusBadRequest, c)
		return
	}
	if albumName, ok := inputData["name"]; ok {
		err = addAlbumToDB(albumName)
	} else {
		logAndSetResponse("must have key \"name\"", http.StatusBadRequest, c)
		return
	}
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
	c.JSON(http.StatusCreated, gin.H{"message":"album created"})
}

//@tags Image
//@summary add an image 
//@Param image body multipart.Form true "Image that needs to be uploaded"
//@success 201 {object} map[string]string "Status Accepted."
//@failure 500 {object} map[string]string "on internal server error"
//@Router /album/{albumname} [POST]
func AddImage(c *gin.Context) {
	
	// essential variables required to store variables in database/ file system
	var albumName 		string
	var imageName		string
	var imageHash		string
	var imageData		= bytes.NewBuffer(nil)
	var imageLocation 	string
	//var al			album
	//var img 		image		

	// Read uploaded file from request.
	f, uploadedFile, err := c.Request.FormFile("file")
	defer f.Close()
	if nil != err {
		message := "unable read image file. : "+err.Error()
		logAndSetResponse(message, http.StatusInternalServerError, c)
		return
	}
	// compute hash of the file
	if _, err := io.Copy(imageData, f); err != nil {
		message := "unable read image file. : "+err.Error()
		logAndSetResponse(message, http.StatusInternalServerError, c)
		return
	}
	if nil != err {
		message := "unable read image file. : "+err.Error()
		logAndSetResponse(message, http.StatusInternalServerError, c)
		return
	}
	imageName			= uploadedFile.Filename
	imageHash			= computeHash(imageData.Bytes())
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
	notifier.Notify(
		map[string]string{
			"entity"	:"image",
			"album"		: albumName,
			"image"		: imageName,
			"action"	: "create",
		},
		notifier.CreateOp,
	)
	c.JSON(http.StatusCreated, gin.H{"message":"image added"})
}
