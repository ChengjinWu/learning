package main

import "fmt"

func main() {
	var x chan int
	x = make(chan int)
	go func() {
		x <- 1
	}()
	fmt.Println(<-x)
}
