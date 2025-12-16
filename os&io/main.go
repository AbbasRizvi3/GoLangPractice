package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// os is used for file handling but io is used as input and output on anything
	// io can work with files, stdin, stdout, ds, etc
	// io is an interface while os implements that interface to provide operations, io can be used directly too

	// file3, err3 := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0644)
	// if err3 != nil {
	// 	println("Error opening or creating file")
	// }
	// io.Copy(file3, os.Stdin)
	// defer file3.Close()
	// this is make whatever u type into terminal to be shown in test.txt

	file, err := os.OpenFile("File.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		println("Error opening or creating file")
	}
	defer file.Close()
	file.Write([]byte("Helllo\n testing 1 2 3"))
	file.Seek(0, 0)
	buf := make([]byte, 100)
	size, err := file.Read(buf)
	if err == nil {
		fmt.Printf(string(buf[:size]))
	} else {
		fmt.Printf("error")
	}

	file.Seek(0, 0)
	buf2 := make([]byte, 100)
	size2, err2 := io.ReadFull(file, buf2)
	if err2 == nil {
		fmt.Printf("Error readin from file")
	} else {
		fmt.Println(string(buf2[:size2]))
	}
	file.Seek(0, 0)

	f, e := os.OpenFile("Sample.txt", os.O_RDWR|os.O_CREATE, 0644)
	if e != nil {
		fmt.Println("Error creating or opening Sample.txt file")
	} else {
		defer f.Close()
		io.Copy(f, file)
		f.Seek(0, 0)
		data, er := io.ReadAll(f)
		if er == nil {
			fmt.Println("showing copied data")
			fmt.Println(string(data))

		}
	}
}
