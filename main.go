package main

import (
	"github.com/imilano/auth/config"
	"github.com/imilano/auth/handler"
	"log"
)

func main() {
	handler.InitDataBase(config.DBAddr)
	s,err := handler.NewServer(config.IP+":"+config.PORT)
	if err != nil {
		log.Fatalln(err)
	}

	s.Serv()
}