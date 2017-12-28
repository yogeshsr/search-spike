package main


import (
	"log"
	"net/http"
)


func main() {
	router := NewRouter()
	port := ":8080"
	log.Printf("starting http://localhost%s\n",port)
	log.Fatal(http.ListenAndServe(port, router))

}
