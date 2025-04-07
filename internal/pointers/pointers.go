package pointers

import "fmt"

func PointerBasic() {
	var a int = 42
	var b *int = &a // b is a pointer to int

	fmt.Printf("a: %v\n", a)   // 42
	fmt.Printf("b: %v\n", b)   // memory address of a
	fmt.Printf("*b: %v\n", *b) // 42

	*b = 100                                  // modify value at pointer
	fmt.Printf("a after *b = 100: %v\n\n", a) // 100
}
