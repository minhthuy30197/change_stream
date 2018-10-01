package main

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

const (
	database       = "test_db"
	collection     = "post"
	replicaSetName = "mongo-rs"
)

type Post struct {
	Tile    string `json:"tile" bson:"tile"`
	Content string `json:"content" bson:"content"`
}
type IDELem struct {
	Data string `json:"data" bson:"_data"`
}
type NSELem struct {
	DB   string `json:"db" bson:"db"`
	Coll string `json:"coll" bson:"coll"`
}
type DocumentKeyElem struct {
	ID objectid.ObjectID `json:"id" bson:"_id"`
}
type CSElem struct {
	ID            IDELem          `json:"id" bson:"_id"`
	OperationType string          `json:"operationType" bson:"operationType"`
	FullDocument  Post            `json:"fullDocument" bson:"fullDocument"`
	NS            NSELem          `json:"ns" bson:"ns"`
	DocumentKey   DocumentKeyElem `json:"documentKey" bson:"documentKey"`
}

func main() {
	// Kết nối với MongoDB
	client, err := mongo.NewClient("mongodb://user1:example@localhost:27017/test_db")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(database).Collection(collection)
	ctx := context.Background()

	var pipeline interface{}

	// Theo dõi sự kiện Change stream
	cur, err := collection.Watch(ctx, pipeline)
	if err != nil {
		return
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		elem := CSElem{}
		// Parse json
		if err := cur.Decode(&elem); err != nil {
			log.Fatal(err)
		}

		// In ra màn hình sự kiện change stream
		log.Println(elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
