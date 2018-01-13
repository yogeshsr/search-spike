package main

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/olivere/elastic"
)

func Index(client *elastic.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		searchParams, ok := r.URL.Query()["search"]
		w.Header().Add("Access-Control-Allow-Origin", "*")
		if !ok || len(searchParams) < 1 {
			log.Println("Url Param 'search' is missing")
			w.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(w).Encode("expected /?search=search query"); err != nil {
				panic(err)
			}
			return
		}
		searchParam := searchParams[0]

		log.Println("Url Param 'search' is: " + string(searchParam))

		vouchers := VoucherSearch(client, searchParam)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "    ")
		if err := encoder.Encode(&vouchers); err != nil {
			panic(err)
		}
	}
}