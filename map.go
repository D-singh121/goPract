package main

import (
	"fmt"
)

func main() {
	// creation of a map with make

	empsal := make(map[string]int)
	empsal["ram"] = 2000
	empsal["arjun"] = 3000
	empsal["raju"] = 7000
	empsal["das"] = 5000
	empsal["seeta"] = 8000
	empsal["sumi"] = 9000
	fmt.Println(empsal)
	fmt.Println(empsal["raju"])
	delete(empsal, "sumi")
	fmt.Println(empsal)

	for name, sal := range empsal {
		fmt.Printf("employee name is %v and salary is %v \n ", name, sal)
	}
}
