package com

import "cloud.google.com/go/pubsub"

const (
	// Address to consumer gRPC
	Address = "goconsumer:50051"
)

var (
	// Topic used for PubSub
	Topic *pubsub.Topic
)
