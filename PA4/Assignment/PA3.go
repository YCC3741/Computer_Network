package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// connecting part
	conn, errc := net.Dial("tcp", "127.0.0.1:12002")
	check(errc)
	defer conn.Close()

	// producing the file
	fmt.Print("Please input the filename: ")

	text := ""
	fmt.Scanln(&text)

	f, err := os.Open(text)
	check(err)
	defer f.Close()

	// read file size
	file_status, err_stat := f.Stat() // for getting information of file
	check(err_stat)
	file_size := fmt.Sprint(file_status.Size()) // file_size

	// sent file size
	writer := bufio.NewWriter(conn)
	_, errw := writer.WriteString(file_size)
	check(errw)
	_, errw = writer.WriteString("\n") // remember to add "\n"
	check(errw)
	writer.Flush()

	fmt.Println("Send the file size first: " + file_size)

	// sent file contains
	scanner_for_file := bufio.NewScanner(f) // for file contains

	for scanner_for_file.Scan() {
		_, errw = writer.WriteString(scanner_for_file.Text()) // WriteString() contains the new line
		check(errw)
		_, errw = writer.WriteString("\n")
		check(errw)
	}
	writer.Flush() // sent file contains

	// receive the message from server
	scanner_for_server := bufio.NewScanner(conn) // for recepting the message from server
	if scanner_for_server.Scan() {
		fmt.Println("Server says:", scanner_for_server.Text())
	}

}
