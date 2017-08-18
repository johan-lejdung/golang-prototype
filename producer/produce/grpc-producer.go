package produce

import (
	"encoding/json"
	"golang-prototype/producer/com"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// GRPCMsg Produces a message using gRPC
var GRPCMsg = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	newMsg := &com.ComMsg{
		Msg:    "",
		Sender: 0,
	}

	// Get the payload
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, newMsg); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	// Send to consumer
	returnMsg := sendMsgToConsumerGRPC(newMsg)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(returnMsg); err != nil {
		panic(err)
	}
})

func sendMsgToConsumerGRPC(msg *com.ComMsg) *com.StatusReport {

	// Set up a connection to the server.
	conn, err := grpc.Dial(com.Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to gRPC: %v", err)
		panic(err)
	}
	defer conn.Close()
	c := com.NewRouteMsgClient(conn)
	log.Printf("Connected to gRPC Server")

	// Contact the server and print out its response.
	r, err := c.SendMsg(context.Background(), msg)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Produced: %s | Sender: %d", msg.Msg, msg.Sender)
	return r
}
