package produce

import (
	"context"
	"encoding/json"
	"golang-proto/producer/com"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"
)

// Msg Produces a message
func Msg(w http.ResponseWriter, r *http.Request) {

	newMsg := &com.ComMsg{
		Msg:    "This is a test message.",
		Sender: 1,
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

	// Check for error
	_, err = proto.Marshal(newMsg)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if errPanic := json.NewEncoder(w).Encode("FAILURE"); errPanic != nil {
			panic(errPanic)
		}
	} else {
		// Send to consumer
		returnMsg := sendMsgToConsumer(newMsg)
		// Send the parentNode
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(returnMsg); err != nil {
			panic(err)
		}
	}
}

func sendMsgToConsumer(msg *com.ComMsg) *com.StatusReport {

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

func sendMsgToConsumerPubSub(msg *com.ComMsg) *com.StatusReport {
	return &com.StatusReport{Status: com.StatusReport_SUCCESS, Message: msg}
}
