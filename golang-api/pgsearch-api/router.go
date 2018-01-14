package pgsearch_api

import (
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
)

func Router() http.Handler {
	rtr := mux.NewRouter()
	db := GetConncetion()
	db.SetMaxOpenConns(0)
	db.SetMaxIdleConns(10)

	rtr = RouteMap(rtr, db)
	return rtr
}

func RouteMap(mainRouter *mux.Router, db *sql.DB) *mux.Router {

	mainRouter.Handle("/", PGSearchHandler(db)).Methods("GET")

	return mainRouter
}

