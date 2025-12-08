package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var temp int
var mutex sync.Mutex
var rwMutex sync.RWMutex
var wg sync.WaitGroup

func sayHello() {
	fmt.Print("Hello")
	wg.Done()
}

func incrTemp() {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	temp += 1
}

func readTemp() int {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return temp
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
	wg.Done()
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
	wg.Done()
}

// func incrCounter() {
// 	mutex.Lock()
// 	counter = counter + 1
// 	mutex.Unlock()
// }

func incrCounter() {
	mutex.Lock()
	defer mutex.Unlock()
	counter += 1
	wg.Done()
}

func main() { // is already on a goroutine by default

	wg.Add(1)
	go sayHello()        // launches another goroutine but takes time to load so we add sleep
	fmt.Println("World") // prints out first

	wg.Add(1)
	go func() {
		fmt.Print("hey")
		wg.Done()
	}()
	fmt.Printf("\n")

	// if we run a simple loop, here the output will be 55555 as the loop might go forward and the gorouine isnt launced
	// reason being is that loops used values as a reference
	// a fix it to pass the value as a func parameter so it goes to stack for each loop func

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Printf("%d\n", i)
		}()
		wg.Done()
	}

	// it doesnt happen in this case but the fix can be
	print("------------------------------------------------")
	print("\n")
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Printf("%d\n", n)
			defer wg.Done()
		}(i)

	}

	print("------------------------------------------------")
	print("\n")

	wg.Add(2)
	go numbers()
	go alphabets()

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go incrCounter()

	}
	fmt.Printf("counter: %d\n", counter)

	print("------------------------------------------------")
	print("\n")

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			incrTemp()
			garbage := readTemp()
			fmt.Printf("Temp: %d\n", garbage)
		}()
	}

	wg.Wait()

}
