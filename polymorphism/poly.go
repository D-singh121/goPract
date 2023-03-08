// package main

// import "fmt"

// ----- in golang polymorphism can be acheived with the help of Interfaces.
//----- interfaces  are the collection  of methods

// type Animal interface {
// 	speak()
// }

// type Cat struct {
// }

// type Dog struct {
// }

// type Cow struct {
// }

// func (c Cat) speak() {
// 	fmt.Println("meow")
// }

// func (d Dog) speak() {
// 	fmt.Println("woof")
// }
// func (co Cow) speak() {
// 	fmt.Println("mehhh")
// }
// func makeAnimalSpeak(a Animal) {
// 	a.speak()
// }

// func main() {
// 	c := Cat{}
// 	d := Dog{}
// 	co := Cow{}
// 	makeAnimalSpeak(c)
// 	makeAnimalSpeak(d)
// 	makeAnimalSpeak(co)

// }
