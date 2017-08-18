package produce

import (
	"encoding/json"
	"fmt"
	"golang-prototype/producer/com"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
)

// PubSubMsg Produces a message for pubsub
var PubSubMsg = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
	returnMsg := sendMsgToConsumerPubSub(newMsg)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(returnMsg); err != nil {
		panic(err)
	}
})

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

	log.Printf("Produced: %s | Sender: %d", msg.Msg, msg.Sender)
	return &com.StatusReport{Status: com.StatusReport_SUCCESS, Message: msg}
}
