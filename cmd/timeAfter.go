package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	exit := make(chan int)
	fmt.Println("start")
	time.AfterFunc(time.Second, func() {
		fmt.Println("one second after")
		exit <- 0
	})
	<-exit
	cpu := runtime.NumCPU()
	fmt.Println(cpu)
	runtime.GOMAXPROCS(cpu)
}
