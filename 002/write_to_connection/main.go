package main

import (
	"fmt"
	"io"
	"log"
	"net"
)


func main() {
	li, err := net.Listen("tcp",":8080")
	if  err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for{
		conn, err := li.Accept()
		if err != nil{
			log.Println(err)
		}

		io.WriteString(conn, "\nHello from TCPserver\n")
		fmt.Fprintln(conn, "How  ir your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()
	}
}