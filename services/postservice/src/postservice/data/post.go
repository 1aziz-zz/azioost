package data

import "gopkg.in/mgo.v2/bson"

type Post struct {
	Id     bson.ObjectId `bson:"_id" json:"_id,omitempty"`
	Title  string        `bson:"title" json:"title"`
	Body   string        `bson:"body" json:"body"`
}
