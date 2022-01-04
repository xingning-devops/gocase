package main

import (
	"httpserver/model"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", model.Index)
	http.HandleFunc("/healthz", model.Healthz)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}

}
