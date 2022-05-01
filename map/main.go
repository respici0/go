package main

import "fmt"

// a maps keys & values must be all the same types
func main() {
	// more than one way to declare a map in go

	// 1.
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// 2. Usual way of declaring a map, with different syntax
	// var colors map[string]string
	// colors := make(map[string]string) // same as line 16 just a different way to declare a map
	// colors["white"] = "#fff" // there is no dot notation with maps .. only square brace syntax - this is the way to create a new type in map

	// delete(colors, "white") // delete keyword to delete items in a map

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c { // color, hex can be seen as key, value
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// Difference between maps & structs
// Map
// 1. All keys must be the same type
// 2. All values must be the same type 
// 3. Keys are indexed - we can iterate over them
// 4. Use to represent a collection of related properties
// 5. Don't need to know all the keys at compile time
// 6. Reference Type!

// Struct 
// 1. Values can be of different types
// 2. keys don't support indexing
// 3. You need to know all the different fields at compile time
// 4. Use to represent a "thing" with a lot of different properties
// 5. Value Type!