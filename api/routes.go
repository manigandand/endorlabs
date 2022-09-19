package api

import (
	"net/http"

	stre "github.com/manigandand/endorlabs/store"

	"github.com/go-chi/chi/v5"
	"github.com/manigandand/adk/api"
	"github.com/manigandand/endorlabs/store/adapter"
)

// store connection copy
var store adapter.ObjectDB

// Routes - all the registered routes
func Routes(router chi.Router) {
	store = stre.Store

	router.Get("/", api.IndexHandeler)
	router.Get("/top", api.HealthHandeler)
	router.Route("/v1", InitV1Routes)
}

// InitV1Routes ...
func InitV1Routes(r chi.Router) {
	r.Route("/objects", func(r chi.Router) {
		r.Method(http.MethodPost, "/persons", api.Handler(savePersonHandler))
		r.Method(http.MethodPost, "/animals", api.Handler(saveAnimalsHandler))

		r.Route("/{objectID}", func(r chi.Router) {
			r.Method(http.MethodGet, "/", api.Handler(getObjectByIDHandler))
			r.Method(http.MethodDelete, "/", api.Handler(deleteObjectHandler))
		})

		r.Method(http.MethodGet, "/name/{objectName}", api.Handler(getObjectByNameHandler))
		r.Method(http.MethodGet, "/kind/{kindName}", api.Handler(listObjectsHandler))
	})
}
