package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Account struct {
	Id		   primitive.ObjectID	`bson:"_id"`
	Name       string 				`json:"name"`
	Cpf        string 				`json:"cpf"`
	Balance    float64 				`json:"balance"`
	Created_at time.Time			`json:"created_at"`
}

func (account *Account) Save() {
	account.Id = primitive.NewObjectID()
	account.Created_at = time.Now()
	account.Balance = 30.00
}

func (account *Account) GetAccount(db *mongo.Client, id string) error {
	accountId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": accountId}
	err = GetCollection(db, "accounts").FindOne(context.TODO(), filter).Decode(&account)

	return err
}

func (account *Account) InsertAccount(db *mongo.Client) (*mongo.InsertOneResult, error) {
	insertResult, err := GetCollection(db, "accounts").InsertOne(context.TODO(), account)
	return insertResult, err
}

func (account *Account) UpdateBalance(db *mongo.Client, newBalance float64) (*mongo.UpdateResult, error) {
	result, err := GetCollection(db, "accounts").UpdateOne(context.TODO(),
		bson.M{"_id": account.Id},
		bson.D{
			{"$set", bson.D{{"balance", newBalance}}},
		},)

	if err != nil {
		log.Fatal(err)
	}

	return result, err
}

func GetAccounts(db *mongo.Client) []*Account {
	var accounts []*Account
	cur, err := GetCollection(db, "accounts").Find(context.TODO(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var account Account
		err := cur.Decode(&account)

		if err != nil {
			log.Fatal(err)
		}

		accounts = append(accounts, &account)
	}

	cur.Close(context.TODO())
	return accounts
}

