package main

import (
	"os"
)

type config struct {
	port        string
	version     string
	serviceName string
}

func readConfig() *config {
	config := &config{
		port:        "8083",
		version:     "0.1.0",
		serviceName: "goose",
	}
	if os.Getenv("PORT") != "" {
		config.port = os.Getenv("PORT")
	}
	return config
}
