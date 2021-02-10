package main

import (
	"fmt"

	pub "arithmetic-service/publisher"
	sub "arithmetic-service/subscriber"
)

func main() {
	fmt.Println("Arithmetic Service Started")
	sub.Subscribe()
	pub.Publish()
	select {}
}
