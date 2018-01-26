package data

import "gopkg.in/mgo.v2/bson"

type Comment struct {
	Id     bson.ObjectId `bson:"_id" json:"_id,omitempty"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
