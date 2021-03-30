package db

import (
	"context"
	"github.com/imilano/auth/config"
	"github.com/imilano/auth/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type DataBase struct {
	db *mongo.Database
}

// New create db and collection and return db  to caller
func New(dbname string, collection string) (*DataBase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DBURI))
	if err != nil {
		log.Println("fail to connect to mongodb")
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	coll := client.Database(dbname).Collection(collection)
	return &DataBase{coll.Database()}, nil
}

// Insert account
func (d *DataBase) InsertOneAccount(ctx context.Context, account *model.Account) (*mongo.InsertOneResult, error) {
	res, err := d.db.Collection(config.ACCOUNT).InsertOne(ctx, toBson(account))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query account
func (d *DataBase) QueryAccountByMobilePhone(ctx context.Context, mobilePhone string) (*model.Account, error) {
	var res *model.Account

	filter := bson.D{{"mobilephone", mobilePhone}}
	err := d.db.Collection(config.ACCOUNT).FindOne(ctx, filter).Decode(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query with full account info
func (d *DataBase) QueryAccountByAccountFullInfo(ctx context.Context, account *model.Account) (*model.Account, error) {
	var res *model.Account

	filter := toBson(account)
	err := d.db.Collection(config.ACCOUNT).FindOne(ctx, filter).Decode(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query by id
func (d *DataBase) QueryAccountById(ctx context.Context, id interface{}) (*model.Account, error) {
	var a *model.Account

	filter := bson.D{{"id", id}}
	err := d.db.Collection(config.ACCOUNT).FindOne(ctx, filter).Decode(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func toBson(account *model.Account) bson.D {
	return bson.D{
		{"id", account.ID},
		{"mobilephone", account.MobilePhone},
		{"loginpassword", account.LoginPassword},
		{"createat", account.CreateAt},
		{"name", account.Name},
		{"level", account.Level},
		{"deleted", account.Deleted},
		{"region", account.Region},
		{"accountAvatar", account.AccountAvatar},
		{"skin", account.Skin},
		{"qq", account.QQ},
		{"wechat", account.WeChat},
		{"updateat", account.UpdateAt},
	}
}
