package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/olivere/elastic"
)

func Router() http.Handler {
	rtr := mux.NewRouter()
	client, _ := GetElasticClient()

	rtr = RouteMap(rtr, client)
	return rtr
}

func RouteMap(mainRouter *mux.Router, client *elastic.Client) *mux.Router {

	mainRouter.Handle("/", Index(client)).Methods("GET")

	return mainRouter
}

