package com

import "cloud.google.com/go/pubsub"

const (
	// Port used for gRPC
	Port = ":50051"
)

var (
	// Subscription used for PubSub
	Subscription *pubsub.Subscription
)
