package main

import (
	"encoding/json"
	"net/http"
)

type Sample struct{
	Foo string  `json:"foo"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	sample := Sample{Foo: "bar"}
	if err := json.NewEncoder(w).Encode(&sample); err != nil {
		panic(err)
	}
}