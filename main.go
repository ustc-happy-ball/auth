package main

import (
	"github.com/imilano/auth/config"
	"github.com/imilano/auth/db"
	"github.com/imilano/auth/handler"
	pb "github.com/imilano/auth/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis,err := net.Listen("tcp",":"+config.PORT)
	if err != nil {
		log.Fatalf("fail to listen: %v",err)
	}

	d,err := db.New(config.DB,config.ACCOUNT)
	if err != nil {
		log.Fatalf("fail to create database: %v",err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s,&handler.Auth{DB: d})
	if err := s.Serve(lis);err != nil {
		log.Fatalf("failed to serve: %v",err)
	}
}