package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch1)
	}()
	ch1 <- 1
	close(ch1)
	v, ok := <-ch1
	fmt.Printf("%v %v\n", v, ok)
	wg.Wait()
}
