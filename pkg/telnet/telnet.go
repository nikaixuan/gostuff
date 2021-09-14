package telnet

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func telnetServer(address string, exitChan chan int) {
	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}
	fmt.Println("listening to " + address)
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		go handleSession(conn, exitChan)
	}
}

func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("Session started: ")
	reader := bufio.NewReader(conn)

	for {
		str, err := reader.ReadString('\n')

		if err == nil {
			str = strings.TrimSpace(str)

			if !processCommand(str, exitChan) {
				conn.Close()
				break
			}

			conn.Write([]byte(str + "\r\n"))
		} else {
			fmt.Println(err)
			conn.Close()
			break
		}
	}
}

func processCommand(str string, exitChan chan int) bool {
	if strings.HasPrefix(str, "@close") {
		fmt.Println("session closed")
		return false
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("shutdown")

		exitChan <- 0
		return false
	}

	fmt.Println(str)
	return true
}

func main() {
	ch := make(chan int)
	go telnetServer("127.0.0.1:7090", ch)

	code := <-ch

	os.Exit(code)
}
