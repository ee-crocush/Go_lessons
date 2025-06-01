package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// addr - сетевой адрес
const addr = "localhost:12345"

// protocol - протокол сетевой службы
const protocol = "tcp4"

func main() {
	conn, err := net.Dial(protocol, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Printf("Connected to %s\n", addr)

	for {
		_, err = conn.Read(make([]byte, 0))
		if err != nil {
			log.Fatal(err)
		}

		reader := bufio.NewReader(conn)
		b, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Client get: ", string(b))
	}

}
