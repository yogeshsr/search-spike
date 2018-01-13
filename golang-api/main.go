package main

import (
	"log"
)

func main() {
	port := ":8080"
	StartAPI(port)
	log.Printf("starting http://localhost%s\n",port)
}
