package main

import "fmt"

func main() {
	fmt.Println("all about Type assertion ")
	var i interface{} = "hello its a string"

	p := i.(string)
	fmt.Println(p)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
}
