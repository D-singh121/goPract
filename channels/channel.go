package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("This is all about the channels")
	//----two types of channel are present in golang 1. Buffer 2.Unbuffer
	//----buffer channels  has fixed size define whereas unbuffer channels has not any define fixed size.

	mychan := make(chan int, 1)
	wg := &sync.WaitGroup{} //pointer
	// mychan <- 2 // putting value into channel
	// num := <-mychan //-- accessing value from the channel.
	// fmt.Println(num)

	//----now it will show deadloc because channels works when someone is reading to the channel.
	//--- channels are bidirectional so they can send the value and recieve the value simultaneously
	//----- we can define channel in ways  either channel can only receive or channel can only send the values.
	//------  " ch chan <-" send only
	//------  " ch <-chan" Receive only

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		mychan <- 5
		close(mychan)

		wg.Done()
	}(mychan, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		// fmt.Println(<-mychan)
		val, isChannelOpen := <-mychan
		fmt.Println(val, isChannelOpen)
		wg.Done()
	}(mychan, wg)

	wg.Wait()
}
