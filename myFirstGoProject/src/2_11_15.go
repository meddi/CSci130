package main 

import (
	"fmt"
)

// Defining constants
const (
	win = 100
	gamesPerSeries = 10
)

type convert func(int) string

func value(x int) string {
	return fmt.Sprintf("%v", x)
}

func quote(fn convert) string {
	return "" + fn(123) + ""
}

// Function that return a function
func plus() (func(v int) int) {
	return func(v int) int {
		return v+2
	}
}

// Another practical use of function returing function
func plusX(x int) (func(v int) int) {
	return func(v int) int {
		return v+x
	}
}

func assert(cond bool, msg string) {
	if !cond {
		print("assertion failed: ", msg, "\n")
		panic(1)
	}
}


func main() {
	var result string
	
	result = value(123)
	fmt.Println(result)
	
	result = quote(value)
	fmt.Println(result)
	
	result = quote(func(x int) string { return fmt.Sprintf("%b", x) })
	fmt.Println(result)
	
	foo := func(x int) string { return "foo" }
	result = quote(foo)
	fmt.Println(result)
	
	//
	_ = convert(foo)
	
	p := plus()
	fmt.Println("3+2 is: ", p(3))
	
	px := plusX(3)
	fmt.Println("3+100 is: ", px(win))
	
	i5 := 5

	switch true {
	case i5 < 5:
		assert(false, "<")
	case i5 == 5:
		assert(true, "!")
	case i5 > 5:
		assert(false, ">")
	}	
	
	switch x:=5; true {
	case i5 < x:
		assert(false, "<")
	case i5 == x:
		assert(true, "!")
	case i5 > x:
		assert(false, ">")
	default:
		assert(false, "default")
	}
	
}

