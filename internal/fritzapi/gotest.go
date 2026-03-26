package main

import (
	"fmt"
	"sync"
	"time"
)

var globalstarttime string = time.Now().String()

func worker(id int, wg *sync.WaitGroup, waittime int) { // * tells the program that wg is a pointer to a WaitGroup, allowing the worker function to signal when it's done
	defer wg.Done()
	fmt.Println("Worker", id, time.Now().String(), "start")
	time.Sleep(time.Duration(waittime) * time.Second)
	fmt.Println("Worker", id, time.Now().String(), "end")
}

func main() {
	fmt.Println("Global Start time:", globalstarttime)
	var wg sync.WaitGroup //wg := sync.WaitGroup{}

	for i := 1; i <= 10; i++ { // Initializing i, creating loop condition, incrementing i for each loop. i is a worker and this loop creates 3 workers which start working concurrently
		wg.Add(1)
		go worker(i, &wg, i) // & is a pointer to the WaitGroup, allowing the worker function to signal when it's done
	}

	wg.Wait()
	fmt.Println("All workers finished")
}
