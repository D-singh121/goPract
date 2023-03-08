package main

import "fmt"

type Calculation interface {
	calculate() int
	displayResult()
}
type sum struct {
	num1, num2 int
}

type multiply struct {
	num1, num2 int
}

func (s sum) calculate() int {
	return s.num1 + s.num2
}
func (m multiply) calculate() int {
	return m.num1 * m.num2
}

// display the sum and multiply
func (s sum) displayResult() {
	fmt.Printf("Sum = %d \n", s.calculate())
}

func (m multiply) displayResult() {
	fmt.Printf("Multiply = %d \n", m.calculate())
}

func main() {
	sm := sum{33, 45}
	mul := multiply{15, 15}

	cc := []Calculation{sm, mul}
	for _, c := range cc {
		c.displayResult()
	}
}
