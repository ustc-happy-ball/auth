package handler

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/imilano/auth/config"
	pb "github.com/imilano/auth/proto/auth"
	"github.com/xtaci/kcp-go"
	"log"
	rand2 "math/rand"
	"strconv"
	"testing"
	"time"
)

func receive(sess *kcp.UDPSession) {
	for {
		buf := make([]byte,1024)
		n,err := sess.Read(buf)
		if err != nil {
			log.Println(err)
		}

		msg := &pb.GMessage{}
		err = proto.Unmarshal(buf[:n],msg)
		if err != nil {
			log.Println(err)
		}

		switch msg.MsgCode {
		case pb.MsgCode_PING_PONG:
		case pb.MsgCode_SIGN_IN:
			log.Println("Receive signIn response")
			log.Printf("%+v",msg.Response.SignInResponse)
		case pb.MsgCode_SIGN_UP:
			log.Println("Receive signUp response")
			log.Printf("%+v",msg.Response.SignUpResponse)
		case pb.MsgCode_REGISTER_ADDR:
			log.Println("Receive register response")
			log.Printf("%+v",msg.Response.RegisterResponse)
		default:
			log.Println("Unknown response type")
		}
	}
}


func TestAuth(t *testing.T) {
	fmt.Println("Starting to test auth service")
	var rand int

	if sess,err :=  kcp.DialWithOptions(config.IP+":"+"8889",nil,0,0); err == nil {
		go receive(sess)
		for  rand != 3{
			log.Println("Preparing data to send")
			req  := &pb.GMessage{
				MsgType:  pb.MsgType_REQUEST,
			}

			s := rand2.Int()
			phone := strconv.Itoa(s)
			switch rand%3 {
			case 1:
				req.MsgCode = pb.MsgCode_SIGN_IN
				req.Request = &pb.Request{SignInRequest: &pb.SignInRequest{
					MobilePhone: phone,
					Password:    "2222",
				}}
				log.Printf("Sending SignIn request, seqID %d\n",rand)
			case 0:
				req.MsgCode = pb.MsgCode_SIGN_UP
				req.Request = &pb.Request{SignUpRequest: &pb.SignUpRequest{
					MobilePhone: phone,
					Password:    "2222",
				}}
				log.Printf("Sending SignUp request, seqID %d\n",rand)
			case 2:
				req.MsgCode = pb.MsgCode_REGISTER_ADDR
				req.Request = &pb.Request{RegisterRequest: &pb.RegisterRequest{}}
				log.Printf("Sending RegisterAddr request, seqID %d\n",rand)
			}

			req.SeqId = int32(rand)
			data,err := proto.Marshal(req)
			if err != nil {
				log.Println(err)
			}

			if _,err := sess.Write(data); err == nil {
				log.Println("Send data done")
			}

			rand++
			time.Sleep(3*time.Second)
		}
	} else {
		log.Fatalln(err)
	}
}
