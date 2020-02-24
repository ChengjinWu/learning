package cpu

import (
	"testing"
	"time"
)

func Test_CPU(t *testing.T) {
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				break
			default:
			}
		}
	}()

	time.Sleep(time.Second * 150000)
	quit <- true
}
