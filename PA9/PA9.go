package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"crypto/tls"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	cert, _ := tls.LoadX509KeyPair("server.cer", "server.key")
 	config := tls.Config{Certificates: []tls.Certificate{cert}}

	fmt.Println("Launching server...")
	ln, _ := tls.Listen("tcp", ":12002", &config)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		reader := bufio.NewReader(conn)      //使用 NewReader 读取文件
		req, err := http.ReadRequest(reader) //http.ReadRequest()是一個定義在net/http中，可以讀取textual input，然後存成一個"Request" object
		requestfile := req.URL.String()
		check(err)

		// fmt.Println(requestfile) // "/server-test.html"

		_, errf := os.Stat(requestfile[1:]) // os.Stat()直接讀取file name，然後return stat，如果此file找不到，errf會被assign *PathError

		if errf == nil {
			fmt.Printf("Method: %s\n", req.Method)
			// file_size := fmt.Sprint(f.Size()) // file_size
			// fmt.Println("File Size = ", file_size)
			content, err := ioutil.ReadFile(requestfile[1:])
			check(err)
			text := string(content)
			fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprintf(conn, "Date: ...\r\n")
			fmt.Fprintf(conn, "\r\n")
			fmt.Fprintf(conn, text+"\r\n")
			fmt.Fprintf(conn, "\r\n")

		} else {
			// fmt.Println("File Not Found")
			fmt.Fprintf(conn, "HTTP/1.1 404 Not Found\r\n")
			fmt.Fprintf(conn, "Date: ...\r\n")
			fmt.Fprintf(conn, "\r\n")
			fmt.Fprintf(conn, "File not found\r\n")
			fmt.Fprintf(conn, "\r\n")
			//fmt.Println(file_name)
			//check(errf)
		}
		conn.Close()

		//記得用curl 127.0.0.1:12002/server-test.html 來測試
		//127.0.0.1是local ip ，或是看你run server此時的ip address

	}
}
