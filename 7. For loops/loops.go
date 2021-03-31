package main

import (
	"fmt"
)

func main() {
	// Simple for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// Same as above - only difference i is scoped to the main function
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	// Same as above
	for i := 0; i < 5; {
		fmt.Println(i)
		i++
	}
	// Same as above
	for i < 5 {
		fmt.Println(i)
		i++
	}
	// While loop
	for {
		fmt.Println(i)
		i++
		if i == 50 {
			break
		}
	}

	// Dual for loop
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	// Nested loop
	for i := 1; i <= 3; i++ {
	Loop: // This is a label to tell the loop where to break
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	// Lopping slices, arrays, maps, strings, channels
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
}
