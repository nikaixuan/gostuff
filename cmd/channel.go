package main

import (
	"fmt"
	"math/rand"
	"time"
)

var nums []int
var ch = make(chan []int)

func main() {
	//go receiveNumber()
	//go addNumber()
	//time.Sleep(time.Second*10)

	c := make(chan int)
	go printer(c)

	for i := 1; i < 11; i++ {
		c <- i
	}

	// inform goroutine no data is sent
	c <- 0

	// receive goroutine is end
	<-c
}

func addNumber() {
	time.Sleep(time.Second * 3)
	for i := 0; i < 3; i++ {
		num := rand.Intn(10)
		nums = append(nums, num)
		fmt.Println("add number")
		ch <- nums
		fmt.Println("add finish")
	}
}

func receiveNumber() {
	fmt.Println("blocking and wait")
	for i := range ch {
		fmt.Println(i)
	}
}

func printer(ch chan int) {
	for {
		data := <-ch
		// receive no data
		if data == 0 {
			break
		}
		fmt.Println(data)
	}
	// inform goroutine is end
	ch <- 0
}
