package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	Addr = "localhost:8000"
)

func main() {
	listen, err := net.Listen("tcp", Addr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("Receive error: %s\n", err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	message := make(chan string, 1)
	defer conn.Close()
	defer close(message)

	go write(conn, message)
	read(conn, message)
}

func read(conn net.Conn, message chan string) {
	reader := bufio.NewScanner(conn)
	for reader.Scan() { // 当对端断连，这里会退出
		content := reader.Text()
		// 忽略掉消息处理
		log.Printf("Receive message: %s\n", content)
		message <- content
	}
}

func write(conn net.Conn, message <-chan string) {
	for msg := range message { // 依赖 channel 关闭来关闭 goroutine
		fmt.Fprintln(conn, msg)
	}
}
