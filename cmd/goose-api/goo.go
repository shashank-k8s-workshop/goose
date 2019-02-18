package main

import (
	"net/http"

	"github.com/go-chi/render"
)

type GooRes struct {
	Service string `json:"service"`
	Version string `json:"version"`
	Res     string `json:"res"`
}

func gooHandler(config *config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := GooRes{
			Service: config.serviceName,
			Version: config.version,
			Res:     "Goose says, Hi!",
		}

		render.JSON(w, r, res)
	}
}
