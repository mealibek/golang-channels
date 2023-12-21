package main

import (
	"fmt"
	"sync"
	"time"
)

func dowork(d time.Duration, s string, ch chan string) {
	fmt.Println("Doing work....")
	time.Sleep(d)
	fmt.Println("Work is done...")
	ch <- s
	wg.Done()
}

var wg *sync.WaitGroup

func main() {
	start := time.Now()
	ch := make(chan string)
	wg = &sync.WaitGroup{}
	wg.Add(3)
	go dowork(time.Second*2, "work 1", ch)
	go dowork(time.Second*4, "work 2", ch)
	go dowork(time.Second*6, "work 3", ch)

	go func() {
		for res := range ch {
			fmt.Println(res)
		}
		close(ch)
	}()

	wg.Wait()
	fmt.Printf("Work took %v seconds\n", time.Since(start))
}
