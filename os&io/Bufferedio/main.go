package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("Sample.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Print("Error opening or creating file")
	} else {
		file.Seek(0, 0)
		bfile := bufio.NewReaderSize(file, 8912)
		data, errr := bfile.ReadString('\n')
		if errr != nil {
			fmt.Println("error")
		} else {
			fmt.Println(data) // use loop for reading till eof
		}

	}

	wfile := bufio.NewWriterSize(file, 1024)
	wfile.WriteString("hey hello")
	wfile.Flush()
	wfile.WriteString("\n")
	wfile.WriteString("heyyy")
	wfile.Flush()

	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}

	defer file.Close()
}
