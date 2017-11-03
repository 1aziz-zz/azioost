package services

import (
	"postservice/data"
	"log"
	"gopkg.in/mgo.v2/bson"
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

func (pc *PostService) Add(post *data.Post) error {
	session, err, collection := pc.db.GetCollection(pc.collectionName)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	return collection.Insert(&post)
}

func (pc *PostService) Edit(post *data.Post) error {
	session, err, collection := pc.db.GetCollection(pc.collectionName)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	update := bson.M{"Body": post.Body}

	return collection.Update(post.ID, update)
}

func (pc *PostService) Remove(post data.Post) {
	session, err, collection := pc.db.GetCollection(pc.collectionName)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	collection.Remove(post)
}

func (pc *PostService) Get(id string) data.Post {
	session, err, collection := pc.db.GetCollection(pc.collectionName)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	post := data.Post{}
	err2 := collection.FindId(bson.ObjectIdHex(id)).One(&post)
	if err2 != nil {
		log.Fatal(err2)
	}
	return post
}
