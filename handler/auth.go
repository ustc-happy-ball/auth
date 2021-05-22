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
	req := reqMsg.Request.SignUpRequest

	// if phone number is wrong
	if !matchPhone(req.MobilePhone) {
		return &pb.GMessage{
			MsgType:  pb.MsgType_RESPONSE,
			MsgCode:  pb.MsgCode_SIGN_UP,
			Response: nil,
			SeqId:    reqMsg.SeqId,
			ErrNum:   pb.ErrNum_WRONG_PHONE_FORMAT,
		},nil
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

	//if !config.DEBUG {
	accountAddRsp, err := (*RemoteDataBase.account).AccountAdd(ctx, accountAddReq)
	if err != nil {
		log.Printf("fail to add account: %v\n", err)
		return &pb.GMessage{
			MsgType:  pb.MsgType_RESPONSE,
			MsgCode:  pb.MsgCode_SIGN_UP,
			SeqId:    reqMsg.SeqId,
			ErrNum:   pb.ErrNum_DUPLICATE_PHONE,
		},nil
	}
	//}

	// Add account info to player collection
	playerAddReq := db.PlayerAddRequest{Player: &db.Player{
		PlayerId:     int32(playerID),
		AccountId:    accountAddRsp.ObjectId,
		HighestScore: 0,
		HighestRank:  0,
		CreateAt:     time.Now().UnixNano(),
		UpdateAt:     time.Now().UnixNano(),
	}}

	doneCh := make(chan int)
	go func() {
		log.Println("Add player info to db: ", playerAddReq.Player)
		ctx := context.Background()
		_,err := (*RemoteDataBase.player).PlayerAdd(ctx,&playerAddReq)
		if err != nil {
			log.Println("fail to add player info to db: ", playerAddReq.Player)
		}

		doneCh <- 1
	}()
	<- doneCh

	return &pb.GMessage{
		MsgType:  pb.MsgType_RESPONSE,
		MsgCode:  pb.MsgCode_SIGN_UP,
		Response: &pb.Response{SignUpResponse: &pb.SignUpResponse{
			IsSignUp: true,
			PlayerId: int32(playerID),
			Addr:     &pb.Address{
				Ip:   config.REMOTE_CLB,
				Port: int32(config.REMOTE_PORT),
			},
		}},
		SeqId:    reqMsg.SeqId,
		ErrNum:   pb.ErrNum_REGULAR_MSG,
	},nil
}

//func (a *Auth) addPlayer(req *db.PlayerAddRequest) error {
//	log.Println("Add player info to db: ", req.Player)
//	ctx := context.Background()
//	_,err := (*RemoteDataBase.player).PlayerAdd(ctx,req)
//	if err != nil {
//		log.Println("fail to add player info to db: ", req.Player)
//		return err
//	}
//
//	return nil
//}

// SignIn for signing in account
func (a *Auth) SignIn(reqMsg *pb.GMessage) (*pb.GMessage,error) {
	req := reqMsg.Request.SignInRequest
	if !matchPhone(req.MobilePhone) {
		return &pb.GMessage{
			MsgType:  pb.MsgType_RESPONSE,
			MsgCode:  pb.MsgCode_SIGN_IN,
			Response: nil,
			SeqId:    reqMsg.SeqId,
			ErrNum:   pb.ErrNum_WRONG_PHONE_FORMAT,
		},nil
	}

	accountFindReq := &db.AccountFindByPhoneRequest{Phone: req.MobilePhone}
	ctx,cancel := context.WithTimeout(context.Background(),2 * time.Second)
	defer cancel()


	if config.DEBUG {
		return &pb.GMessage{
			MsgType:  pb.MsgType_RESPONSE,
			MsgCode:  pb.MsgCode_SIGN_IN,
			Response: nil,
			SeqId:    reqMsg.SeqId,
			ErrNum:   pb.ErrNum_REGULAR_MSG,
		},nil
	}
	// If account doesn't exist or has been deleted
	accountFindRsp, err := (*RemoteDataBase.account).AccountFindByPhone(ctx, accountFindReq)
	if err != nil {
		log.Printf("fail to find account: %v", err)
		return &pb.GMessage{
			MsgType:  pb.MsgType_RESPONSE,
			MsgCode:  pb.MsgCode_SIGN_IN,
			Response: nil,
			SeqId:    reqMsg.SeqId,
			ErrNum:   pb.ErrNum_ACCOUNT_NOT_EXIST,
		},nil
	}

	if accountFindRsp.Account.Delete == true {
		return &pb.GMessage{
			MsgType:  pb.MsgType_RESPONSE,
			MsgCode:  pb.MsgCode_SIGN_IN,
			Response: nil,
			SeqId:    reqMsg.SeqId,
			ErrNum:   pb.ErrNum_ACCOUNT_NOT_EXIST,
		},nil
	}

	// If password mismatch
	if accountFindRsp.Account.LoginPassword != req.Password {
		return &pb.GMessage{
			MsgType:  pb.MsgType_RESPONSE,
			MsgCode:  pb.MsgCode_SIGN_IN,
			Response: nil,
			SeqId:    reqMsg.SeqId,
			ErrNum:   pb.ErrNum_PASSWORD_MISMATCH,
		},nil
	}

	getPlayerInfoReq := db.AccountFindPlayerByAccountIdRequest{AccountId: accountFindRsp.Account.ObjectId}
	ch := make(chan *db.AccountFindPlayerByAccountIdResponse)
	go func() {
		log.Println("Get player info...")
		getPlayerInfoRsp,err := (*RemoteDataBase.account).AccountFindPlayerByAccountId(context.Background(), &getPlayerInfoReq)
		if err != nil {
			log.Println("Fail to get player info according to account object id")
		}
		ch <- getPlayerInfoRsp
	}()
	getPlayerInfoRsp := <- ch

	log.Printf("GetPlayerInfoRsp: %+v\n", getPlayerInfoRsp)
	return &pb.GMessage{
		MsgType:  pb.MsgType_RESPONSE,
		MsgCode:  pb.MsgCode_SIGN_IN,
		Response: &pb.Response{SignInResponse: &pb.SignInResponse{
			IsLogin:  true,
			PlayerId: accountFindRsp.Account.PlayerId,
			HighestRank: getPlayerInfoRsp.PlayerInfo.HighestRank,
			HighestScore: getPlayerInfoRsp.PlayerInfo.HighestScore,
			Addr:     &pb.Address{
				Ip:   config.REMOTE_CLB,
				Port: int32(config.REMOTE_PORT),
			},
		}},
		SeqId:    reqMsg.SeqId,
		ErrNum:   pb.ErrNum_REGULAR_MSG,
	},nil
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
	reg := regexp.MustCompile(`^1(3\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\d|9[0-35-9])\d{8}$`)
	if reg == nil {
		log.Fatalln("regexp err")
	}

	return reg.MatchString(s)
}
