package handler

import (
	"context"
	"github.com/imilano/auth/db"
	pb "github.com/imilano/auth/proto"
	"github.com/imilano/auth/tools"
)


type Auth struct {
	pb.UnimplementedAuthServer
	DB *db.DataBase
}

//func (a *Auth) mustEmbedUnimplementedAuthServer() {
//	panic("implement me")
//}

//func (a *Auth) mustEmbedUnimplementedAuthServer() {
//	panic("implement me")
//}

// TODO just for test
var globalUID = tools.GenerateUUID64()

// SignUp for signing up account, use full account info or just mobilePhone and password
func (a *Auth) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse,error) {
	//uid := tools.GenerateUUID64()
	//account := &model.Account{
	//	ID:            uid,  // TODO 非关系型数据库，如何处理这个ID，这个ID还有必要保留吗?这里的id与MongoDB中的_id有何关系
	//	MobilePhone:   req.MobilePhone,
	//	Name:          "",
	//	LoginPassword: req.Password,
	//	AccountAvatar: "",
	//	Level:         0,
	//	Skin:          "",
	//	Deleted:       false,
	//	Region:        "",
	//	QQ:            "",
	//	WeChat:        "",
	//	CreateAt:      time.Time{},
	//	UpdateAt:      time.Time{},
	//}

	// TODO deal with db connection
	//_,err := a.DB.InsertOneAccount(ctx,account)
	//if err != nil {
	//	log.Println(err)
	//}


	return &pb.SignUpResponse{
		IsSignUp: true,
		PlayerId: globalUID,
},nil
}


func (a *Auth) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	//account := &model.Account{
	//	ID:            0,
	//	MobilePhone:   req.MobilePhone,
	//	Name:          "",
	//	LoginPassword: req.Password,
	//	AccountAvatar: "",
	//	Level:         0,
	//	Skin:          "",
	//	Deleted:       false,
	//	Region:        "",
	//	QQ:            "",
	//	WeChat:        "",
	//	CreateAt:      time.Time{},
	//	UpdateAt:      time.Time{},
	//}

	// TODO deal with db connection
	//account, err := a.DB.QueryAccountByMobilePhone(ctx,req.MobilePhone)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	return &pb.SignInResponse{
		IsLogin:  true,
		PlayerId: globalUID,
	},nil
}

func (a *Auth)  PingPong(ctx context.Context, req *pb.Ping) (*pb.Pong,error) {
	return &pb.Pong{Response: "Hello Back!" },nil
}
