package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("File.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		println("An error occured")
	}
	defer file.Close()

	data := make([]byte, 100)
	copy(data, "dadadas")
	file.Write(data)
	data2 := make([]byte, 100)
	file.Seek(0, 0)
	count, err := file.Read(data2)
	fmt.Println(string(data2[:count]))
}
