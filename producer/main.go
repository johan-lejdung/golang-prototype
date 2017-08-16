package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/pubsub"

	"golang-prototype/producer/com"
	"golang-prototype/producer/route"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("general.env")
	if err != nil {
		log.Fatal("Error loading general.env file")
	}

	initPubSub()
	initRestRouter()
}

func initRestRouter() {
	router := route.NewRouter()
	log.Printf("Router created")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(headersOk, originsOk, methodsOk)(router)))
}

func initPubSub() {
	log.Printf("Initializing PubSub")

	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("PubSub client PRODUCER created 1/2")
	com.Topic = createTopicIfNotExists(client)
	log.Printf("PubSub topic created 2/2")
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
		// Try 3 more times (since both consumer and producer might try to create at same time), ugly fix
		counter := 0
		for counter < 3 {
			time.Sleep(time.Second * time.Duration(random(1, 4)))
			tt := c.Topic(topic)
			if ok, err = t.Exists(ctx); ok && err == nil {
				return tt
			}
			counter++
		}

		log.Fatalf("Failed to create the topic: %v", err)
	}
	return t
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
