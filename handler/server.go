package handler

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/imilano/auth/proto"
	"github.com/xtaci/kcp-go"
	"log"
)

type Server struct {
	Auth *Auth
	Listener *kcp.Listener
}

// NewServer return a kcp Server
func NewServer(addr string) (*Server,error) {
	s := new(Server)
	var err error
	s.Listener,err = kcp.ListenWithOptions(addr,nil,0,0)
	if err != nil {
		log.Fatalln(err)
	}

	return s,nil
}

// Serv to listen connection
func (s *Server)Serv () {
	for {
		sess,err := s.Listener.AcceptKCP()
		if err != nil {
			log.Fatalln(err)
		}

		go handler(sess,s)
	}
}

// handler to handle every connection
func handler(conn *kcp.UDPSession,srv *Server) {
	buf := make([]byte,1024)
	for {
		n,err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		go dispatcher(conn,buf[:n],srv)
	}
}

// dispatcher to dispatch message according to MsyType
func dispatcher(conn *kcp.UDPSession, buf []byte, srv *Server) {
	var msg pb.GMessage
	err := proto.Unmarshal(buf,&msg)
	if err != nil {
		log.Println(err)
	}

	var rsp pb.GMessage
	rsp.MsgType = pb.MsgType_RESPONSE
	//defer conn.Close()

	switch msg.MsgType {
	case pb.MsgType_REQUEST:
		switch msg.MsgCode {
		case pb.MsgCode_SIGN_IN:
			log.Println("Receive SignIn Request")
			r,err := srv.Auth.SignIn(msg.Request.GetSignInRequest())
			log.Println("Get signIn response")
			if err != nil {
				log.Println(err)
			}

			rsp.MsgCode = pb.MsgCode_SIGN_IN
			rsp.Response  = &pb.Response{SignInResponse: r}
		case pb.MsgCode_SIGN_UP:
			log.Println("Receive SignUP Request")
			r,err := srv.Auth.SignUp(msg.Request.GetSignUpRequest())
			if err != nil {
				log.Println(err)
			}

			rsp.MsgCode  = pb.MsgCode_SIGN_UP
			rsp.Response = &pb.Response{SignUpResponse: r}
		case pb.MsgCode_REGISTER_ADDR:
			log.Println("Receive RequestAddress request")
			r,err := srv.Auth.Register(msg.Request.GetRegisterRequest())
			if err != nil {
				log.Println(err)
			}

			rsp.MsgCode = pb.MsgCode_REGISTER_ADDR
			rsp.Response = &pb.Response{RegisterResponse: r}
		case pb.MsgCode_PING_PONG:
			log.Println("Receive PingPong request")
			r,err := srv.Auth.PingPong(msg.Request.GetPing())
			if err != nil {
				log.Println(err)
			}

			rsp.MsgCode = pb.MsgCode_PING_PONG
			rsp.Response = &pb.Response{Pong: r}
		default:
			log.Println("MsgCode unknown")
		}
	case pb.MsgType_NOTIFY:
		log.Println("Receive Notify MsgType")
	case pb.MsgType_RESPONSE:
		log.Println("Receive Response MsgType")
	default:
		log.Println("MsgType unknown")
	}

	b,err := proto.Marshal(&rsp)
	if err != nil {
		log.Println(err)
	}

	log.Println("Write back response to client")
	_,err = conn.Write(b)
	if err != nil {
		log.Println(err)
	}
}

