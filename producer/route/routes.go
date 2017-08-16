package route

import (
	"net/http"

	"golang-prototype/producer/produce"
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
		"ProduceGRPCMsg",
		"POST",
		"/produce/msg",
		produce.GRPCMsg,
	},
	Route{
		"ProducePubSubMsg",
		"POST",
		"/produce/pubsub/msg",
		produce.PubSubMsg,
	},
}
