package subscriber

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"

	s "arithmetic-service/service"
)

// Subscribe ...
func Subscribe() {
	natsURL := "nats://localhost:4222"

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Subscriber Connection Error: %v", err)
	}
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalf("Subscriber Encoding Error: %v", err)
	}
	c.Subscribe("arithmetic", func(request *s.ArithmeticDS) {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf(
			"  Got request: %v %c %v = %v\n",
			request.A, request.OP,
			request.B,
			s.GetResult(request),
		)
	})
}
