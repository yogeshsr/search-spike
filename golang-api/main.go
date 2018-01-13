package main

import (
	"log"
)

func main() {
	port := ":8080"
	//elasticUrl := "http://localhost:9200"
	elasticUrl := "http://elasticsearch:9200"

	StartAPI(port, elasticUrl)
	log.Printf("starting http://localhost%s\n",port)
}
