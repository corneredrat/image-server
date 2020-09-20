package api

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type image struct{
	location	string
	name		string
	hash		string 		`mongo: primary_key`
	counter		int32 		`numer of times the image appears, in different albums`
}

type album struct{
	name		string		`mongo: primary_key`
	images		[]map[string]string
}

func (a album) toBson() (bson.D) {
	if a.images == nil {
		a.images = make([]map[string]string, 0)
	}
	return bson.D {
		{"name"		, a.name},
		{"images"	, a.images},
	}
}

func (a *album) bindBson(bsonObject bson.M) {
	var images []map[string]string
	for _,image := range bsonObject["images"].(bson.A) {
		img 			:= make(map[string]string,0)
		imageFmted		:= image.(primitive.M)
		img["name"] 	= imageFmted["name"].(string)
		img["hash"] 	= imageFmted["hash"].(string)
		images = append(images, img)
	}
	a.name		= bsonObject["name"].(string)
	a.images	= images
}

func (i image) toBson() (bson.D) {
	return bson.D {
		{"name"			, i.name},
		{"hash"			, i.hash},
		{"location"	, i.location,},
		{"counter"		, i.counter},
	}
}

func (i *image) bindBson (bsonObject bson.M) {
	i.name 		= bsonObject["name"].(string)
	i.hash 		= bsonObject["hash"].(string)
	i.location 	= bsonObject["location"].(string)
	i.counter	= bsonObject["counter"].(int32)
}