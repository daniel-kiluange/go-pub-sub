package processor

import (
	"log"
	"os"
	"sync"
	"time"
)

func Process(workers int, bufferSize int) {
	waitGroup := &sync.WaitGroup{}
	channel := make(chan int, bufferSize)

	go publisher(channel)

	for i := 1; i <= workers; i++ {
		waitGroup.Add(1)
		go subscriber(channel, waitGroup, i)
	}
	waitGroup.Wait()
}

func loop(block func()) {
	for {
		block()
	}
}

func publisher(channel chan int) {
	sum := 0
	loop(func() {
		sum++
		channel <- sum
	})
}

func subscriber(channel chan int, workingGroup *sync.WaitGroup, id int) {
	defer workingGroup.Done()
	loop(func() {
		time.Sleep(1000 * time.Millisecond)

		log.Printf("worker id %d value %v t %d", id, <-channel, os.Getgid())
	})
}
