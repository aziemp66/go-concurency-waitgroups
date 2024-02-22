package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go work(i + 1)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main is done")
}

func work(id int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task : ", id, " is done")
}
