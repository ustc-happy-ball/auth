package handler

import (
	"context"
	"github.com/imilano/auth/config"
	"google.golang.org/grpc"
	"log"
	"testing"
	pb "github.com/imilano/auth/proto"
	"time"
)

func TestSignIn(t *testing.T) {
	req := &pb.SignInRequest{
		MobilePhone: "15251859785",
		Password:    "123456",
	}

	conn,err := grpc.Dial("localhost:"+config.PORT,grpc.WithInsecure(),grpc.WithBlock())
	defer  conn.Close()
	if err != nil {
		panic(err)
	}

	c := pb.NewAuthClient(conn)
	ctx,cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	rsp,err := c.SignIn(ctx,req)
	if err != nil {
		log.Println(err)
	}

	log.Printf("%v",rsp.GetPlayerId())
}

func TestSignUp(t *testing.T) {
	req := &pb.SignUpRequest{
		MobilePhone: "15251859785",
		Password:    "123456",
	}

	conn,err := grpc.Dial("localhost:"+config.PORT,grpc.WithInsecure(),grpc.WithBlock())
	defer  conn.Close()
	if err != nil {
		panic(err)
	}

	c := pb.NewAuthClient(conn)
	ctx,cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	rsp,err := c.SignUp(ctx,req)
	if err != nil {
		log.Println(err)
	}

	log.Printf("%v, %v",rsp.GetIsSignUp(),rsp.GetPlayerId())
}

func TestAuth_PingPong(t *testing.T) {
	req := &pb.Ping{Request: "hello"}
	conn,err := grpc.Dial("localhost:"+config.PORT,grpc.WithInsecure(),grpc.WithBlock())
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	client := pb.NewAuthClient(conn)
	ctx,cancel := context.WithTimeout(context.Background(),4*time.Second)
	defer cancel()

	rsp,err := client.PingPong(ctx,req)
	if err != nil {
		log.Printf("error: %v\n",err)
	}
	log.Println(rsp.GetResponse())
}
