package data

import "gopkg.in/mgo.v2/bson"

type Comment struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	Title  string        `api:"title"`
	Body   string        `api:"body"`

}
