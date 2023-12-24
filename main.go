package main

import (
	"log"

	"bipbop/server"
)

func main(){
	log.Print("Listening...")
	server.Serve()
}