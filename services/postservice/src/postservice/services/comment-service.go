package services

import (
	"postservice/data"
	"gopkg.in/mgo.v2/bson"
)

type CommentService struct {
	comments       []data.Comment
	db             DbService
	collectionName string
}

func (cs *CommentService) CommentService(db DbService, collectionName string) {
	cs.db = db
	cs.collectionName = collectionName
}

func (cs *CommentService) GetAll() []data.Comment {
	session, err, collection := cs.db.GetCollection(cs.collectionName)

	if err != nil {
		panic(err)
	}
	defer session.Close()
	var results []data.Comment
	collection.Find(nil).All(&results)
	return results
}

func (cs *CommentService) Add(comment *data.Comment) {

	session, err, collection := cs.db.GetCollection(cs.collectionName)
	comment.Id = bson.NewObjectId()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	err = collection.Insert(&comment)

}

func (cs *CommentService) Edit(comment *data.Comment) error {
	session, err, collection := cs.db.GetCollection(cs.collectionName)

	colQueried := bson.M{"_id": comment.Id}

	if comment.Title != "" {
		err = collection.Update(colQueried, bson.M{"$set": bson.M{"title": comment.Title}})
	}
	if comment.Body != "" {
		err = collection.Update(colQueried, bson.M{"$set": bson.M{"body": comment.Body}})
	}
	if err != nil {
		panic(err)
	}
	defer session.Close()
	return err
}

func (cs *CommentService) Remove(comment data.Comment) {
	session, err, collection := cs.db.GetCollection(cs.collectionName)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	collection.Remove(comment)
}

func (cs *CommentService) Get(id bson.ObjectId) data.Comment {

	session, err, collection := cs.db.GetCollection(cs.collectionName)

	defer session.Close()
	comment := data.Comment{}
	if collection.FindId(id).One(&comment); err != nil {

		return comment
	}

	return comment
}

func (cs *CommentService) CheckCommentBody(body string) bool {
	_, _, collection := cs.db.GetCollection(cs.collectionName)
	count, err := collection.Find(bson.M{"body": body}).Count()
	if err == nil {
		if count == 0 {
			return true
		}
	}
	return false
}

func (cs *CommentService) CommentExists(id string) bool {
	_, _, collection := cs.db.GetCollection(cs.collectionName)
	count, err := collection.Find(bson.M{"id": id}).Count()
	if err == nil {
		if count == 0 {
			return true
		}
	}

	return false
}
