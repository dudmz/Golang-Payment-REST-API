package model

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

func GetCollection(db *mongo.Client, col string) *mongo.Collection {
	return db.Database(os.Getenv("MONGO_DATABASE")).Collection(col)
}
