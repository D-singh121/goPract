package main

//---for checking race condition in progarme run command   " go run --race file_name.go "
import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition starts")
	wg := &sync.WaitGroup{} // ---pointer
	mu := &sync.Mutex{}     //--pointer for handling the race condition while reading and writing to a shared location in memory by goroutines

	score := []int{0}
	wg.Add(3) //----numbers of goroutines running .
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("one Route")
		mu.Lock() //----locking and unlocking the  process of  appending the slice "score"
		score = append(score, 1)
		mu.Unlock()
		wg.Done()
	}(wg, mu)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("two Route")
		mu.Lock()
		score = append(score, 2)
		mu.Unlock()
		wg.Done()

	}(wg, mu)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("three Route")
		mu.Lock()
		score = append(score, 3)
		mu.Unlock()
		wg.Done()
	}(wg, mu)

	wg.Wait()
	fmt.Println(score)
}
