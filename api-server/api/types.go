package api

import (
	"go.mongodb.org/mongo-driver/bson"
)

type image struct{
	location	string
	name		string
	hash		string 		`mongo: primary_key`
}

type album struct{
	name		string		`mongo: primary_key`
	images		[]string
}

func (a album) toBson() (bson.D) {
	return bson.D {
		{"name"		, a.name},
		{"images"	, a.images},
	}
}

func (i image) toBson() (bson.D) {
	return bson.D {
		{"name"			, i.name},
		{"hash"			, i.hash},
		{"localtion"	, i.location,},
	}
} 