package main

import (
	"fmt"
)

func main() {
	//creation of slices

	// var days = []string{"monday", "tuesday", "wednesday", "thursday", "friday"}

	// fmt.Println(days)
	// days = append(days, "sunday")
	// fmt.Println(days[1:3])

	highscore := make([]int, 5)
	highscore[0] = 12
	highscore[1] = 22
	highscore[2] = 33
	highscore[4] = 45
	highscore[3] = 44

	fmt.Println(highscore[:2])
	highscore = append(highscore, 78, 65, 42, 38)
	fmt.Println(highscore)
	fmt.Println("Unsorted array")

	// sort.Ints(highscore) //----it will  increse our slice value in increasing oreder
	// fmt.Println(highscore)
	// fmt.Println("sorted array")
	// fmt.Println(sort.IntsAreSorted(highscore))

	// ---- now we can seprate slices into 2 or more part

	var index int = 2 //---- for removing the 2nd index value from the array .
	highscore = append(highscore[:index], highscore[index+1:]...)
	fmt.Println(highscore)

}
