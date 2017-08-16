package main

import (
	"golang-prototype/consumer/com"
	"golang-prototype/consumer/logger"
	"log"
	"net"
	"os"

	"github.com/golang/protobuf/proto"

	"cloud.google.com/go/pubsub"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/joho/godotenv"
)

// server is used to implement com.RouteMsgServer.
type server struct{}

// ConsumeMsg implements com.RouteMsgServer
func (s *server) SendMsg(ctx context.Context, in *com.ComMsg) (*com.StatusReport, error) {
	logger.LogMsgConsumed(*in)
	return &com.StatusReport{Status: com.StatusReport_SUCCESS, Message: in}, nil
}

func main() {
	err := godotenv.Load("general.env")
	if err != nil {
		log.Fatal("Error loading general.env file")
	}

	// Need to call with goroutine because each will lock the thread waiting for messages
	go initPubSub()
	initGRPC()
}

func initGRPC() {
	log.Printf("Starting gRPC Server")
	lis, err := net.Listen("tcp", com.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	com.RegisterRouteMsgServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("gRPC server created")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initPubSub() {
	log.Printf("Initializing PubSub")

	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("PubSub client CONSUMER created")
	// Use a callback to receive messages via sub.
	com.Subscription = createSubscriptionIfNotExists(client)
	err = com.Subscription.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		newMsg := &com.ComMsg{}
		proto.Unmarshal(m.Data, newMsg)
		logger.LogMsgConsumed(*newMsg)
		m.Ack() // Acknowledge that we've consumed the message.
	})
	if err != nil {
		log.Fatal(err)
	}
}

func createSubscriptionIfNotExists(c *pubsub.Client) *pubsub.Subscription {
	ctx := context.Background()

	sub := os.Getenv("PUBSUB_SUBSCRIPTION")
	// Create a topic to subscribe to.
	s := c.Subscription(sub)
	ok, err := s.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		return s
	}

	s, err = c.CreateSubscription(ctx, sub, pubsub.SubscriptionConfig{Topic: createTopicIfNotExists(c)})
	if err != nil {
		log.Fatalf("Failed to create the subscription: %v", err)
	}
	return s
}

func createTopicIfNotExists(c *pubsub.Client) *pubsub.Topic {
	ctx := context.Background()

	topic := os.Getenv("PUBSUB_TOPIC")
	// Create a topic to subscribe to.
	t := c.Topic(topic)
	ok, err := t.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		return t
	}

	t, err = c.CreateTopic(ctx, topic)
	if err != nil {
		log.Fatalf("Failed to create the topic: %v", err)
	}
	return t
}
