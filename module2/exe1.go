// make it to multiple Producers & Consumers by utilizing goroutine
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
	"sync"
	"time"
)

type Queue struct {
	queue []int      // init slice
	cond  *sync.Cond // pointer to avoid struct copy & we could easily touch
}

func main() {
	// new queue struct
	q := Queue{
		queue: []int{},
		cond:  sync.NewCond(&sync.Mutex{}), // mutex lock as cond, & for pointer to keep consistency.
	}
	// producer goroutine #1
	go func() {
		for {
			randomInt := rand.Intn(100)
			q.Enqueue(randomInt)
			time.Sleep(1 * time.Second)
		}
	}()
	// producer goroutine #2
	go func() {
		for {
			randomInt := rand.Intn(100)
			q.Enqueue(randomInt)
			time.Sleep(1 * time.Second)
		}
	}()
	// consumer goroutine #1
	go func() {
		for {
			q.Dequeue()
			time.Sleep(time.Second)
		}
	}()
	// consumer goroutine #2
	go func() {
		for {
			q.Dequeue()
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(3600 * time.Second) // sleep main, let goroutines work
}

// Enqueue - Queue method, q -> *Queue as receiver
func (q *Queue) Enqueue(elem int) {
	q.cond.L.Lock()                 // lock
	defer q.cond.L.Unlock()         // unlock when method is finished
	q.queue = append(q.queue, elem) // append elem to slice = put elem into
	fmt.Printf("Producing [%d] into queue.\n", elem)
	q.cond.Broadcast() // notify

}

// Dequeue Queue method, q -> *Queue as receiver
func (q *Queue) Dequeue() int {
	q.cond.L.Lock()         // lock
	defer q.cond.L.Unlock() // unlock when method is finished
	for len(q.queue) == 0 {
		fmt.Println("No elem in queue, waiting...")
		q.cond.Wait() // block if no elem in queue
	}
	result := q.queue[0]
	fmt.Printf("Consuming [%d] from queue.\n", result)
	q.queue = q.queue[1:]
	return result
}
