package main

import (
	"fmt"

	pub "github.com/irfansofyana/nats-playground/go-arithmetic-nats-streaming/publisher"
	sub "github.com/irfansofyana/nats-playground/go-arithmetic-nats-streaming/subscriber"
)

func main() {
	fmt.Println("Arithmetic Service Started using NATS streaming")
	sub.Subscribe()
	pub.Publish()
	select {}
}
