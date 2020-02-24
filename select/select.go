package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(10 * time.Second) // sleep one second
		timeout <- true
	}()

	go func() {
		fmt.Println(<-timeout)
	}()
	ch := make(chan int)
	select {
	case <-ch:
	case <-timeout:
		fmt.Println("timeout!")
	}
}
