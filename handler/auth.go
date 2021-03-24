package handler

import (
	"context"
	pb "github.com/imilano/auth/proto"
)


type auth struct {
	pb.UnimplementedAuthServer
}

func (a *auth) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse,error) {

	return nil,nil
}


func (a *auth) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	return nil,nil
}

func (a *auth)  PingPong(ctx context.Context, req *pb.Ping) (*pb.Pong,error) {
	return nil,nil
}
