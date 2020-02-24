package context

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

// rang channel类型的WithCancel示例
// 数字生成器，通过遍历channel来获取通道数据
func Test_WithCancel(t *testing.T) {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func Test_WithDeadline(t *testing.T) {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func Test_WithTimeout(t *testing.T) {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	go func() {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err()) // prints "context deadline exceeded"
		}
	}()
	cancel()
	time.Sleep(time.Second)
}

// 简单地
func Test_WithValue(t *testing.T) {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))
}

func Test_WithTimeoutMulti(t *testing.T) {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Microsecond)
	var input, output int64
	ticker := time.NewTicker(time.Second)
	cancel()
	for {
		go func() {
			atomic.AddInt64(&input, 1)
			select {
			case <-time.After(1 * time.Second):
				fmt.Println("overslept")
			case <-ctx.Done():
				atomic.AddInt64(&output, 1)
				//fmt.Println(ctx.Err()) // prints "context deadline exceeded"
			}
		}()
		select {
		case <-ticker.C:
			fmt.Println(atomic.LoadInt64(&input), atomic.LoadInt64(&output))
		default:

		}
	}
	time.Sleep(time.Second)
	fmt.Println(input, output)
	time.Sleep(time.Second)
}
