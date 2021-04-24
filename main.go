package main

import (
	"github.com/imilano/auth/config"
	"github.com/imilano/auth/handler"
	"log"
)

func main() {
	handler.InitDataBase(config.DBAddr)

	addr := config.IP + ":" + config.PORT
	s,err := handler.NewServer(addr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Starting server, lisen on %v\n",addr)
	s.Serv()
}