package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi")
	sayMessage("Hello go")

	for i := 0; i < 5; i++ {
		sayMessage2Params("Double params", i)
	}

	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name)
	fmt.Println(name)

	fmt.Println("\nSum function with spread")
	sum(1, 2, 3, 4, 5)

	// Anonymous function
	func() {
		fmt.Println("\nAnonymous functoin")
	}()

	fmt.Println("\nSum function with error return")
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

}

// Simple function
func sayMessage(msg string) {
	fmt.Println(msg)
}

// Function with 2 parameters
func sayMessage2Params(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

// Function with pointers
func sayGreeting(greeting, name *string) { // if type is same, can be written only at the end
	fmt.Println(*greeting, *name)
	*name = "Ted" // name gets mutated outside the scope because of the pointer
	fmt.Println(*name)
}

// Function with variatic parameters
func sum(values ...int) { // sum(msg string, values ...int) - this is also ok. There can be one variatic only at the end
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is ", result)
}

// Function with err in return
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
