package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	// "strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12002")
	conn, _ := ln.Accept()
	defer ln.Close()
	defer conn.Close()


	scanner := bufio.NewScanner(conn)
	original_size := ""
	if scanner.Scan() { // Scan() ignores the new line
		original_size = scanner.Text()
	}

	fmt.Printf("Upload file size: %s\n", original_size)
	f, err := os.Create("./whatever.txt")

	check(err)
    defer f.Close()

	writer := bufio.NewWriter(f)
	reader := bufio.NewReader(conn)

	
	line_counter := 1
	// for ok:=scanner.Scan(); ok ; scanner.Scan() {
	// 	//fmt.Printf(strconv.Itoa(line_counter))
	// 	print(ok, line_counter, scanner.Text(), "\n")
	// 	newline := strconv.Itoa(line_counter) + " " + scanner.Text() + "\n"
	// 	_, errw := writer.WriteString(newline)
	// 	check(errw)
	// 	writer.Flush()
	// 	line_counter++
	// 	//fmt.Printf(strconv.Itoa(line_counter))
	// }
	for {
		line, err3 := reader.ReadString('\n')
		if err3 != nil {
			break
		}
		writer.WriteString(fmt.Sprint("%d %s", line_counter, line))
		line_counter++
	}
	

	//fmt.Printf(strconv.Itoa(1))
	f, err = os.Open("./whatever.txt")

	file_status, err_stat := f.Stat() // for getting information of file
	check(err_stat)
	file_size := fmt.Sprint(file_status.Size()) // file_size
	//fmt.Println(file_size)
	

	conn_writter := bufio.NewWriter(conn)
	return_message := fmt.Sprintf("%s bytes received, %s bytes file generatedfile_size", original_size, file_size)
	_, errwf := conn_writter.WriteString(return_message)
	check(errwf)
	conn_writter.Flush()
	fmt.Printf("Output file size: %s\n", file_size)

	




	
}
