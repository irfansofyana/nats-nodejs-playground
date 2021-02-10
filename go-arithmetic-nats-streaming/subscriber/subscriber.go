package subscriber

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	s "github.com/irfansofyana/nats-playground/go-arithmetic-nats-streaming/service"

	"github.com/nats-io/stan.go"
)

// Subscribe ...
func Subscribe() {
	clientID := "arithService"
	clusterID := "myCluster"

	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatalf("Subscriber Connection Error: %v", err)
	}
	defer sc.Close()

	sub, err := sc.Subscribe("arithmetic", func(m *stan.Msg) {
		var r *s.ArithmeticDS
		err := json.Unmarshal(m.Data, r)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(10 * time.Millisecond)
		fmt.Printf(
			"  Got request: %v %c %v = %v\n",
			r.A, r.OP,
			r.B,
			s.GetResult(r),
		)
	},
		stan.DurableName("arithService"), // Please remember what I have already received. // HL
		//stan.DeliverAllAvailable(), // HL
	)

	if err != nil {
		log.Fatal(err)
	}
	defer sub.Close()

	select {}
}
