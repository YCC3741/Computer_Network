package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8083")
	conn, _ := ln.Accept()
	defer conn.Close()
	defer ln.Close()

	writer := bufio.NewWriter(conn)
	_, errw := writer.WriteString("Input filename: ")
	check(errw)

	scanner := bufio.NewScanner(conn)

	message := ""
	if scanner.Scan() {
		message = scanner.Text()
		file, _ := os.Open(message)
		fi, _ := file.Stat()
		fsuze := fi.Size()
		fmt.Println(fsuze)
		b, err := ioutil.ReadFile(message) // just pass the file name
		if err != nil {
			fmt.Print(err)
		}
		ss := string(b)
		fmt.Println(ss)
	}
	// file, _ := os.Open(message)
	// fi, _ := file.Stat()
	// fsuze := fi.Size()

	// writer := bufio.NewWriter(conn)
	// to_user := fmt.Sprintf("Input filename: ")
	// writer.WriteString(to_user)

	writer2 := bufio.NewWriter(conn)
	newline := fmt.Sprintf("%d bytes received\n", len(message))
	writer2.WriteString(newline)
	writer2.Flush()
}
