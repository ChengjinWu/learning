package main

import (
	"fmt"
	"testing"
)

/*
append之后会new一个slice对象,内部对象也是申请的内存
原slice的长度，容量不变
*/
func Test_Context4444444444(t *testing.T) {
	slice := []int{10, 20, 30, 40}
	newSlice := append(slice, 50)
	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Slice[%d] = %p, newSlice[%d] = %p\n", i, &slice[i], i, &newSlice[i])
	}
	newSlice[1] += 10
	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))

}

func Test_Context22222222333(t *testing.T) {
	array := [4]int{10, 20, 30, 40}
	for i, _ := range array {
		fmt.Printf("index = %d, Pointer = %p\n", i, &array[i])
	}
	slice := array[1:3]
	for i, _ := range slice {
		fmt.Printf("index = %d, Pointer = %p\n", i, &slice[i])
	}
	newSlice := append(slice, 50)
	for i, _ := range newSlice {
		fmt.Printf("index = %d, Pointer = %p\n", i, &newSlice[i])
	}
	fmt.Printf("Before slice = %v, Pointer = %p, First = %p, len = %d, cap = %d\n", slice, &slice, &slice[0], len(slice), cap(slice))
	fmt.Printf("Before newSlice = %v, Pointer = %p, First = %p, len = %d, cap = %d\n", newSlice, &newSlice, &newSlice[0], len(newSlice), cap(newSlice))
	newSlice[1] += 10
	fmt.Printf("Before slice = %v, Pointer = %p, First = %p, len = %d, cap = %d\n", slice, &slice, &slice[0], len(slice), cap(slice))
	fmt.Printf("Before newSlice = %v, Pointer = %p, First = %p, len = %d, cap = %d\n", newSlice, &newSlice, &newSlice[0], len(newSlice), cap(newSlice))
	fmt.Printf("After array = %v\n", array)
}

func delete(array *[]int) {

	*array = append((*array)[0:5], (*array)[8:len((*array))]...)
	fmt.Println((*array), len((*array)), cap((*array)))
}

func Test_SliceValuePassing(t *testing.T) {
	array := []int{}
	for i := 0; i < 10; i++ {
		array = append(array, i)
	}
	delete(&array)
	fmt.Println(array, len(array), cap(array))

}

func Test_ArrayValuePassing(t *testing.T) {
	arrayA := [2]int{100, 200}
	var arrayB [2]int

	arrayB = arrayA

	fmt.Printf("arrayA : %p ,  %p , %v\n", &arrayA, &arrayA[0], arrayA)
	fmt.Printf("arrayB : %p ,  %p , %v\n", &arrayB, &arrayB[0], arrayB)

	testArray(arrayA)
}

func testArray(x [2]int) {
	fmt.Printf("func Array : %p , %v\n", &x, x)
}

func Test_ArrayQuotePassing(t *testing.T) {
	arrayA := [2]int{100, 200}
	fmt.Printf("arrayA : %p ,  %p , %v\n", &arrayA, &arrayA[0], arrayA)
	testArrayPoint(&arrayA) // 1.传数组指针
	//arrayB := arrayA[:]
	//testArrayPoint(&arrayB)   // 2.传切片
	//fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
}

func testArrayPoint(x *[2]int) {
	fmt.Printf("func Array : %p , %v\n", x, *x)
	(*x)[1] += 100
	x = &[2]int{12, 43}
}

func TestArrayPoint(t *testing.T) {
	s := []byte("")

	s1 := append(s, 'a')
	s2 := append(s, 'b')

	// 如果有此行，打印的结果是 a b，否则打印的结果是b b
	fmt.Println(s1, "===", s2)
	fmt.Println(string(s1), string(s2))
}
