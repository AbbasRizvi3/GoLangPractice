package main

import (
	"fmt"
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

// type TEMP int
// type TEMP2 = int
// type Human interface {
// 	getName() string
// }

// type Person struct {
// 	Name string
// }

// func (p Person) getName() string {
// 	return p.Name
// }

// var temp int = 5

// const garbage = 3

func DoSomething() {
	fmt.Println("hello")
}

func do_something_bad() { // <- bad naming style (snake_case)
	fmt.Println("bad")
}

func main() {
	// arr := [10]int{1, 2, 3, 4, 5}
	// // var arr2 [10]int
	// arr[0] = 1
	// fmt.Println(arr)

	// var sl []int
	// fmt.Println(sl)

	// sl = append(sl, 2)
	// sl = append(sl, 3)
	// sl = append(sl, 5)
	// fmt.Println(sl)

	// sl = append(sl[:], 1)
	// fmt.Println(sl)
	// fmt.Println(temp)

	// fmt.Println(garbage)

	// const (
	// 	a int   = iota
	// 	b int32 = iota
	// 	c int64 = iota
	// 	d
	// )
	// fmt.Println(a, b, c, d)

	// const (
	// 	e = iota
	// 	f
	// 	g
	// 	h
	// )

	// fmt.Print(e, f, g, h)

	// // var gg TEMP=5
	// // var gg2 TEMP2=6
	// // var gg3 int=7

	// // gg2+=gg3
	// garbage := 5
	// fmt.Println(garbage)

	// var (
	// 	x int    = 5
	// 	y string = "sd"
	// )
	// fmt.Println(x, y)

	// // var temppp uint = -7

	// var c1 complex64 = complex(5, 7)
	// var c2 complex64 = complex(5, 7)
	// c3 := c1 * c2
	// fmt.Println(c3)

	// arrr := []int{1, 2, 3, 4, 5, 6, 6, 7, 7, 3}
	// arrr = append(arrr, 0)
	// copy(arrr[7:], arrr[6:])
	// arrr[6] = 9
	// fmt.Println(arrr)

	var x int
	fmt.Println(x)
	var y string
	fmt.Println(y)

	var sli = make([]int, 5, 10)
	sli = append(sli, 5, 4, 6, 2, 2, 1)
	// sli[6] = 3
	fmt.Println(sli)
	var sli2 = make([]bool, 5)
	// sli2[6] = 3
	fmt.Println(sli2)

	var arr [5]bool
	fmt.Println(arr)

	var s []int
	fmt.Println(s)

	var xx [5]int
	fmt.Println(xx)

	temp := new(int)
	fmt.Println(*temp)

	println("--------------------------------")
	var tem interface{}
	tem = 4
	fmt.Println(tem)

	fmt.Printf("%d\n", "hey")
	ttst := "abbas"
	ttst2 := "feeehu"
	fmt.Printf("This is %s\n", ttst, ttst2)

	test()

	string := "sdad"
	fmt.Print(string)

}
