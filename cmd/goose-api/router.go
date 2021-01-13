package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func routes(config *config) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Heartbeat("/heartbeat"),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer)

	router.Route("/", func(r chi.Router) {
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
		})
		router.Get("/goo", gooHandler(config))
	})
	return router
}
