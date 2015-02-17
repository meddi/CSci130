package main 

import (
	"fmt"
	"errors"
	"math"
	"log"
)

type Salute struct {
	x, y int
	u float64 // u means unsigned
	_ float32 
	A *[]int
	s *string
}

// Function with multiple returns
func SumProdDiff(i, j int) (int, int, int) {
	return i+j, i*j, i-j
}

//
func MySqrt(f float64) (float64, error) {
	//return an error as second parameter in case of invalid input
	if (f < 0) {
		return float64(math.NaN()), errors.New("I won't be able to do sqrt of negative number!")
	}
	
	// otherwise use default square root function
	return math.Sqrt(f), nil
}

func Hi() (int, string) {
	return 1, "Hello."
}

// 
func Hello(input string) (strVal string, err error) {
	if len(input) == 0 {
		return "", errors.New("Empty strings are not accepted!")
	} else {
		return input + " is a proper string!", nil
	}
}

func main() {
	
	// Here we call the function with multiple returns
	sum, prod, diff := SumProdDiff(3, 4)
	fmt.Println("Sum: ", sum , " | Product: ", prod, " | diff: ", diff)
	
	// Multiple return values and error handling
	fmt.Print("First example with -1: ")
	ret1, err1 := MySqrt(-1)
	if err1 != nil {
		fmt.Println("Error! Return values are ", ret1, err1)
	} else {
		fmt.Println("It's Ok! return values are ", ret1, err1)
	}
	
	// Here to practice unused returns
	fmt.Print("Second example with 5: ")
	if ret2, err2 := MySqrt(5); err2 != nil {
		fmt.Println("Error! Return values are ", ret2, err2)
	} else {
		fmt.Println("It's Ok! Return values are ", ret2, err2)
	}
	
	_, strVal := Hi()
	fmt.Println("This is to practice unused return", strVal)
	
	if strVal, err := Hello("Mehdi"); err != nil {
		log.Fatal("You have entered an empty string!")
	} else {
		fmt.Println(strVal)
	}
	
}

