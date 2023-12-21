package main

import (
	"fmt"
	"sync"
	"time"
)

func dowork(d time.Duration, s string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done() // decrements groups by one
	fmt.Println("Doing work....", s)
	time.Sleep(d)
	fmt.Println("Work is done...", s)
	ch <- s
}

func main() {
	start := time.Now()
	ch := make(chan string, 3) // Create a buffered channel with a buffer size of 3
	wg := &sync.WaitGroup{}
	wg.Add(3)

	go dowork(time.Second*2, "work 1", ch, wg)
	go dowork(time.Second*4, "work 2", ch, wg)
	go dowork(time.Second*6, "work 3", ch, wg)

	wg.Wait() // waits until groups are finished...

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Printf("Work took %v seconds\n", time.Since(start))
	time.Sleep(time.Second)
}
