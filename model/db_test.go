package model

import (
	"context"
	"github.com/imilano/auth/config"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Test if we could connect to db
func TestMongoConnection(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// trying to connect to mongo
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Println("fail to connect to mongo")
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to MongoDB")
}



// Test if we can insert a document to mongo
func TestMongoInsert(t *testing.T) {
	db,err := New(config.DB,config.COLLECTION)
	if err != nil {
		log.Fatalln(err)
	}

	accounts := []*Account{
		&Account{
			MobilePhone: "11111",
			Password: "2222",
			CreateAt: time.Now(),
		},
	}
	ctx,cancel := context.WithTimeout(context.Background(),2 *time.Second)
	defer cancel()

	for i  := range accounts {
		log.Printf("Insert %+v", accounts[i])

		res, err := db.InsertOneAccount(ctx,accounts[i])
		if err != nil {
			log.Println(err.Error())
		}

		db.QueryAccountByMobilePhone()




	}

}