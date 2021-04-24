package handler

import (
	db "github.com/imilano/auth/proto/db"
	"google.golang.org/grpc"
	"log"
)

// remote database client
var RemoteDataBase *DataBase

type DataBase struct {
	account *db.AccountServiceClient
}

func InitDataBase(addr string) {
	log.Println("Starting to initialize db connection......")
	conn,err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	accountClient := db.NewAccountServiceClient(conn)
	RemoteDataBase = &DataBase{account: &accountClient}
}

