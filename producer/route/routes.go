package route

import (
	"net/http"

	"golang-proto/producer/produce"
)

// Route struct for routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all the routes
type Routes []Route

var routes = Routes{
	Route{
		"ProduceMsg",
		"POST",
		"/produce/msg",
		produce.Msg,
	},
}
