package main

import "fmt"

// //------variadic function
// func sum(nums ...int) int {
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	return total
// }

// ------------------anonymous function
func solve(x, y int) (int, int) {
	return x / y, x % y
}

func main() {
	// =-----  therem is two type of function 1.variadic function .  2. Anonymous function .
	// ---varidic function denotes by ... ("triple dot ") ,it takes unlimited argument of same types.
	// fmt.Println(sum(2, 3, 5, 5))
	// fmt.Println(sum(2, 3, 5, 5, 4, 3))

	//------ -----------anonymous function
	//---we can ignore some value from multiple returning value
	_, mod := solve(5, 4)
	// fmt.Print("div is ", div, "\n")
	fmt.Print("mod is ", mod, "\n")

}
