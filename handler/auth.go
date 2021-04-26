package handler

import (
	"context"
	"github.com/imilano/auth/config"
	pb "github.com/imilano/auth/proto/auth"
	db "github.com/imilano/auth/proto/db"
	"github.com/imilano/auth/tools"
	"log"
	"regexp"
	"time"
)

type Auth struct {}

// SignUp for signing up account, use full account info or just mobilePhone and password
func (a *Auth) SignUp(reqMsg *pb.GMessage) (*pb.GMessage,error) {
	var msg pb.GMessage
	msg.SeqId = reqMsg.SeqId
	msg.MsgCode = pb.MsgCode_SIGN_UP
	msg.MsgType = pb.MsgType_RESPONSE
	req := reqMsg.Request.SignUpRequest

	// if phone number is wrong
	if !matchPhone(req.MobilePhone) {
		msg.ErrNum = pb.ErrNum_WRONG_PHONE_FORMAT
		return &msg,nil
	}

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
	// TODO 如果 account 的信息重复了

	msg.ErrNum = pb.ErrNum_REGULAR_MSG
	msg.Response.SignUpResponse = &pb.SignUpResponse{
		IsSignUp: true,
		PlayerId: int32(playerID),
		Addr:     &pb.Address{
			Ip:   config.REMOTE_CLB,
			Port: int32(config.REMOTE_PORT),
		},
	}
	return &msg,nil
}

// SignIn for signing in account
func (a *Auth) SignIn(reqMsg *pb.GMessage) (*pb.GMessage,error) {
	var msg pb.GMessage
	msg.MsgCode = pb.MsgCode_SIGN_IN
	msg.SeqId = reqMsg.SeqId
	msg.MsgType = pb.MsgType_RESPONSE

	req := reqMsg.Request.SignInRequest
	if !matchPhone(req.MobilePhone) {
		msg.ErrNum = pb.ErrNum_WRONG_PHONE_FORMAT
		return &msg,nil
	}

	accountFindReq := &db.AccountFindByPhoneRequest{Phone: req.MobilePhone}
	ctx,cancel := context.WithTimeout(context.Background(),2 * time.Second)
	defer cancel()

	// If account doesn't exist or has been deleted
	accountFindRsp,err := (*RemoteDataBase.account).AccountFindByPhone(ctx,accountFindReq)
	if err != nil {
		log.Printf("fail to find account: %v",err)
		msg.ErrNum = pb.ErrNum_ACCOUNT_NOT_EXIST
		return &msg,nil
	}

	if accountFindRsp.Account.Delete == true {
		msg.ErrNum = pb.ErrNum_ACCOUNT_NOT_EXIST
		return &msg,nil
	}

	// If password mismatch
	if accountFindRsp.Account.LoginPassword != req.Password {
		msg.ErrNum = pb.ErrNum_PASSWORD_MISMATCH
		return &msg,nil
	}


	msg.ErrNum = pb.ErrNum_REGULAR_MSG
	msg.Response.SignInResponse = &pb.SignInResponse{
		IsLogin:  true,
		PlayerId: accountFindRsp.Account.PlayerId,
		Addr:     &pb.Address{
			Ip:  config.REMOTE_CLB,
			Port: int32(config.REMOTE_PORT),
		},
	}

	return &msg,nil
}

// Register to notify dgs address to client
func (a *Auth) Register(reqMsg *pb.GMessage) (*pb.GMessage,error) {
	return &pb.GMessage{
		MsgType:  pb.MsgType_RESPONSE,
		MsgCode:  pb.MsgCode_REGISTER_ADDR,
		Response: &pb.Response{RegisterResponse: &pb.RegisterResponse{Addr: &pb.Address{
			Ip: config.REMOTE_CLB,
			Port: int32(config.REMOTE_PORT),
		}}},
		SeqId:   reqMsg.SeqId,
		ErrNum:   pb.ErrNum_REGULAR_MSG,
	},nil
}

// PingPong for test
func (a *Auth)  PingPong(req *pb.GMessage) (*pb.GMessage,error) {
	return &pb.GMessage{
		MsgType:  pb.MsgType_REQUEST,
		MsgCode:  pb.MsgCode_PING_PONG,
		Response: &pb.Response{Pong: &pb.Pong{Rsp: "Hello back!"}},
		SeqId:    req.SeqId,
		ErrNum:   pb.ErrNum_REGULAR_MSG,
	},nil
}

func matchPhone(s string) bool {
	reg := regexp.MustCompile("^1(3\\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\\d|9[0-35-9])\\d{8}$")
	if reg == nil {
		log.Fatalln("regexp err")
	}

	return reg.MatchString(s)
}
