package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// addr - сетевой адрес
const addr = "0.0.0.0:12345"

// protocol - протокол сетевой службы
const protocol = "tcp4"

func main() {
	listner, err := net.Listen(protocol, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listner.Close()

	fmt.Println("Listening on", protocol, addr)

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

// handleConn Обработчик соединения
func handleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	b, err := reader.ReadBytes('\n')
	if err != nil {
		log.Println(err)

		return
	}

	msg := strings.TrimSuffix(string(b), "\n")
	msg = strings.TrimSuffix(msg, "\r")
	if msg == "time" {
		_, err = conn.Write([]byte(time.Now().String() + "\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}
