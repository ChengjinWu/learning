package main

import (
	"learning/rpc/gRpc/pb"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func TestMainClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:12345", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewSimpleClient(conn)

	//for {
	line := strconv.Itoa(rand.Int())
	if err != nil {
		log.Fatal(err)
	}
	reply, err := c.GetLine(context.Background(), &pb.SimpleRequest{Data: string(line)})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Reply: %v, Data: %v", reply, reply.Data)
	time.Sleep(time.Second)
	//}
}
