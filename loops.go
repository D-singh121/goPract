package main

import "fmt"

// ------------ only for loop is available in golang , but we can impliment while loop also but syntax would be different.
func main() {
	// start := 1
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// fmt.Println(num)

	// for i := 0; i < len(num); i++ {
	// 	fmt.Println(num[i])
	// }

	//--------range loop
	for _, j := range num {
		fmt.Println(j)
	}

	//-----------for loop as a while loop

	// for start < 100 {
	// 	start += start
	// 	if start == 32 {
	// 		break
	// 	}
	// 	fmt.Println("value is ", start)
	// }

	//-----this also for while loop
	// j := 0
	// for j <= 5 {
	// 	fmt.Println(j)
	// 	j++
	// }

}
