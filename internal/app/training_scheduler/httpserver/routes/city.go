package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/oleglacto/traning_scheduler/internal/app/training_scheduler/httpserver/handlers"
	"net/http"
)

func initCityRoutes(r chi.Router) {
	r.Route("/city", func(r chi.Router) {
		r.Method(http.MethodGet, "/", Handler(handlers.GetAllCities))
		r.Method(http.MethodPost, "/new", Handler(handlers.AddCity))
	})
}
