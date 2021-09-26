package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		wg sync.WaitGroup
	)

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Println(i+1)
		}(i)
	}

	fmt.Println("wait...")
	wg.Wait()
	fmt.Println("wait ok!")


	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println("100")
		}()
	}
}
