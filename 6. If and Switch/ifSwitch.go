package main

import (
	"fmt"
)

func main() {
	// Example 1
	if true {
		fmt.Println("hello")
	}

	// Example 2
	statePopulations := map[string]int{ // Literal syntax for declaring maps
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	// Example 3
	number := 50
	guess := 50
	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	}
	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
	}

	// SWITCH
	// Example 1
	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}

	// Example 2
	switch 5 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	// Example 3 - Tagless syntax - This one is allowed to overlap
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough // cancels break and continues to the next case.
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// Example 4 - Type switch
	var t interface{} = 1
	switch t.(type) {
	case int:
		fmt.Println("t is an int")
		break
		fmt.Println("some text before the break")
	case float64:
		fmt.Println("t is a float64")
	case string:
		fmt.Println("t is string")
	default:
		fmt.Println("t is another type")
	}

}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}
