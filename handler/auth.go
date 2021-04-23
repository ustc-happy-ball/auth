package handler

import (
	"context"
	"github.com/imilano/auth/config"
	pb "github.com/imilano/auth/proto/auth"
	db "github.com/imilano/auth/proto/db"
	"github.com/imilano/auth/tools"
	"log"
	"time"
)

type Auth struct {}

// SignUp for signing up account, use full account info or just mobilePhone and password
func (a *Auth) SignUp(req *pb.SignUpRequest) (*pb.SignUpResponse,error) {
	playerID := tools.GenerateUUID32()
	accountAddReq := &db.AccountAddRequest{Account: &db.Account{
		ObjectId:      "",
		PlayerId: 		int32(playerID),
		LoginPassword: req.Password,
		Phone:         req.MobilePhone,
		RecentLogin:   time.Now().UnixNano(),
		CreateAt:      time.Now().UnixNano(),
		UpdateAt:      time.Now().UnixNano(),
	}}

	ctx,cancel := context.WithTimeout(context.Background(),2 * time.Second)
	defer cancel()
	_,err := (*RemoteDataBase.account).AccountAdd(ctx,accountAddReq)
	if err != nil {
		log.Printf("fail to add account: %v\n",err)
	}

	return &pb.SignUpResponse{
		IsSignUp: true,
		PlayerId: int32(playerID),
		Addr: &pb.Address{
			Ip:   config.REMOTE_CLB,
			Port: int32(config.REMOTE_PORT),
		},
	},nil
}

func (a *Auth) SignIn(req *pb.SignInRequest) (*pb.SignInResponse, error) {
	accountFindReq := &db.AccountFindByPhoneRequest{Phone: req.MobilePhone}

	ctx,cancel := context.WithTimeout(context.Background(),2 * time.Second)
	defer cancel()
	accountFindRsp,err := (*RemoteDataBase.account).AccountFindByPhone(ctx,accountFindReq)
	if err != nil {
		log.Printf("fail to find account: %v",err)
	}

	// if password does not match
	if accountFindRsp.Account.LoginPassword != req.Password || accountFindRsp.Account.Delete == true{
		return &pb.SignInResponse{
			IsLogin:  false,
			PlayerId: 0,
			Addr:     nil,
		},nil
	}


	return &pb.SignInResponse{
		IsLogin:  true,
		PlayerId: accountFindRsp.Account.PlayerId,
		Addr: &pb.Address{
			Ip:   config.REMOTE_CLB,
			Port: int32(config.REMOTE_PORT),
		},
	},nil
}

func (a *Auth) Register(req *pb.RegisterRequest) (*pb.RegisterResponse,error) {
	return  &pb.RegisterResponse{Addr: &pb.Address{
		Ip:  config.REMOTE_CLB,
		Port: int32(config.REMOTE_PORT),
	}},nil
}

func (a *Auth)  PingPong(req *pb.Ping) (*pb.Pong,error) {
	return &pb.Pong{Rsp: "Hello back!"},nil
}
