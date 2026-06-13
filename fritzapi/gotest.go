package main

import (
	"fmt"
	"sync"
	"time"
)

var globalstarttime string = time.Now().String()

func worker(id int, wg *sync.WaitGroup, waittime int, wtype string) { // * tells the program that wg is a pointer to a WaitGroup, allowing the worker function to signal when it's done
	defer wg.Done()
	fmt.Println(time.Now(), "Worker", wtype, id, "start")
	time.Sleep(time.Duration(waittime) * time.Second)
	fmt.Println(time.Now(), "Worker", wtype, id, "end")
}

func main() {
	fmt.Println("Global Start time:", globalstarttime)
	var wg sync.WaitGroup //wg := sync.WaitGroup{}

	for i := 1; i <= 3; i++ { // Initializing i, creating loop condition, incrementing i for each loop. i is a worker and this loop creates 3 workers which start working concurrently
		wg.Add(2)
		go worker(i, &wg, i, "TypeA")
		go worker(i, &wg, i, "TypeB")
		wg.Wait()
		fmt.Println(time.Now(), "non-worker", i) // & is a pointer to the WaitGroup, allowing the worker function to signal when it's done
	}

	wg.Wait()

	fmt.Println(time.Now(), "All workers finished")
}
