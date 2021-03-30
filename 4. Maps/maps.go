package main

import "fmt"

func main() {
	// MAPS
	// Similar to objects in Javascript
	// Slices, maps, and other functions can't be keys to maps
	statePopulations := map[string]int{ // Literal syntax for declaring maps
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	statePopulations2 := make(map[string]int) // Declaring map with make
	statePopulations2 = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	sp := statePopulations
	delete(sp, "Illinois") // Deleting from sp also deletes from statePopulations

	statePopulations["Georgia"] = 10310371     // Add a value to map
	delete(statePopulations, "Ohio")           // Delete a value
	checkAL, ok := statePopulations["Alabama"] // CheckAL returns 0 which is confusing, ok returns true or false
	_, ok2 := statePopulations["Alabama"]      // Alternative to above
	fmt.Println(statePopulations)              // Map returns aren't in order as in Arrays and Slices
	fmt.Println(statePopulations2)
	fmt.Println(statePopulations["California"])
	fmt.Println("Check Alabama: ", checkAL, ok)
	fmt.Println("Check Alabama: ", ok2)
	fmt.Println("Map length: ", len(statePopulations))
}
