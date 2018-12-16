package main

import (
	"github.com/go-chi/render"
	"net/http"
)

type PingRes struct {
	Service string `json:"service"`
	Version string `json:"version"`
	Res     string `json:"res"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	res := PingRes{
		Service: service,
		Version: "0.1.0",
		Res: "pong",
	}

	render.JSON(w, r, res)
}
