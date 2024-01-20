package routes

import (
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
)

var templates = template.Must(template.ParseGlob("internal/templates/**"))

func Router() http.Handler {
	r := chi.NewRouter()

	r.Route("/todos", func(r chi.Router) {
		r.Get("/", getTODOs)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", getTODO)
			r.Delete("/", deleteTODO)
		})
	})

	return r
}
