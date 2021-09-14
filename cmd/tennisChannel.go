package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wgt sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wgt.Add(2)
	court := make(chan int)

	go player("B", court)
	go player("A", court)

	court <- 0
	wgt.Wait()
}

func player(s string, ch chan int) {
	defer wgt.Done()
	for {
		n, ok := <-ch
		if !ok {
			fmt.Printf("%s win, close \n", s)
			return
		}
		ra := rand.Intn(100)
		if ra%7 == 0 {
			close(ch)
			fmt.Printf("%s lost, close, ball %v \n", s, ra)
			return
		}
		n += 1
		fmt.Printf("%s hit, round %v, ball %v \n", s, n, ra)
		ch <- n
	}
}
