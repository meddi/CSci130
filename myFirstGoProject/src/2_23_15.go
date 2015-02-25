package main 

import "fmt"
import "bytes"

type path []byte
type bin int
type StringRef []byte

func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
    if i >= 0 {
        *p = (*p)[0:i]
    }
}

func (b bin) String() string {
	return fmt.Sprintf("%b", b)
}

func (s StringRef) String() string {
	return string(s[:])
}

func main() {
	// One way to declare slice
	letters := []string{"a", "b", "c", "d", "e"}
	fmt.Println(letters)
	// another way is use a built-in function
	var s []byte
	s = make([]byte, 5, 5)
	fmt.Println(s)
	
	// Append to a slice
	letters = append(letters, "d")
	fmt.Println(letters)
	
	// Indexing in slices
	fmt.Println(letters[2:])


    pathName := path("/usr/bin/tso") // Conversion from string to path.
    pathName.TruncateAtFinalSlash()
    fmt.Printf("%s\n", pathName)
    
    // Method attached to string
    fmt.Println(bin(42))
    
    // Method with returns
    fmt.Printf("foo=%s", StringRef("bar"))

}

