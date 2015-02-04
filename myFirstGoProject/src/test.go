package main 

import (
	"fmt"
	"unsafe"
)

type Salute struct {
	x, y int
	u float64 // u means unsigned
	_ float32 
	A *[]int
	s *string
}

type myType struct {
	val1 int
	val2 string
}

func main() {
	
	x, y := 2, 5
	
	fmt.Println("X = ", x, "Y = ", y)
	
	swap(&x, &y)

	fmt.Println("X = ", x, "Y = ", y)
	
	var msg1 string
	msg1 = "First Declared, then Assigned!"
	fmt.Println(msg1)
	
	msg2 := "Declared and Assigned!"
	fmt.Println(msg2)
	
	msg3 := myType{10, "Bill"}
	fmt.Println(msg3.val2)
	
	ptr := unsafe.Pointer(&msg3)
	fmt.Printf("Addr: %v Value1 : %d Value2: %s\n", ptr, msg3.val1, msg3.val2)
	
	changeVal(&msg3)
	
	fmt.Printf("Addr: %v Value1 : %d Value2: %s\n", ptr, msg3.val1, msg3.val2)
}

func swap(x *int, y *int) {
	p := *x
	*x = *y
	*y = p
}

func changeVal(val *myType) {
	val.val1 = 20
	val.val2 = "Jill"
	
	ptr := unsafe.Pointer(&val)
	
	fmt.Printf("Addr: %v Value1: %d Value2: %s\n", ptr, val.val1, val.val2)
}
-------------------------OUTPUT----------------------------------------------------------
X =  2 Y =  5
X =  5 Y =  2
First Declared, then Assigned!
Declared and Assigned!
Bill
Addr: 0xc20801e040 Value1 : 10 Value2: Bill
Addr: 0xc20804c020 Value1: 20 Value2: Jill
Addr: 0xc20801e040 Value1 : 20 Value2: Jill
