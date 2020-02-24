package main

import (
	"fmt"
	"github.com/lunny/log"
	"learning/rpc/rpc_codec/client"
	"math/rand"
	"strconv"
)

func main() {
	callTimes := 10

	for i := 0; i < callTimes; i++ {
		go func() {
			//使用hello world加一个随机数作为参数
			argString := "hello world " + strconv.Itoa(rand.Int())
			resultString, err := client.Echo(argString)
			if err != nil {
				log.Fatal("error calling:", err)
			}
			if resultString != argString {
				fmt.Println("error")
			} else {
				fmt.Printf("echo:%s\n", resultString)
			}
		}()
	}
}
