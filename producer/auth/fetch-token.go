package auth

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type jwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

// FetchToken fetches a token
var FetchToken = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// generate token
	token, err := generateNewJWTAuthToken()
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	resp := &Response{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
})

// GenerateNewJWTAuthToken Generates a JWT from a User object and a secret
func generateNewJWTAuthToken() (string, error) {
	/* Set up a global string for our secret */
	mySigningKey, err := base64.URLEncoding.DecodeString(os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", err
	}

	claims := &jwtCustomClaims{
		"TestName",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}

	/* Create the token */
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)
	/* Finally, write the token to the browser window */
	return tokenString, nil
}
