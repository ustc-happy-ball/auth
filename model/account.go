package model

import (
	"context"
	"github.com/imilano/auth/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Account struct {
	ID          int64
	MobilePhone string
	Name        string
	Password    string
	//Avatar      string
	Level       int
	Skin        string
	Deleted     bool
	Region      string
	QQ          string
	Wechat      string
	CreateAt    time.Time
	UpdateAt    time.Time
}

//func init () {
//	var once sync.Once
//	once.Do(func() {
//		createDB(config.DB,config.COLLECTION)
//	})
//}

type DataBase struct {
	DB *mongo.Database
}


// New create db and collection and return db  to caller
func New(dbname string, collection string) (*DataBase, error) {
	ctx,cancel := context.WithTimeout(context.Background(),2 *time.Second)
	defer cancel()


	client,err := mongo.Connect(ctx,options.Client().ApplyURI(config.DBURI))
	if err != nil {
		log.Println("fail to connect to mongo")
	}

	defer func() {
		if err = client.Disconnect(ctx);err != nil {
			panic(err)
		}
	}()

	err = client.Database(dbname).CreateCollection(ctx, collection)
	if err != nil {
		log.Println(err)
		return nil,err
	}
	return &DataBase{client.Database(dbname)},nil
}

// Insert account
func (d *DataBase) InsertOneAccount(ctx context.Context, account *Account) (*mongo.InsertOneResult,error) {
	res,err := d.DB.Collection(config.COLLECTION).InsertOne(ctx, toBson(account))
	if err != nil {
		return nil,err
	}

	return res, nil
}

// Query account
func (d *DataBase) QueryAccountByMobilePhone(ctx context.Context, account *Account) (bool,error) {
	var res *Account
	err := d.DB.Collection(config.COLLECTION).FindOne(ctx,toBson(account)).Decode(res)
	if err != nil {
		return false, err
	}

	return true,nil
}

// Query by id
func (d *DataBase) QueryAccountById(ctx context.Context, id interface{}) (*Account,error) {

	return nil,nil
}

func toBson(account *Account) bson.D {
	return bson.D{
		{"mobilePhone", account.MobilePhone},
		{"password", account.Password},
		{"createAt", account.CreateAt},
	}
}
