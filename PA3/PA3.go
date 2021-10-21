package main
import "fmt"
import "bufio"
import "net"
import "os"

func check(e error) {
 if e != nil {
 panic(e)
 }
}
func main() {
	UploadFilename := ""
	fmt.Print("Input filename: ")
	fmt.Scan(&UploadFilename)

	conn, errc := net.Dial("tcp", "127.0.0.1:12000")
	check(errc)
	defer conn.Close()

	input, err := os.Open(UploadFilename)
	check(err)
	defer input.Close()

	reader := bufio.NewScanner(input)
	writer := bufio.NewWriter(conn)

	f, errf := input.Stat()
	check(errf)
	filesize := f.Size()
	fmt.Printf("Send the file size first: %d\n", filesize)

	writer.WriteString(fmt.Sprintf("%d",filesize))
	writer.WriteString("\n")

	for reader.Scan() {
		//fmt.Println(reader.Text())
		writer.WriteString(reader.Text())
		writer.WriteString("\n")
	}

	writer.Flush()

	scanner := bufio.NewScanner(conn)
 	if scanner.Scan() {
 		fmt.Printf("Server says: %s\n", scanner.Text())
 	}
}