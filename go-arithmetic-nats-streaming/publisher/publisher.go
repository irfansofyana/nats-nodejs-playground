package publisher

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	s "github.com/irfansofyana/nats-playground/go-arithmetic-nats-streaming/service"

	stan "github.com/nats-io/stan.go"
)

// Publish ...
func Publish() {
	clientID := "publisher"
	clusterID := "test-cluster"

	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			request := s.ArithmeticDS{
				A:  int64(rand.Int63n(100) + 1),
				B:  int64(rand.Int63n(100) + 1),
				OP: s.RandomOP(rand.Intn(4)),
			}
			ath, err := json.Marshal(request)
			if err != nil {
				log.Fatal(err)
			}
			if err := sc.Publish("arithmetic", ath); err != nil {
				log.Fatalf("FATAL: could not publish: %v", err)
			}
			fmt.Printf("Published: %v %c %v\n", request.A, request.OP, request.B)
			time.Sleep(3 * time.Second)
		}
	}()
}
