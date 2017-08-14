package com

import "cloud.google.com/go/pubsub"

const (
	Address = "goconsumer:50051"
)

var (
	Topic *pubsub.Topic
)
