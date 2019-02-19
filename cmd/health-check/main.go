package main

import (
	"net/http"
	"os"
)

func main() {
	port := "8083"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	url := "http://localhost:" + port + "/goo"
	if _, err := http.Get(url); err != nil {
		os.Exit(1)
	}
}
