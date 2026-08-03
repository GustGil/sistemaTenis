package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	client *mongo.Client
)

func Connect() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	MongoString := os.Getenv("MONGOSTRING")

	client, err = mongo.Connect(
		options.Client().
			ApplyURI(MongoString),
	)

	if err != nil {
		panic(err)
		return
	}

	fmt.Println("Connected to MongoDB!")
	return
}

func Disconnect() error {
	if client != nil {
		return nil
	}
	return client.Disconnect(context.Background())
}

func GetDatabase(db string) *mongo.Database {
	return client.Database(db)
}

func GetCollection(database, collection string) *mongo.Collection {
	db := GetDatabase(database)
	return db.Collection(collection)
}

func InsertAny(database, collection string, data any) {
	c := GetCollection(database, collection)

	result, err := c.InsertOne(context.Background(), data)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.InsertedID)
}

func ProjectionAny(database, collection string, Param *bson.M) []bson.M {
	c := GetCollection(database, collection)
	opts := options.Find().SetProjection(Param)

	cursor, err := c.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		panic(err)
		return nil
	}

	var result []bson.M
	if err := cursor.All(context.Background(), &result); err != nil {
		panic(err)
		return nil
	}
	return result
}
