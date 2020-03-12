package main

import (
	"github.com/lunny/log"
	"math/rand"
	"net/rpc/jsonrpc"
	"strconv"
	"testing"
	"time"
)

func TestMainClient(t *testing.T) {
	client, err := jsonrpc.Dial("tcp", "localhost:12345") // 只改动这一行
	if err != nil {
		log.Fatal(err)
	}

	//for {
	line := strconv.Itoa(rand.Int())
	if err != nil {
		log.Fatal(err)
	}
	var reply Reply
	err = client.Call("Listener.GetLine", []byte(line), &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Reply: %v, Data: %v", reply, reply.Data)
	time.Sleep(time.Second)
	//}
}
