package main

import (
	"fmt"
	"time"
)

func greetings(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Millisecond)
		fmt.Println(s)

	}
}

func main() {
	go greetings("welcome") // ---- went to goroutine call
	greetings("To India")
}
