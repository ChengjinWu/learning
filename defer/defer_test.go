package _defer_test

import (
	"fmt"
	"testing"
)

func Test_111(t *testing.T) {
	fmt.Println("return:", a()) // 打印结果为 return: 0
}

func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i
}

func Test_3333333333(t *testing.T) {
	fmt.Println("return:", b()) // 打印结果为 return: 2
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i // 或者直接 return 效果相同
}

func Test_344444433(t *testing.T) {
	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}
func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}
func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
