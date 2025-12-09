package main

import (
	"fmt"
	"os"
)

func main() {
	//creating a file

	file, err := os.Create("sample.txt")
	if err != nil {
		fmt.Print("Error creating file\n")
	} else {
		fmt.Print("File created\n")
	}
	file.Close()

	//opening a file

	filee, errr := os.OpenFile("sample.txt", os.O_RDWR, 0644)
	if errr != nil {
		fmt.Print("Error opening file\n")
	} else {
		fmt.Print("File Opened\n")
	}

	//writing to file
	filee.WriteString("This is a sample text\n")
	//read from file
	data, errr := os.ReadFile("sample.txt")
	if errr != nil {
		fmt.Print("cannot read")
	} else {
		fmt.Printf("The read data is: %s\n", data)
	}

	filee.Close()

}
