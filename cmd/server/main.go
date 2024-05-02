package main

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/chrisarmitage/cheat-sheet-layout/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()

	mux.Get("/", handleIndex())
	mux.Get("/template", handleTemplateIndex())

	mux.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))


	slog.Info("starting application", "addr", "localhost:3000")

	err := http.ListenAndServe("localhost:3000", mux)
	if err != nil {
		slog.Error("fatal error", "err", err)
	}
}

func handleIndex() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(views.Index()).ServeHTTP(w, r)
	}
}

func handleTemplateIndex() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(views.TemplateIndex()).ServeHTTP(w, r)
	}
}