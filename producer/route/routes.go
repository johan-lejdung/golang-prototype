package route

import (
	"net/http"

	"golang-prototype/producer/auth"
	"golang-prototype/producer/produce"
)

// Route struct for routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.Handler
}

// Routes contains all the routes
type Routes []Route

var routes = Routes{
	Route{
		"ProduceGRPCMsg",
		"POST",
		"/produce/grpc/msg",
		produce.GRPCMsg,
	},
	Route{
		"ProducePubSubMsg",
		"POST",
		"/produce/pubsub/msg",
		produce.PubSubMsg,
	},
	Route{
		"JWTFetchToken",
		"GET",
		"/produce/jwt/fetch",
		auth.FetchToken,
	},
	Route{
		"JWTAuthToken",
		"POST",
		"/produce/jwt/auth",
		auth.JwtMiddleware.Handler(auth.AuthToken),
	},
}
