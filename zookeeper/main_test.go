package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	wg = sync.WaitGroup{}
)

func print(exist <-chan string) {
	for {
		fmt.Println(44444444)
		var ok bool
		var result string
		select {
		case result, ok = <-exist:
			fmt.Println(ok)
			if result == "" {
				fmt.Println("无")
			} else {
				fmt.Println(result)
			}
			time.Sleep(time.Second)
		}

		if !ok {
			fmt.Println("break")
			break
		}
	}
	wg.Done()
}

func Test_chann(t *testing.T) {
	exist := make(chan string, 10)
	fmt.Println("enter")
	wg.Add(1)
	go func(ch chan<- string) {
		defer wg.Done()
		fmt.Println(222222222)
		ch <- "nihao"
		ch <- "nihao"
		ch <- "nihao"
		ch <- "nihao"
		ch <- "nihao"
		ch <- "nihao"
		fmt.Println(3333333)
		close(exist)
		fmt.Println(55555555555)
	}(exist)
	wg.Add(1)
	go print(exist)
	wg.Wait()
}
