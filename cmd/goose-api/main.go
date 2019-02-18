package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	config := readConfig()
	router := routes(config)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":"+config.port, router))
}
