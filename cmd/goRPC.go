package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go rpcServer(c)

	response, err := rpcClient(c, "2")
	if err == nil {
		fmt.Println(response)
	}
}

func rpcClient(ch chan string, req string) (string, error) {
	ch <- req

	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(3 * time.Second):
		fmt.Println("time out")
		return "", fmt.Errorf("time out")
	}
}

func rpcServer(ch chan string) {
	for {
		data := <-ch
		fmt.Printf("received %s \n", data)

		ch <- "roger"
	}
}
