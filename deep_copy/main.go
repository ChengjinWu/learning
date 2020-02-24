package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
)

type User struct {
	data  map[int]string
	array []int
}

type Employee struct {
	Name      string
	Age       int32
	DoubleAge int32
	EmployeId int64
	SuperRule string
}

func (employee *Employee) Role(role string) {
	employee.SuperRule = "Super " + role
}

func main() {
	var (
		user = User{
			data: map[int]string{
				11:   "23423",
				3432: "23421423",
			},
			array: []int{
				11, 232, 321312, 432423,
			},
		}
		user22 = User{data: map[int]string{}}
	)
	fmt.Printf("%p\n", user.data)
	fmt.Printf("%p\n", user22.data)
	fmt.Printf("%p\n", user.array)
	fmt.Printf("%p\n", user22.array)

	copier.Copy(&user22, &user)

	jb, _ := json.Marshal(user22.data)
	fmt.Println(string(jb))
	fmt.Printf("%p\n", user.data)
	fmt.Printf("%p\n", user22.data)
	fmt.Printf("%p\n", user.array)
	fmt.Printf("%p\n", user22.array)
	fmt.Printf("%p\n", &user)
	fmt.Printf("%p\n", &user22)
}
