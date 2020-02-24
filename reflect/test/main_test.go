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
	values := mtV.MethodByName("String").Call(nil)
	fmt.Println("Before:", len(values), values[0])
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(18)
	mtV.MethodByName("SetI").Call(params)
	params[0] = reflect.ValueOf("reflection test")
	mtV.MethodByName("SetName").Call(params)
	fmt.Println("After:", mtV.MethodByName("String").Call(nil)[0])
}
func Test_3333333333(t *testing.T) {
	myType := &MyType{22, "wowzai"}
	mtV := reflect.ValueOf(&myType).Elem()
	fmt.Println("Before:", mtV.Method(2).Call(nil)[0])
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(18)
	mtV.Method(0).Call(params)
	params[0] = reflect.ValueOf("reflection test")
	mtV.Method(1).Call(params)
	fmt.Println("After:", mtV.Method(2).Call(nil)[0])
}

type MyStruct struct {
	name string
}

func (this *MyStruct) GetName(str string) string {
	this.name = str
	return this.name
}

func Test_4444444444444(t *testing.T) {

	// 备注: reflect.Indirect -> 如果是指针则返回 Elem()
	// 首先，reflect包有两个数据类型我们必须知道，一个是Type，一个是Value。
	// Type就是定义的类型的一个数据类型，Value是值的类型

	// 对象
	s := "this is string"

	// 获取对象类型 (string)
	fmt.Println(reflect.TypeOf(s))

	// 获取对象值 (this is string)
	fmt.Println(reflect.ValueOf(s))

	// 对象
	var x float64 = 3.4

	// 获取对象值 (<float64 Value>)
	fmt.Println(reflect.ValueOf(x))

	// 对象
	a := &MyStruct{name: "nljb"}

	// 返回对象的方法的数量 (1)
	fmt.Println(reflect.TypeOf(a).NumMethod())

	// 遍历对象中的方法
	for m := 0; m < reflect.TypeOf(a).NumMethod(); m++ {
		method := reflect.TypeOf(a).Method(m)
		fmt.Println(method.Type)         // func(*main.MyStruct) string
		fmt.Println(method.Name)         // GetName
		fmt.Println(method.Type.NumIn()) // 参数个数
		fmt.Println(method.Type.In(1))   // 参数类型
	}

	// 获取对象值 (<*main.MyStruct Value>)
	fmt.Println(reflect.ValueOf(a))

	// 获取对象名称
	fmt.Println(reflect.Indirect(reflect.ValueOf(a)).Type().Name())

	// 参数
	i := "Hello"
	v := make([]reflect.Value, 0)
	v = append(v, reflect.ValueOf(i))

	// 通过对象值中的方法名称调用方法 ([nljb]) (返回数组因为Go支持多值返回)
	fmt.Println(reflect.ValueOf(a).MethodByName("GetName").Call(v))

	// 通过对值中的子对象名称获取值 (nljb)
	fmt.Println(reflect.Indirect(reflect.ValueOf(a)).FieldByName("name"))

	// 是否可以改变这个值 (false)
	fmt.Println(reflect.Indirect(reflect.ValueOf(a)).FieldByName("name").CanSet())

	// 是否可以改变这个值 (true)
	fmt.Println(reflect.Indirect(reflect.ValueOf(&(a.name))).CanSet())

	// 不可改变 (false)
	fmt.Println(reflect.Indirect(reflect.ValueOf(s)).CanSet())

	// 可以改变
	// reflect.Indirect(reflect.ValueOf(&s)).SetString("jbnl")
	fmt.Println(reflect.Indirect(reflect.ValueOf(&s)).CanSet())

}
