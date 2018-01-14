package main

import (
	"github.com/codegangsta/negroni"

	"github.com/sirupsen/logrus"
	"github.com/gorilla/context"
	"github.com/yogeshsr/search-spike/golang-api/pgsearch-api"
)

func startServer(port string, elasticUrl string) {
	router := Router(elasticUrl)
	server := negroni.New(negroni.NewRecovery())
	server.UseHandler(context.ClearHandler(router))
	server.Run(port)
}

func StartAPI(port string, elasticUrl string) {
	logrus.Info("Starting ES Search Service")
	startServer(port, elasticUrl)
}

func StartPGAPI(port string) {
	logrus.Info("Starting PG Search Service")
	startPGServer(port)
}

func startPGServer(port string) {
	router := pgsearch_api.Router()
	server := negroni.New(negroni.NewRecovery())
	server.UseHandler(context.ClearHandler(router))
	server.Run(port)
}