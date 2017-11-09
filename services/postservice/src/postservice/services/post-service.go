package services

import (
	"postservice/data"
	"gopkg.in/mgo.v2/bson"
	"fmt"

)

type PostService struct {
	posts          []data.Post
	db             DbService
	collectionName string
}

func (pc *PostService) Init(db DbService, collectionName string) {
	pc.db = db
	pc.collectionName = collectionName
}

func (pc *PostService) GetAll() []data.Post {
	session, err, collection := pc.db.GetCollection(pc.collectionName)

	if err != nil {
		panic(err)
	}
	defer session.Close()
	var results []data.Post
	collection.Find(nil).All(&results)
	return results
}

func (pc *PostService) Add(post *data.Post)  {

	session, err, collection := pc.db.GetCollection(pc.collectionName)
	post.Id = bson.NewObjectId()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	err = collection.Insert(&post)

}

func (pc *PostService) Edit(post *data.Post) error {
	session, err, collection := pc.db.GetCollection(pc.collectionName)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	update := bson.M{"Body": post.Body}

	return collection.Update(post.Id, update)
}

func (pc *PostService) Remove(post data.Post) {
	session, err, collection := pc.db.GetCollection(pc.collectionName)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	collection.Remove(post)
}

func (pc *PostService) Get(id bson.ObjectId) data.Post {
	fmt.Println(1)

	session, err, collection := pc.db.GetCollection(pc.collectionName)

	defer session.Close()
	post := data.Post{}
	if collection.FindId(id).One(&post); err != nil {

		return post
	}

	return post
}


