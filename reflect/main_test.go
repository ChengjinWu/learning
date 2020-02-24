package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func prints(i int) string {
	fmt.Println("i =", i)
	return strconv.Itoa(i)
}

func Test_main(t *testing.T) {
	fv := reflect.ValueOf(prints)
	params := make([]reflect.Value, 1)                 //参数
	params[0] = reflect.ValueOf(20)                    //参数设置为20
	rs := fv.Call(params)                              //rs作为结果接受函数的返回值
	fmt.Println("result:", rs[0].Interface().(string)) //当然也可以直接是rs[0].Interface()
}

type MyType struct {
	i    int
	name string
}

func (mt *MyType) SetI(i int) {
	mt.i = i
}

func (mt *MyType) SetName(name string) {
	mt.name = name
}

func (mt *MyType) String() string {
	return fmt.Sprintf("%p", mt) + "--name:" + mt.name + " i:" + strconv.Itoa(mt.i)
}

func Test_22222222(t *testing.T) {
	myType := &MyType{22, "wowzai"}
	//fmt.Println(myType)     //就是检查一下myType对象内容
	//println("---------------")
	mtV := reflect.ValueOf(&myType).Elem()
	fmt.Println("Before:", mtV.MethodByName("String").Call(nil)[0])
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(18)
	mtV.MethodByName("SetI").Call(params)
	params[0] = reflect.ValueOf("reflection test")
	mtV.MethodByName("SetName").Call(params)
	fmt.Println("After:", mtV.MethodByName("String").Call(nil)[0])
}
