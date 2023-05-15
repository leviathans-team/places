package models

import (
	"github.com/jmoiron/sqlx"
	"gopkg.in/mgo.v2"
)

type Meta struct {
	PlaceId  int64  `bson:"placeId"`
	PlaceUrl string `bson:"placeUrl"`
}

type Filter struct {
	FilterId   int64  `json:"filterId"`
	FilterName string `json:"filterName"`
}

type Place struct {
	PlaceId      int64  `bson:"placeId"`
	PlaceName    string `bson:"placeName"`
	PlaceDestiny string `bson:"placeDestiny"`
	PlaceRoom    int64  `bson:"placeRoom"`
	PlaceCount   int64  `bson:"placeCount"`
}

type ConnectBd struct {
	Database *sqlx.DB
}

var Connection ConnectBd

type MongoConnect struct {
	session *mgo.Session
}

var Mongo MongoConnect
