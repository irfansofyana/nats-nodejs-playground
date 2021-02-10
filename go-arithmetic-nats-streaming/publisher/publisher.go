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
	clientID := "arithService"
	clusterID := "test-cluster"

	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	go func() {
		for {
			request := &s.ArithmeticDS{
				A:  int64(rand.Int63n(100)),
				B:  int64(rand.Int63n(100)),
				OP: s.RandomOP(rand.Intn(4)),
			}
			ath, err := json.Marshal(request)
			if err != nil {
				log.Fatal(err)
			}
			if err := sc.Publish("arithmetic", ath); err != nil {
				log.Printf("Error couldn't reach NATS server: %v", err)
			}
			fmt.Printf("Published: %v %c %v\n", request.A, request.OP, request.B)
			time.Sleep(3 * time.Second)
		}
	}()
}
