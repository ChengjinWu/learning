package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"log"
)

func main() {
	js, err := simplejson.NewJson([]byte(`{
        "test": {
            "array": [1, "2", 3],
            "arraywithsubs": [
                {"subkeyone": 1},
                {"subkeytwo": 2, "subkeythree": 3}
            ],
            "bignum": 8000000000
        }
    }`))

	if err != nil || js == nil {
		log.Fatal("something wrong when call NewFromReader")
	}
	fmt.Println(js) //&{map[test:map[array:[1 2 3] arraywithsubs:[map[subkeyone:1] map[subkeytwo:2 subkeythree:3]] bignum:8000000000]]}

	arr, err := js.Get("test").Get("array").Array()
	if err != nil || arr == nil {
		log.Fatal("something wrong when call Get and Array")
	}
	fmt.Println(arr) //[1 2 3]

	//使用下面的Must类方法就不用判断而err了
	fmt.Println(js.Get("test").Get("array").MustArray()) //[1 2 3]

	fmt.Println(js.Get("test").Get("arraywithsubs").GetIndex(4).MustMap()) //map[subkeyone:1]

	fmt.Println(js.Get("test").Get("bignum").MustInt64()) //8000000000
}
