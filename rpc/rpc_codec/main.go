package main

import (
	"fmt"
	"github.com/lunny/log"
	"learning/rpc/rpc_codec/client"
	"learning/rpc/rpc_codec/server"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

//main函数
func main() {
	go server.RegisterAndServeOnTcp() //先启动服务端
	time.Sleep(1e9)
	wg := new(sync.WaitGroup) //waitGroup用于阻塞主线程防止提前退出
	callTimes := 10
	wg.Add(callTimes)
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
			wg.Done()
		}()
	}
	wg.Wait()
}
