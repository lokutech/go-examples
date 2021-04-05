package main

import (
	"fmt"
)

func main() {
	var aa int = 42
	var bb int = aa
	fmt.Println(aa, bb)
	aa = 27
	fmt.Println(aa, bb)

	// With dereferencing
	fmt.Println("\nWith dereferencing")
	var a int = 42
	var b *int = &a // *-pointer &-addressOfOperator
	//           & is referring to the memory location of a
	fmt.Println(a, *b)
	a = 27             // b changes as well
	fmt.Println(a, *b) // *-is in this case dereferencing operator; gets the value of that memory address
	pointerArray()
}

func pointerArray() {
	fmt.Println("\nPointer array")
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v %p %p\n", a, b, c)
}
