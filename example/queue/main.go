package main

import (
	"fmt"
	"time"
)

// func main() {
// 	n := 1000
// 	queue := make(chan int, n)

// 	for i := 1; i <= n; i++ {
// 		queue <- i
// 	}

// 	go startConsumer(queue, "C1")
// 	go startConsumer(queue, "C2")

// 	time.Sleep(time.Second * 10)
// }

// func startConsumer(queue chan int, name string) {
// 	for {
// 		time.Sleep(time.Second)
// 		log.Println(name, <-queue)
// 	}
// }

func main() {
	numberOfRequests := 100
	maxWorkNumber := 5
	queueChan := make(chan int, numberOfRequests)
	doneChan := make(chan int)

	for i := 1; i <= numberOfRequests; i++ {
		queueChan <- i
	}

	for i := 1; i <= maxWorkNumber; i++ {
		go func(name string) {
			for v := range queueChan {
				crawl(name, v)
			}

			fmt.Printf("%s is done\n", name)
			doneChan <- i
		}(fmt.Sprintf("%d", i))
	}

	close(queueChan)

	for i := 1; i < maxWorkNumber; i++ {
		<-doneChan
	}
}

func crawl(name string, v int) {
	time.Sleep(time.Second / 3)
	fmt.Printf("Worker %s is crawling: %d \n", name, v)
}
