package main

import (
	"github.com/imilano/auth/config"
	"github.com/imilano/auth/handler"
	pb "github.com/imilano/auth/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis,err := net.Listen("tcp",":" + config.PORT)
	if err != nil {
		log.Fatalf("fail to listen: %v",err)
	}

	//d,err := db.New(config.DB,config.ACCOUNT)
	//ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	//defer cancel()
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DBURI))
	//d := db.DataBase{db:client.Database(config.DB)}
	//d := &db.DataBase{}
	//if err != nil {
	//	log.Fatalf("fail to create database: %v",err)
	//}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s,&handler.Auth{})
	if err := s.Serve(lis);err != nil {
		log.Fatalf("failed to serve: %v",err)
	}
}