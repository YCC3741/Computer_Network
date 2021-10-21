package main

import (
	"bufio"
	"fmt"
	"net"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	conn, errc := net.Dial("tcp", "127.0.0.1:12002")
	check(errc)
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	len, errw := writer.WriteString("Hello World!\n") // WriteString() contains the new line
	check(errw)

	fmt.Printf("Send a string of %d bytes\n", len)
	writer.Flush()

	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		fmt.Printf("Server replies: %s\n", scanner.Text())
	}
}
