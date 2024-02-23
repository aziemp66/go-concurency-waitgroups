package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go work(wg)

	wg.Wait()
	println("process done")
}

func work(wg sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
}
