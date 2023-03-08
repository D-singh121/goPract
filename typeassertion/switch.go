package main

import (
	"fmt"
)

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n ", v, len(v))
	default:
		fmt.Printf("i don't know about types %T\n", v)

	}
}
func main() {
	fmt.Println("sitch type assertion")

	do(20)
	do("Devesh")
	do(true)
}
