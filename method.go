package main

import "fmt"

type User struct {
	ID      int
	Name    string
	Email   string
	Roll_No int
}

func (u User) Visit() {
	fmt.Println(u.Roll_No)
}

func main() {
	Dee := User{2, "devesh", "choudhary11@gmail.in", 121}
	fmt.Println(Dee)
	Dee.Visit()
}
