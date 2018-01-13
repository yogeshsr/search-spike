package main

import (
	"github.com/codegangsta/negroni"

	"github.com/sirupsen/logrus"
	"github.com/gorilla/context"
)

func startServer(port string) {
	router := Router()
	server := negroni.New(negroni.NewRecovery())
	server.UseHandler(context.ClearHandler(router))
	server.Run(port)
}

func StartAPI(port string) {
	logrus.Info("Starting Points Proxy Service")
	startServer(port)
}
