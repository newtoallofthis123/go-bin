package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	names := []string{"John", "Paul", "George", "Ringo"}
	firsts := []string{}
	for index := range names {
		var first = names[index]
		fmt.Printf("Adding %s to DB\n", first)
		firsts = append(firsts, first)
		fmt.Printf("Firsts: %v\n", firsts)
	}
}
