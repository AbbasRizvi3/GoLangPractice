package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"rsc.io/quote"
) //external library, use go add . to download all dependencies or go add <dependency-name>

// Demonstration of use of variables, print function, data types, etc

func main() {
	var number int = 5
	fmt.Println(number)
	fmt.Println(5 + 7)
	var name = "abbas"
	fmt.Println(name)

	age := 67
	//age="67" not allowed
	println(age) //library used mostly when coding
	print(age)   //builtin used for debugging mostly

	fmt.Println(quote.Go())

	fmt.Printf("the name is %s and he is %d years old\n", name, age)

	// signed and unsigned numbers, ranges are diff line int8,32 etc and on exceeding the limit, it is wrapped around
	var signed_num int8 = 127
	fmt.Printf("the unsigned number is %d", signed_num)
	signed_num++
	fmt.Printf("the unsigned number is %d\n", signed_num)

	//floating numbers 32 and 64 bit
	var float1 float64 = 3.2145454541545452343232
	println(float1)
	var float2 float32 = float32(float1)
	println(float2)
	//floating numbers
	var complex_number complex128 = complex(10, 12)
	fmt.Printf("the complex number is: %v\n", complex_number)
	fmt.Printf("the real part is %v\n", real(complex_number))
	fmt.Printf("The imaginary part is %v\n", imag(complex_number))

	complex_number2 := complex(5, 7)
	complex := complex_number + complex_number2
	fmt.Printf("%v\n", complex)

	//parsing strings into bool
	bool_value := "1"
	ischeck, _ := strconv.ParseBool(bool_value)
	fmt.Printf("%v\n", ischeck)

	//parsing booleans into strings
	check := false
	var boolstring string = strconv.FormatBool(check)
	fmt.Printf("the converted string boolean value is %s\n", boolstring)

	if 7 > 5 && 9 > 3 {
		fmt.Print("greater\n")
	} else {
		fmt.Print("lesser\n")
	}

	var temp bool
	fmt.Print(temp)

	const pi float32 = 3.14
	const (
		x = 5
		y = 6
		z = 7
	)
	fmt.Print(x, y, z)
	const tempoo = 5
	const (
		Feeehu  = 5
		Abbas   = 6
		Mamoona = 7
		Roha    = 8
	)
	fmt.Print("\n")
	fmt.Print(Feeehu, Abbas, Mamoona, Roha)

	fmt.Print("\n")

	const (
		Feeehuu = "snake"
		Abbass
		Mamoonaaa
		Rohaaa
	)

	fmt.Printf("%s %s %s %s", Feeehuu, Abbass, Mamoonaaa, Rohaaa)
	print("\n")

	// fmt.Print("\n")
	// const temp2 = iota
	// fmt.Print(temp2)
	// fmt.Print("\n")

	// const (
	// 	Sunday = iota
	// 	Monday
	// 	Tuesday
	// 	Wednesday
	// 	Thursday
	// 	Friday
	// 	Saturday
	// )

	// fmt.Print(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	// fmt.Print("\n")

	// const (
	// 	KB = 1 << (iota * 10)
	// 	MB
	// 	GB
	// )
	// fmt.Print(KB, MB, GB)
	// fmt.Print("\n")

	// const (
	// 	offset = 2*iota + 1
	// 	_
	// 	value
	// 	_
	// 	result
	// )

	// fmt.Print(offset, value, result)
	// fmt.Print("\n")

	test := 10
	if true {
		test = 15
		fmt.Printf("%d", test)
	}
	fmt.Printf("\n")
	fmt.Printf("%d", test)

	fmt.Printf("\n")
	var charact rune = '4'
	fmt.Printf("%c\n", charact)

	s := "a\nb\nc"
	fmt.Printf("%q", s)

	var s2 string = "fareeha snake hai"
	fmt.Printf("\n")
	fmt.Printf("%d\n", len(s2))

	// %d is used for int, %s for strings, %v for raw, %q for adding quotationg, %c for rune/character, and %x for hexa value

	var agee int = 30
	var string_age string = strconv.Itoa(agee)
	fmt.Printf("%s\n", string_age)

	sample_str := "hello 你好"
	fmt.Printf("%d\n", utf8.RuneCountInString(sample_str))
	// simple len counts the number of bytes but string contains unicode and so can mismatch the calculations

}
