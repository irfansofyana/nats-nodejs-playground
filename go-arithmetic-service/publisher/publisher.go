package publisher

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	s "arithmetic-service/service"

	nats "github.com/nats-io/nats.go"
)

// Publish ...
func Publish() {
	natsURL := "nats://localhost:4222"

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Publisher Connection Error: %v", err)
	}
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalf("Publiser Encoding Error: %v", err)
	}

	go func() {
		for {
			request := &s.ArithmeticDS{
				A:  int64(rand.Int63n(100) + 1),
				B:  int64(rand.Int63n(100) + 1),
				OP: s.RandomOP(rand.Intn(4)),
			}
			if err := c.Publish("arithmetic", request); err != nil {
				log.Printf("Error couldn't reach NATS server: %v", err)
				continue
			}
			fmt.Printf("Published: %v %c %v\n", request.A, request.OP, request.B)
			time.Sleep(3 * time.Second)
		}
	}()
}
