// queue：
// length 10，elem type is:  int

// producer：
// produce elem every 1s, blocked when queue is full

// consumer：
// consume elem every 1s, blocked when queue is empty
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// queue
	queue := make(chan int, 10)

	// producer in goroutine
	// produce per 1s
	go func() {
		for {
			randomInt := rand.Intn(100)
			fmt.Printf("Producing [%d] into queue.\n", randomInt)
			queue <- randomInt
			time.Sleep(1 * time.Second)
		}
	}()

	// consumer in main thread
	// consumer per 1s
	for v := range queue {
		fmt.Printf("Consuming [%d] from queue.\n", v)
	}
}
