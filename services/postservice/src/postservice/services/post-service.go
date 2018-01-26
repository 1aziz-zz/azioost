package services

import (
	"postservice/data"
	"gopkg.in/mgo.v2/bson"
)

type PostService struct {
	posts          []data.Post
	db             DbService
	collectionName string
}

func (ps *PostService) PostService(db DbService, collectionName string) {
	ps.db = db
	ps.collectionName = collectionName
}

func (ps *PostService) GetAll() []data.Post {
	session, err, collection := ps.db.GetCollection(ps.collectionName)

	if err != nil {
		panic(err)
	}
	defer session.Close()
	var results []data.Post
	collection.Find(nil).All(&results)
	return results
}

func (ps *PostService) Add(post *data.Post) {

	session, err, collection := ps.db.GetCollection(ps.collectionName)
	post.Id = bson.NewObjectId()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	err = collection.Insert(&post)

}

func (ps *PostService) Edit(post *data.Post) error {
	session, err, collection := ps.db.GetCollection(ps.collectionName)

	colQueried := bson.M{"_id": post.Id}

	if post.Title != "" {
		err = collection.Update(colQueried, bson.M{"$set": bson.M{"title": post.Title}})
	}
	if post.Body != "" {
		err = collection.Update(colQueried, bson.M{"$set": bson.M{"body": post.Body}})
	}
	if err != nil {
		panic(err)
	}
	defer session.Close()
	return err
}

func (ps *PostService) Remove(post data.Post) {
	session, err, collection := ps.db.GetCollection(ps.collectionName)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	collection.Remove(post)
}

func (ps *PostService) Get(id bson.ObjectId) data.Post {

	session, err, collection := ps.db.GetCollection(ps.collectionName)

	defer session.Close()
	post := data.Post{}
	if collection.FindId(id).One(&post); err != nil {

		return post
	}

	return post
}

func (ps *PostService) CheckPostBody(body string) bool {
	_, _, collection := ps.db.GetCollection(ps.collectionName)
	count, err := collection.Find(bson.M{"body": body}).Count()
	if err == nil {
		if count == 0 {
			return true
		}
	}
	return false
}

func (ps *PostService) PostExists(id string) bool {
	_, _, collection := ps.db.GetCollection(ps.collectionName)
	count, err := collection.Find(bson.M{"id": id}).Count()
	if err == nil {
		if count == 0 {
			return true
		}
	}

	return false
}
