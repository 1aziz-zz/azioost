package data

import "gopkg.in/mgo.v2/bson"

type Comment struct {
	//ID     bson.ObjectId `bson:"_id,omitempty"`

	Id     bson.ObjectId `bson:"_id" json:"_id,omitempty"`
	CommentId string        `bson:"commentId" json:"commentId"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
