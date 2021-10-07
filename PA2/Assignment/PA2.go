package main

import (
	"bufio"
	"fmt"
	"os"
)

func check_file(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input_name := ""
	fmt.Print("Please input the 'input filenames': ")
	fmt.Scan(&input_name)

	output_name := ""
	fmt.Print("Please input the 'output filenames': ")
	fmt.Scan(&output_name)

	input, err := os.Open(input_name)
	check_file(err)
	defer input.Close()

	output, err := os.Create(output_name)
	check_file(err)
	defer output.Close()

	counter := 1

	reader := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)

	for reader.Scan() {
		fmt.Fprintf(output, "%d ", counter)
		writer.WriteString(reader.Text())
		writer.WriteString("\n")
		writer.Flush()
		counter++
	}

}
