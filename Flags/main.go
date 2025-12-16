package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	// the simple use of flags returns  a pointer o * is needed when using the pointer variable
	// name := flag.String("name", "Guest", "Enter your name")
	// age := flag.Int("age", 0, "Enter your age")
	// flag.Parse()
	// fmt.Printf("name: %s\n", *name)
	// fmt.Printf("age: %d\n", *age)

	var name2 string
	var age2 int
	// the addition of var makes a flag variable and so no pointer dereference is required
	flag.StringVar(&name2, "name", "guest", "Enter your name")
	flag.IntVar(&age2, "age", 0, "Enter your age")

	b := flag.Bool("check", false, "enter bool type")

	flag.Parse()
	fmt.Printf("%s\n", name2)
	fmt.Printf("%d\n", age2)
	fmt.Println("%v\n", *b)
	fmt.Print(flag.Args())
	fmt.Print(flag.Arg(0))

	// now lets try flagset, to add flags to a flagset, we use os.flag which can see all flags defined in cmd line
	// we cannot use args as it can only see args, not flags

	println("-------------------------------------------")
	ctime := time.Now()
	println(ctime.Date())
	println(ctime.Day())
	println(ctime.Hour())
	println(ctime.Month())

	println("---------------------------------------------")
	currentTime := time.Now()
	fmt.Println("The time is", currentTime)

	fmt.Printf("%d-%d-%d %d:%d:%d\n",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second())

}
