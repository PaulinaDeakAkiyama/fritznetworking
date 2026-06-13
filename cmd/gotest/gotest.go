package main

import (
	"fmt"
	"sync"
	"time"
)

var globalstarttime string = time.Now().String()

func worker(id int, wg *sync.WaitGroup, waittime int, wtype string) {
	defer wg.Done()
	fmt.Println(time.Now(), "Worker", wtype, id, "start")
	time.Sleep(time.Duration(waittime) * time.Second)
	fmt.Println(time.Now(), "Worker", wtype, id, "end")
}

func main() {
	fmt.Println("Global Start time:", globalstarttime)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(2)
		go worker(i, &wg, i, "TypeA")
		go worker(i, &wg, i, "TypeB")
		wg.Wait()
		fmt.Println(time.Now(), "non-worker", i)
	}

	wg.Wait()

	fmt.Println(time.Now(), "All workers finished")
}
