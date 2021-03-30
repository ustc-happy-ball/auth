package db

import (
	"context"
	"fmt"
	"github.com/imilano/auth/config"
	"github.com/imilano/auth/model"
	"github.com/imilano/auth/tools"
	"go.mongodb.org/mongo-driver/bson"
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

	coll := client.Database(config.DB).Collection(config.ACCOUNT)
	accounts := []*model.Account{
		&model.Account{
			ID: tools.GenerateUUID64(),
			MobilePhone: "11111",
			LoginPassword: "2222",
			CreateAt: time.Now(),
		},
	}

	for i  := range accounts {
		log.Printf("Insert %+v", accounts[i])

		res,err := coll.InsertOne(ctx,accounts[i])
		log.Println(res)
		if err != nil {
			log.Println(err.Error())
		}

		var account model.Account
		err = coll.FindOne(ctx,bson.D{{"id",accounts[i].ID}}).Decode(&account)

		if err != nil {
			log.Println(err)
		}

		fmt.Printf("Fetch %+v\n",&account)
	}
}

func TestDataBase_InsertOneAccount(t *testing.T) {
	db,err := New(config.DB,config.ACCOUNT)
	if err != nil {
		log.Println(err)
	}

	accounts := []*model.Account{
		&model.Account{
			ID: tools.GenerateUUID64(),
			MobilePhone: "11111",
			LoginPassword: "2222",
			CreateAt: time.Now(),
		},
	}


	ctx,cancel := context.WithTimeout(context.Background(),2*time.Second)
	defer cancel()
	for i  := range accounts {
		log.Printf("Insert %+v", accounts[i])

		res,err := db.InsertOneAccount(ctx,accounts[i])
		log.Println(res)
		if err != nil {
			log.Println(err.Error())
		}

		acc,err := db.QueryAccountByMobilePhone(ctx,accounts[i].MobilePhone)
		if err != nil {
			log.Println(err)
		}

		fmt.Printf("Fetch %+v\n",acc)
	}
}