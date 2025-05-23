package interface_writer

import (
	"bufio"
	"os"
)

func main() {
	_, _ = os.Stdout.Write([]byte("Hello, World!!!"))

	f, err := os.OpenFile("out.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	_, _ = f.Write([]byte("Hi"))
	_, _ = f.Write([]byte(" "))
	_, _ = f.Write([]byte("there"))
	_, _ = f.Write([]byte("!!!"))

	//listener, err := net.Listen("tcp", ":8080")
	//if err != nil {
	//	panic(err)
	//}
	//defer listener.Close()
	//
	//for {
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	buf := make([]byte, 1024)
	//
	//	n1, err := conn.Read(buf)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println("Income: ", string(buf[:n1]))
	//
	//	n2, err := conn.Write(buf[:n1])
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println("Outcome: ", string(buf[:n2]))
	//
	//	conn.Close()
	//}

	writer := bufio.NewWriter(os.Stdout)

	writer.Write([]byte("Hello"))
	writer.Flush()
	writer.Write([]byte(", "))
	writer.Write([]byte("World"))
	writer.Write([]byte("!!!")) // break
	writer.Flush()
}
