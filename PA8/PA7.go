package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12002")
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		reader := bufio.NewReader(conn)      //使用 NewReader 读取文件
		req, err := http.ReadRequest(reader) //http.ReadRequest()是一個定義在net/http中，可以讀取textual input，然後存成一個"Request" object
		requestfile := req.URL.String()
		check(err)

		// fmt.Println(requestfile) // "/server-test.html"

		f, errf := os.Stat(requestfile[1:]) // os.Stat()直接讀取file name，然後return stat，如果此file找不到，errf會被assign *PathError

		if errf == nil {
			file_size := fmt.Sprint(f.Size()) // file_size
			fmt.Println("File Size = ", file_size)
		} else {
			fmt.Println("File Not Found")
			//fmt.Println(file_name)
			//check(errf)
		}
		conn.Close()

		//記得用curl 127.0.0.1:12002/server-test.html 來測試
		//127.0.0.1是local ip ，或是看你run server此時的ip address

	}
}
