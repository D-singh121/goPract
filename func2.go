package main

import "fmt"

func main() {
	//call by value vs pass by value
	x := 10
	y := 20

	// fmt.Println("before calling swap", x, y)
	// swap(x, y)
	// fmt.Println("after calling ", x, y)

	//--------------- call by reference
	fmt.Println("before calling swap", x, y)
	swap(&x, &y)
	fmt.Println("after calling ", x, y)

}

// ------------for call by vlaue
func swap(x, y int) {
	temp := x
	x = y
	y = temp
}

// --------------for call by refrence value
func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

//-----------pass by value and pass by reference
//-- pass by value ;= copy of variable is stored in the anthor location and it does not aaffect the orginal value  out of the function .
