package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	numChan, resChan := make(chan int), make(chan int)
	var wg sync.WaitGroup
	
	rand.NewSource(time.Now().UnixNano())
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rand.NewSource(time.Now().UnixNano())
			num := rand.Intn(101)
			numChan <- num
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := <-numChan
			resChan <- result * result
		}()
	}
	go func() {
		wg.Wait()
		close(resChan)
		close(numChan)
	}()
	for square := range resChan {
		fmt.Println(square)
	}
}
