package api

import (
	"context"
	"time"
	"fmt"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	log "github.com/sirupsen/logrus"
	"github.com/corneredrat/image-server/api-server/config"
)

func getClient() (*mongo.Client, context.Context , context.CancelFunc,error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://"+config.Options.Database.URL+":"+config.Options.Database.PORT))
	if err != nil {
		return nil, nil, nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, nil, nil,  err
	}
	return client, ctx, cancel, nil

}

func getAlbums(albumNames []string) (gin.H ,error) {
	albums		:=	make(map[string]interface{})
	var results []bson.M
	client, context, cancel, err := getClient()
	// cleanup
	defer client.Disconnect(context)
	defer cancel()
	if err != nil {
		message := fmt.Sprintf("unable to initialize mongo client. : %v", err.Error())
		return nil, errors.New(message)
	}
	database 			:= client.Database("nokiatask")
	albumsCollection 	:= database.Collection("albums")
	if len(albumNames) == 0 {
		cursor, err := albumsCollection.Find(context, bson.D{}, options.Find())
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return nil, errors.New("AlbumNotFound")
			} else {
				return nil, err
			}
		}
		// load cursor contents into a slice
		if err = cursor.All(context, &results); err != nil {
			return nil, err
		}
		// iterate through all of the results
		for _, result := range results {
			var albm album
			albm.bindBson(result)
			albums[albm.name] = albm.images 
		}
	} else {
		for _,albumName := range albumNames {
			cursor, err := albumsCollection.Find(context, bson.D{bson.E{Key: "name", Value:albumName,}}, options.Find())
			if err != nil {
				if err == mongo.ErrNoDocuments {
					return nil, errors.New("AlbumNotFound")
				} else {
					return nil, err
				}
			}
			// load cursor contents into a slice
			if err = cursor.All(context, &results); err != nil {
				return nil, err
			}
			// iterate through all of the results
			for _, result := range results {
				var albm album
				albm.bindBson(result)
				albums[albm.name] = albm.images 
			}	
		}
	} 
	return albums, nil
}

func addAlbumToDB(albumName string) error {
	
	var result 				bson.M
	var al 					album
	var albumAreadyPresent	bool
	client, context, cancel,  err := getClient()
	// cleanup
	defer client.Disconnect(context)
	defer cancel()
	if err != nil {
		message := fmt.Sprintf("unable to initialize mongo client. %v: ", err.Error())
		return errors.New(message)
	}
	database 			:= client.Database("nokiatask")
	albumsCollection 	:= database.Collection("albums")
	err = albumsCollection.FindOne(context, bson.D{{"name",albumName}}, options.FindOne()).Decode(&result)
	albumAreadyPresent 	= true
	if err != nil {
		if err == mongo.ErrNoDocuments {
			albumAreadyPresent = false
		} else {
			return err
		}
	} 
	if albumAreadyPresent {
		return errors.New("AlreadyExsists")
	} else {
		al.name 	= albumName
		_, err		= albumsCollection.InsertOne(context, al.toBson(), options.InsertOne())
		if err != nil {
			return err
		}
	}
	return nil
}

/**
	WARNING: THIS DOES NOT PERFORM CHECKS JUST DELETES THE ENTRY.
*/
func deleteAlbumFromDB(albumName string) error {
	client, context, cancel, err := getClient()
	// cleanup
	defer client.Disconnect(context)
	defer cancel()
	if err != nil {
		message := fmt.Sprintf("unable to initialize mongo client. : %v", err.Error())
		return errors.New(message)
	}
	database 			:= client.Database("nokiatask")
	albumsCollection 	:= database.Collection("albums")
	_, err = albumsCollection.DeleteOne(context, bson.D{{"name",albumName}},  options.Delete())
	if err != nil {
		return err
	}
	return nil
}



func getImageFromAlbum(imageName string, albumName string) (image, error) {
	
	var result 		bson.M
	var albm 		album
	var img 		image
	var imageHash 	string
	imageFound		:= false
	client, context, cancel, err := getClient()
	// cleanup
	defer client.Disconnect(context)
	defer cancel()
	if err != nil {
		message := fmt.Sprintf("unable to initialize mongo client. : %v", err.Error())
		return img, errors.New(message)
	}
	database 			:= client.Database("nokiatask")
	albumsCollection 	:= database.Collection("albums")
	err = albumsCollection.FindOne(context, bson.D{{"name",albumName}}, options.FindOne()).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return img, errors.New("AlbumNotFound")
		} else {
			return img, err
		}
	}
	
	albm.bindBson(result)
	for _, image := range albm.images {
		if imageName == image["name"] {
			imageFound			= true
			imageHash 			= image["hash"]
			imagesCollection	:= database.Collection("images")
			err = imagesCollection.FindOne(context, bson.D{{"hash",imageHash}}, options.FindOne()).Decode(&result)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					return img, errors.New("ImageNotFound")
				}
				return img, err
			}
			img.bindBson(result)
		}
	}
	if imageFound {
		return img, nil
	} else {
		return img, errors.New("ImageNotFound")
	}
	
}

func addImageToAlbum(albumName, imageName string, imageHash string) error {
	
	var result 	bson.M
	var albm 	album
	client, context, cancel, err := getClient()
	// cleanup
	defer client.Disconnect(context)
	defer cancel()
	if err != nil {
		message := fmt.Sprintf("unable to initialize mongo client. : %v", err.Error())
		return errors.New(message)
	}
	database 			:= client.Database("nokiatask")
	albumsCollection 	:= database.Collection("albums")
	// check if the album exsists in the database
	err = albumsCollection.FindOne(context, bson.D{{"name",albumName}}, options.FindOne()).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("AlbumDoesntExsist")
		} else {
			return err
		}
	}
	// check if image already exsists in the album
	albm.bindBson(result)
	for _, image := range albm.images {
		if imageName == image["name"] {
			return errors.New("ImageExistsInAlbum")
		}
	}
	// add image to album
	imgObject 	:= map[string]string {
		"name":imageName,
		"hash":imageHash,
	}
	albm.images = append(albm.images, imgObject)
	filter 		:= bson.D{{"name", albumName}}
	update 		:= bson.D{{"images",albm.images}}
	opts		:= options.Update()
	_, err 		= albumsCollection.UpdateOne(context, filter,bson.D{{"$set",update}} , opts)
	if err != nil {
    	return err
	}
	return nil
}

func addImage(albumName string, imageName string, imageLocation string, imageHash string) error {
	err := addImageToAlbum(albumName, imageName, imageHash)
	if err != nil {
		if err.Error() == "AlbumDoesntExsist" {
			return err 
		}
		if err.Error() == "ImageExistsInAlbum" {
			return err
		} else {
			msg := fmt.Sprintf("error while adding image to album's entries: %v",err.Error())
			return errors.New(msg)
		}
	}
	err = addImageToDB(imageName, imageLocation, imageHash)
	if err != nil {
		msg := fmt.Sprintf("error while adding image to database: %v",err.Error())
		return errors.New(msg)
	}
	return nil
}

func addImageToDB(imageName string, imageLocation string, imageHash string)  error{
	log.Info("adding image to database..")
	var result	bson.M
	var img		image
	client, context, cancel, err := getClient()
	// cleanup
	defer client.Disconnect(context)
	defer cancel()
	if err != nil {
		message := fmt.Sprintf("unable to initialize mongo client. : %v", err.Error())
		return errors.New(message)
	}
	database 			:= client.Database("nokiatask")
	imagesCollection 	:= database.Collection("images")
	err = imagesCollection.FindOne(context, bson.D{{"hash",imageHash}}, options.FindOne()).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			img.name 		= imageName
			img.location 	= imageLocation
			img.hash 		= imageHash
			img.counter		= 1
			_, err 			:= imagesCollection.InsertOne(context, img.toBson(),options.InsertOne())
			if err != nil {
				return err
			}
			return nil
		} else {
			return err
		}
	}
	// at this point, it appears that image already exsists in the database, so we need to increase the counter - the number of times it is referenced.
	log.Info("image seems to be already present in the db, updating counter.")
	img.bindBson(result)
	img.counter = img.counter + 1
	filter 		:= bson.D{{"hash", imageHash}}
	update 		:= bson.D{{"counter",img.counter}}
	opts		:= options.Update()
	_, err 		= imagesCollection.UpdateOne(context, filter,bson.D{{"$set",update}} , opts)
	if err != nil {
    	return err
	}
	return nil	
}

func deleteImageFromAlbum(imageName string, albumName string) error {
	var result 		bson.M
	var albm 		album
	var img 		image
	var imageHash 	string
	client, context, cancel, err := getClient()
	// cleanup
	defer client.Disconnect(context)
	defer cancel()
	if err != nil {
		message := fmt.Sprintf("unable to initialize mongo client. : %v", err.Error())
		return errors.New(message)
	}
	database 			:= client.Database("nokiatask")
	albumsCollection 	:= database.Collection("albums")
	err = albumsCollection.FindOne(context, bson.D{{"name",albumName}}, options.FindOne()).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("AlbumNotFound")
		} else {
			return err
		}
	}
	albm.bindBson(result)

	if len(albm.images) == 0 {
		return errors.New("ImageNotFound")
	}

	for _, image := range albm.images {
		imageHash 			= image["hash"]
		imagesCollection	:= database.Collection("images")
		err = imagesCollection.FindOne(context, bson.D{{"hash",imageHash}}, options.FindOne()).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return errors.New("ImageNotFound")
			}
			return err
		}
		img.bindBson(result)
		// Update album collection
		images 		:= albm.images
		imgElement 	:= map[string]string {"name":imageName, "hash":imageHash}
		pos 		:= findImageInArray(imgElement, images)
		if pos == -1 {return errors.New("ImageNotFound")}
		log.Info(string(pos))
		images[len(images)-1], images[pos] = images[pos], images[len(images)-1]
		images = images[:len(images)-1]
		filter 		:= bson.D{{"name", albumName}}
		update 		:= bson.D{{"images",images}}
		opts		:= options.Update()
		_, err 		= albumsCollection.UpdateOne(context, filter,bson.D{{"$set",update}} , opts)
		if err != nil {
			return err
		}
		// Update image collection
		if img.counter == 1 {
			log.Warning("This is the last occurance of that image in the filesystem. Deleting...")
			filter 		:= bson.D{{"hash", img.hash}}
			_, err := imagesCollection.DeleteOne(context, filter,  options.Delete())
			if err != nil {
				return err
			}
			deleteFile(img.location)
			if err != nil {
				return err
			} 
		} else {
			img.counter = img.counter - 1
			filter 		:= bson.D{{"hash", img.hash}}
			update 		:= bson.D{{"counter",img.counter}}
			opts		:= options.Update()
			_, err 		= imagesCollection.UpdateOne(context, filter,bson.D{{"$set",update}} , opts)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
