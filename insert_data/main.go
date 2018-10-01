package main

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

const (
	database       = "test_db"
	collection     = "post"
)
type Post struct {
	Tile    string
	Content string
}

func main() {
	client, err := mongo.NewClient("mongodb://user1:example@localhost:27017/test_db")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(database).Collection(collection)
	res, err := collection.InsertOne(context.Background(), Post{"Greeter", "Hello Change stream. You are awesome!"})
	if err != nil { log.Fatal(err) }
	id := res.InsertedID
	log.Println(id)
}

