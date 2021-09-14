package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)
	go doWork("A")
	go doWork("B")

	time.Sleep(time.Second)
	fmt.Println("shutdown")
	atomic.StoreInt64(&shutdown, 1)

	// need this because need the main goroutine to wait for two worker goroutine
	wg.Wait()
}

func doWork(s string) {
	defer wg.Done()
	for {
		fmt.Printf("%s start \n", s)
		time.Sleep(time.Millisecond * 250)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("shutting %s down", s)
			break
		}
	}
}
