package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"os"
	"time"
)

var zookeeperServers = []string{"localhost:2181"}
var namespacePath = "/namespacePath"
var childPath = "subNamespacePath"

func runServer(id chan int, childId string) {
	conn, connChan, err := zk.Connect(zookeeperServers, 3*time.Second)
	if err != nil {
		fmt.Println("get connection from zookeeper error")
		return
	}
	connEvent := <-connChan
	if connEvent.State == zk.StateConnected {
		fmt.Println("connect to zookeeper server success!")
	} else {
		fmt.Println("connect to zookeeper err", connEvent.State)
	}

	startSelectLeader(conn, childPath, childId)
	fmt.Println(namespacePath + "/" + childPath)
	children, state, childChan, err := conn.ExistsW(namespacePath + "/" + childPath)
	if err != nil {
		fmt.Println("watch children error, ", err)
	}
	fmt.Println("watch children result, ", children, state)
	isKilled := false
	for {
		select {
		case childEvent, ok := <-childChan:
			fmt.Println("####################################", ok)
			fmt.Println(childEvent.Type, childEvent.Path, childEvent.State, childEvent.Server)
			switch childEvent.Type {
			case zk.EventNodeDeleted:
				fmt.Println("znode delete event ", childEvent)
				fmt.Println("start select leader")
				startSelectLeader(conn, childPath, childId)
			case zk.EventNodeCreated:
				fmt.Println("the zookeeper state is ", state)
			case zk.EventNotWatching:
				startSelectLeader(conn, childPath, childId)
			default:
				time.Sleep(time.Second)
				fmt.Println(childEvent.Type)
				startSelectLeader(conn, childPath, childId)
				children, state, childChan, err = conn.ExistsW(namespacePath + "/" + childPath)
				if err != nil {
					fmt.Println("watch children error, ", err)
				}
				fmt.Println("watch children result, ", children, state)
			}
		case killedId := <-id:
			fmt.Println("I am killed", killedId)
			isKilled = true
			break
		}
		if isKilled {
			break
		}
	}

	fmt.Println("please close the connection")
	conn.Close()
}

func startSelectLeader(conn *zk.Conn, childPath string, childId string) {
	path, err := conn.Create(namespacePath, nil, 0, zk.WorldACL(zk.PermAll))
	if err == nil {
		fmt.Println("create root path success", path)
	} else {
		fmt.Printf("create root path failure err:%v\n ", err)
	}

	childpath, err := conn.Create(namespacePath+"/"+childPath, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	fmt.Println(childpath)
	if err == nil {
		fmt.Printf("create childPath success and child path is %v child id is %v\n", childpath, childId)
	} else {
		fmt.Printf("create child path failure err:%v\n ", err)
	}
}

func main() {
	id0 := make(chan int)
	id1 := make(chan int)
	id := os.Args[1]
	go runServer(id0, id)

	time.Sleep(time.Duration(5) * time.Second)
	<-id1
}
