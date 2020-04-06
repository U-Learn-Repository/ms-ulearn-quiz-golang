package data

import "gopkg.in/mgo.v2/bson"

type DataJSON struct {
	Id   bson.ObjectId `form:"id" bson:"_id,omitempty"`
	Data string        `form:"data" bson:"data"`
}