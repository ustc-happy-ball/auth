package handler

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/imilano/auth/config"
	pb "github.com/imilano/auth/proto/auth"
	"github.com/xtaci/kcp-go"
	"log"
	"math/rand"
	"regexp"
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

		log.Println(msg.ErrNum)
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
	var times int

	raddr := config.REMOTE_CLB + ":" + strconv.Itoa(config.REMOTE_PORT)
	_ = raddr
	phoneNum := rand.Intn(1000) + 15251859785
	if sess,err :=  kcp.DialWithOptions(raddr,nil,0,0); err == nil {
		go receive(sess)

		var oldPhone string
		for  times != 12 {
			log.Println("Preparing data to send")
			req  := &pb.GMessage{
				MsgType:  pb.MsgType_REQUEST,
			}


			phone := strconv.Itoa(phoneNum)
			switch times%3 {
			case 1:
				req.MsgCode = pb.MsgCode_SIGN_IN
				req.Request = &pb.Request{SignInRequest: &pb.SignInRequest{
					MobilePhone: oldPhone,
					Password:    "2222",
				}}
				log.Printf("Sending SignIn request, seqID %d\n", times)
			case 0:
				req.MsgCode = pb.MsgCode_SIGN_UP
				req.Request = &pb.Request{SignUpRequest: &pb.SignUpRequest{
					MobilePhone: phone,
					Password:    "2222",
				}}
				log.Printf("Sending SignUp request, seqID %d\n", times)
			case 2:
				req.MsgCode = pb.MsgCode_REGISTER_ADDR
				req.Request = &pb.Request{RegisterRequest: &pb.RegisterRequest{}}
				log.Printf("Sending RegisterAddr request, seqID %d\n",times)
			}

			req.SeqId = int32(times)
			data,err := proto.Marshal(req)
			if err != nil {
				log.Println(err)
			}

			if _,err := sess.Write(data); err == nil {
				log.Println("Send data done")
			}

			times++
			oldPhone = phone
			phoneNum++
			time.Sleep(3*time.Second)
		}
	} else {
		log.Fatalln(err)
	}
}

func TestPhone(t *testing.T) {
	reg := regexp.MustCompile("^1(3\\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\\d|9[0-35-9])\\d{8}$")
	if reg == nil {
		log.Fatalln("regexp err")
	}

	strs := []string{
		"111",
		"jfdkfj",
		"133333333333",
		"15251859786",
	}
	for i := range strs {
		if reg.MatchString(strs[i]) {
			log.Println(strs[i], "match")
		}
	}
}
