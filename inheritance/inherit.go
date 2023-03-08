package main

import "fmt"

type Vehicle struct {
	Name  string
	Model int
}

type Car struct {
	Vehicle  //  embedding the struct vehicle
	Maxspeed int
	Variant  string
}

func main() {

	c := Car{}
	c.Name = "Tata harrior"
	c.Model = 2022
	c.Maxspeed = 150
	c.Variant = "petrol convertival"
	fmt.Println(c.Name)
	fmt.Println(c)

}
