package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup // pointer

// waitgroup has 3 function
// 1.- wg.Add()
// 2.- wg.Wait()
// 3.- wg.Done()
var mu sync.Mutex // pointer

//-- mu.Lock()
//--mu.UnLock()

func main() {
	website := []string{
		"https://www.google.com",
		"https://www.pkg.go.dev",
		"https://www.fb.com",
		"https://www.github.com",
	}

	for _, web := range website {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println(" sorry we did't get the endpoint")
	} else {
		mu.Lock() //---for protecting by race condition
		signals = append(signals, endpoint)
		defer mu.Unlock()

		fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
	}

}
