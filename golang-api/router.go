package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/olivere/elastic"
	"github.com/yogeshsr/search-spike/golang-api/elasticsearch-api"
)

func Router(elasticUrl string) http.Handler {
	rtr := mux.NewRouter()
	client, _ := elasticsearch_api.GetElasticClient(elasticUrl)

	rtr = RouteMap(rtr, client)
	return rtr
}

func RouteMap(mainRouter *mux.Router, client *elastic.Client) *mux.Router {

	mainRouter.Handle("/", elasticsearch_api.Index(client)).Methods("GET")

	return mainRouter
}

