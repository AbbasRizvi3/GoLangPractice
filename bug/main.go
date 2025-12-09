package main

import (
	_ "fmt"
	_ "math"
	_ "os"
)

// func a() {
// 	panic("something went wrong")
// }

// func b() {
// 	a()
// }

// func main() {
// 	b()
// }

// var mutex sync.Mutex
// var wg sync.WaitGroup

// func main() {
// 	m := make(map[string]string)
// 	wg.Add(1)
// 	go func() {
// 		mutex.Lock()
// 		defer mutex.Unlock()
// 		defer wg.Done()
// 		m["1"] = "a" // First conflicting access.
// 	}()
// 	mutex.Lock()
// 	m["2"] = "b" // Second conflicting access.
// 	mutex.Unlock()
// 	wg.Wait()
// 	for k, v := range m {
// 		fmt.Println(k, v)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(5)
// 	var i int
// 	for i = 0; i < 5; i++ {
// 		go func(j int) {
// 			fmt.Println(j) // Not the 'i' you are looking for.
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

// func main() {
// 	// creating a file

// 	file, err := os.Create("samplefile.txt")
// 	if err != nil {
// 		fmt.Print("Error creating file")
// 	} else {
// 		file.WriteString("hello")
// 	}

// }
