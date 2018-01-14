package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
)

var (
	serverMode = kingpin.Command("startES", "Starting Elastic based search api").Default()
	indexCreationMode = kingpin.Command("startPG", "Starting Postgres free text based search api")
)


func main() {

	port := ":8080"

	// cmd line mode will have precedence over env defined

	cmdLineMode := kingpin.Parse()
	envMode := os.Getenv("mode")
	mode := envMode
	if envMode == "" {
		mode = cmdLineMode
	}

	log.Printf("starting mode: %s", mode)

	switch mode {
		case "startPG":
			StartPGAPI(port)
		default:
			elasticUrl := "http://localhost:9200"
			//elasticUrl := "http://elasticsearch:9200"
			StartAPI(port, elasticUrl)
	}
	log.Printf("starting http://localhost%s\n", port)

}
