package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// data := make([]byte, 100)
	file, err := os.ReadFile(os.Args[1]) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))
}
