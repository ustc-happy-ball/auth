package handler

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/imilano/auth/proto/auth"
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

	log.Println("Creating server......")
	var err error
	s.Listener,err = kcp.ListenWithOptions(addr,nil,0,0)
	if err != nil {
		log.Println(err)
		return nil,err
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

		log.Println("Accept connection, starting to handle......")
		go handler(sess,s)
	}
}

// handler to handle every connection
func handler(conn *kcp.UDPSession,srv *Server) {
	defer conn.Close()
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

	var rsp *pb.GMessage
	switch msg.MsgType {
	case pb.MsgType_REQUEST:
		switch msg.MsgCode {
		case pb.MsgCode_SIGN_IN:
			log.Printf("Receive SignIn Request, seqId %d\n",msg.SeqId)
			rsp,_= srv.Auth.SignIn(&msg)
		case pb.MsgCode_SIGN_UP:
			log.Printf("Receive SignUP Request,seqId %d\n",msg.SeqId)
			rsp,_ = srv.Auth.SignUp(&msg)
		case pb.MsgCode_REGISTER_ADDR:
			log.Printf("Receive RequestAddress request, seqId %d\n",msg.SeqId)
			rsp,_ = srv.Auth.Register(&msg)

		case pb.MsgCode_PING_PONG:
			log.Println("Receive PingPong request")
			rsp,_ = srv.Auth.PingPong(&msg)
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

	b,err := proto.Marshal(rsp)
	if err != nil {
		log.Println(err)
	}

	log.Printf("Write back response to client, seqID %d\n",msg.SeqId)
	_,err = conn.Write(b)
	if err != nil {
		log.Println(err)
	}
}


