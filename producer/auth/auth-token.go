package auth

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

// JwtMiddleware to handle authentication
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		decoded, err := base64.URLEncoding.DecodeString(os.Getenv("JWT_SECRET"))
		if err != nil {
			return nil, err
		}
		return decoded, nil
	},
})

// AuthToken has a middleware applied to it, wont ever get here unless authenticated
var AuthToken = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	resp := &Response{
		Message: "All good, you are free to enter!",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
})
