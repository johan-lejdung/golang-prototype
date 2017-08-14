package produce

import (
	"encoding/json"
	"fmt"
	"golang-proto/producer/com"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
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

// PubSubMsg Produces a message for pubsub
func PubSubMsg(w http.ResponseWriter, r *http.Request) {

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
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if errPanic := json.NewEncoder(w).Encode("FAILURE"); errPanic != nil {
			panic(errPanic)
		}
	} else {
		// Send to consumer
		returnMsg := sendMsgToConsumerPubSub(newMsg)
		// Send the parentNode
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(returnMsg); err != nil {
			panic(err)
		}
	}
}

func sendMsgToConsumerPubSub(msg *com.ComMsg) *com.StatusReport {
	ctx := context.Background()

	byteArr, _ := proto.Marshal(msg)
	pubsubMsg := &pubsub.Message{
		Data: byteArr,
	}

	if _, err := com.Topic.Publish(ctx, pubsubMsg).Get(ctx); err != nil {
		log.Fatal(fmt.Sprintf("Could not publish message: %v", err), 500)
		return &com.StatusReport{Status: com.StatusReport_ERROR, Message: msg}
	}

	log.Printf("PubSub message published.")
	return &com.StatusReport{Status: com.StatusReport_SUCCESS, Message: msg}
}
