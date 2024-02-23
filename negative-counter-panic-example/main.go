package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer wg.Done()
		time.Sleep(time.Second)
	}()
	wg.Wait()
	fmt.Println("operation finish")
}
