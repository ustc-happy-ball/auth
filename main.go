package main

import (
	"flag"
	"github.com/imilano/auth/config"
	"github.com/imilano/auth/handler"
	"log"
)

func cmd() {
	var debug bool
	flag.BoolVar(&debug,"DEBUG", false,"Whether to open debug flag")
	flag.Parse()

	if debug == true {
		log.Println("Starting DEBUG mode...")
		config.DEBUG = true
	}
}

func main() {
	//cmd()
	handler.InitDataBase(config.DBAddr)
	addr := config.IP + ":" + config.PORT
	s,err := handler.NewServer(addr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Starting server, lisen on %v\n",addr)
	s.Serv()
}