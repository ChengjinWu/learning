package _select

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func Test_22222222(t *testing.T) {
	totalCh := make(chan int, 1)
	//timer := time.NewTicker(10*time.Second)
	i := 0
	for {

		time.Sleep(2 * time.Second)
		go func(d int) {
			totalCh <- d
			fmt.Println(d)
			time.Sleep(3 * time.Second)
			<-totalCh
		}(i)
		i++
		fmt.Println("--->>>", i)
	}
}

func TimeOut(f func() interface{}) (v interface{}, err error) {
	var c chan interface{}

	go func() {
		c <- f()
	}()
	fmt.Println(time.Now().Unix())
	select {
	case v = <-c:
		return v, nil
	case <-time.After(6 * time.Second):
		fmt.Println(time.Now().Unix())
		return nil, errors.New("time out")
	}
}

func Test_select(t *testing.T) {
	for i := 0; i < 1000; i++ {
		result, err := TimeOut(func() interface{} {
			fmt.Println("start")
			time.Sleep(2 * time.Second)
			fmt.Println("end")
			return 7
		})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
	time.Sleep(9 * time.Second)
}
