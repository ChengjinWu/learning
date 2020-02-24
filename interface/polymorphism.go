package main

import (
	"fmt"
)

type base interface { //类似基类定义
	virtualfunc() int //类似虚函数抽象定义
}

type der1 int //类似派生类1定义

func (der1) virtualfunc() int { //类似派生类1虚函数具体实现
	fmt.Printf("I'm der1\n")
	return 1
}

type der2 struct { //类似派生类2定义
	//nothing
}

func (der2) virtualfunc() int { //类似派生类2虚函数具体实现
	fmt.Printf("I'm der2\n")
	return 2
}

func somefunc(b base) { //作为某个函数的形参
	b.virtualfunc()
}
