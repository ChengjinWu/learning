package channel

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

// 关闭通道后，获取该通道信息的语句全部执行，取得默认值
func Test_CloseChannel(t *testing.T) {
	ch := make(chan int32, 1)
	var count int32 = 0
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(<-ch, atomic.AddInt32(&count, 1))

		}()
	}
	close(ch)
	time.Sleep(time.Second)
}

func Test_CloseChannel22(t *testing.T) {
	ch := make(chan int32)
	var count int32 = 0

	go func() {
		for i := 0; i < 1000000; i++ {
			select {
			case a, ok := <-ch:
				fmt.Println(a, ok, atomic.AddInt32(&count, 1))
			default:
				fmt.Println("nihoa")
			}
			time.Sleep(time.Second)
		}
	}()
	for i := 0; i < 100; i++ {
		ch <- int32(i)
		time.Sleep(time.Second)
	}
	//close(ch)
	time.Sleep(time.Second)
}
