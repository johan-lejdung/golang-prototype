package route

import (
	"net/http"

	"golang-proto/producer/logger"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router using the private variable routes
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, routeIndex := range routes {
		var handler http.Handler

		handler = routeIndex.HandlerFunc
		handler = logger.Logger(handler, routeIndex.Name)
		router.
			Methods(routeIndex.Method).
			Path(routeIndex.Pattern).
			Name(routeIndex.Name).
			Handler(handler)
	}

	return router
}
