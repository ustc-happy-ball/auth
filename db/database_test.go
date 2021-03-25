package db

import (
	"context"
	"fmt"
	"github.com/imilano/auth/config"
	"github.com/imilano/auth/model"
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
	db,err := New(config.DB,config.ACCOUNT)

	if err != nil {
		log.Fatalln(err)
	}

	accounts := []*model.Account{
		&model.Account{
			MobilePhone: "11111",
			LoginPassword: "2222",
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

		r,err := db.QueryAccountById(ctx,res.InsertedID)
		if err != nil {
			log.Println(err)
		}

		fmt.Printf("Fetch %+v\n",r)

	}

}