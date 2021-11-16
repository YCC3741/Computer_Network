package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	original_size := ""
	if scanner.Scan() { // Scan() ignores the new line
		original_size = scanner.Text()
	}

	fmt.Printf("Upload file size: %s\n", original_size)
	f, err := os.Create("./whatever.txt")

	check(err)
	//defer f.Close()

	writer := bufio.NewWriter(f)
	conn_writter := bufio.NewWriter(conn)
	//reader := bufio.NewReader(conn)

	line_counter := 1
	num := 0
	all, _ := strconv.Atoi(original_size)
	for scanner.Scan() {
		//fmt.Printf(strconv.Itoa(line_counter))

		newline := strconv.Itoa(line_counter) + " " + scanner.Text() + "\n"
		_, errw := writer.WriteString(newline)
		check(errw)
		writer.Flush()

		num += len(scanner.Text())
		//print(line_counter, " ", num, " ", all, "\n")
		if num == all-line_counter {
			break
		}
		line_counter++
		//fmt.Printf(strconv.Itoa(line_counter))
	}

	//fmt.Printf(strconv.Itoa(1))
	//f, err = os.Open("./whatever.txt")

	file_status, err_stat := f.Stat() // for getting information of file
	check(err_stat)
	file_size := fmt.Sprint(file_status.Size()) // file_size
	//fmt.Println(file_size)

	return_message := fmt.Sprintf("%s bytes received, %s bytes file generated", original_size, file_size)
	_, errwf := conn_writter.WriteString(return_message)
	check(errwf)
	conn_writter.Flush()

	fmt.Printf("Output file size: %s\n", file_size)
	time.Sleep(5 * time.Second)
	conn.Close()

}
func main() {
	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", ":12002")
	defer ln.Close()

	for {
		conn, _ := ln.Accept()

		go handleConnection(conn)

	}
}
