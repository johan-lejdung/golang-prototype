package logger

import (
	"golang-prototype/consumer/com"
	"log"
	"net/http"
	"time"
)

// Logger provides logs for the accesses to the go server using the routes in routes.go
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func LogMsgConsumed(msg com.ComMsg) {
	log.Printf("Consumed: %s | Sender: %d", msg.Msg, msg.Sender)
}
