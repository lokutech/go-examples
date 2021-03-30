package main

import (
	"fmt"
)

func main() {
	// ARRAYS
	// Array size doesn't change
	// Arrays can be any type but all elements must be same type
	grades := [3]int{10, 20, 40}                          // Assign an array
	grades2 := [...]int{5, 6, 7}                          // Assign an array way2
	var students [3]string                                // Assign an array with 0 values
	var nested [3][3]int                                  // Assign nested array
	nested2 := [3][3]int{{1, 0, 2}, {1, 0, 2}, {1, 0, 2}} // Assign nested array

	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Grades2: %v\n", grades2)
	fmt.Printf("Students: %v\n", students)
	students[0] = "Loku" // Specify an element value in an array
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[0])           // Display a single element of an array
	fmt.Printf("Number of Students: %v\n", len(students)) // Display a single element of an array
	fmt.Printf("Nested array: %v\n", nested)
	fmt.Printf("Nested array 2: %v\n", nested2)

	arr := [...]int{99, 88, 77}
	var barr = arr
	barr[0] = 55 // Does does mutate reference value arr, but barr = &arr mutates arr.
	fmt.Println(barr, arr)

	// SLICES
	// Slice size changes
	a := []int{1, 2, 3} // Assign a slice
	b := a[:]           // slice of all elements
	c := a[1:]          // slice from 2nd element to end
	d := a[:3]          // slice first 3 elements
	e := a[1:3]         // slice 2nd and 3rd elements
	fmt.Printf("Slice: %v\n", a)
	fmt.Printf("Slice length: %v\n", len(a))
	fmt.Printf("Slice capacity: %v\n", cap(a)) // Only available for slices
	fmt.Printf("b  %v\n", b)
	fmt.Printf("c  %v\n", c)
	fmt.Printf("d  %v\n", d)
	fmt.Printf("e  %v\n", e)

	sliceMake := make([]int, 3, 100) // making a slice with make (type, length, capacity)
	fmt.Printf("Slice with make %v\n", sliceMake)
	fmt.Printf("Length %v\n", len(sliceMake))
	fmt.Printf("Capacity %v\n", cap(sliceMake))

	// How append a slice
	sliceToAppend := []int{4, 5, 6}
	var appendedSlice = append(sliceToAppend, 7, 8, 9, 10)
	// var appendedSlice = append(sliceToAppend, int[]{7, 8, 9, 10}...) // Same thing with spread operator
	fmt.Printf("Appended slice %v\n", appendedSlice)

}
