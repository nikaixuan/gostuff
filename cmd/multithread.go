package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func producer(header string, channel chan<- string) {
	for {
		channel <- fmt.Sprintf("%s : %v", header, rand.Int31())
		time.Sleep(time.Second)
	}
}

func consumer(channel <-chan string) {
	for {
		message := <-channel
		fmt.Println(message)
	}
}

var count int
var rwLock sync.RWMutex

func main() {
	//// channel test
	//channel := make(chan string)
	//go producer("123", channel)
	//go producer("234", channel)
	//consumer(channel)
	//
	//// lock test
	//var a = 0
	//var lock sync.Mutex
	//for i:=0;i<10;i++ {
	//	go func(idx int) {
	//		lock.Lock()
	//		defer lock.Unlock()
	//		a = a+1
	//		fmt.Printf("%v is %v", i, a)
	//		fmt.Println()
	//	}(i)
	//}
	//
	//ch := make(chan string)
	//go func() {
	//	lock.Lock()
	//	defer lock.Unlock()
	//	fmt.Println("lock 2s")
	//	time.Sleep(time.Second*2)
	//	fmt.Println("finish lock")
	//	ch<-"l1"
	//}()
	//go func() {
	//	fmt.Println("waiting for lock")
	//	lock.Lock()
	//	defer lock.Unlock()
	//	fmt.Println("finish lock 2")
	//	ch<-"l2"
	//}()
	//for i:=0;i<2;i++ {
	//	<-ch
	//}

	// Test read and write lock
	chann := make(chan struct{}, 10)
	for i := 0; i < 5; i++ {
		go read(i, chann)
	}
	for i := 0; i < 5; i++ {
		go write(i, chann)
	}
	for i := 0; i < 10; i++ {
		<-chann
	}

	// check CPU using
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// Read and write lock
func read(n int, ch chan struct{}) {
	rwLock.Lock()
	fmt.Printf("goroutine %v, can read, other can read, cannot write \n", n)
	v := count
	fmt.Printf("goroutine %v, read value: %v \n", n, v)
	rwLock.Unlock()
	ch <- struct{}{}
}

func write(n int, ch chan struct{}) {
	rwLock.Lock()
	fmt.Printf("goroutine %v, can write, other cannot write or read \n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("goroutine %v, write value: %v \n", n, count)
	rwLock.Unlock()
	ch <- struct{}{}
}
