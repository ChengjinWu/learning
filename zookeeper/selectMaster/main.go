package main

import (
	"github.com/gin-gonic/gin/json"
	"log"
	"net"
	"time"
	"tripod/define"
	"tripod/devkit"
	"tripod/zkutil"
)

var RegNet *net.IPNet

func masterSlaveHA() {
	serviceName := devkit.GetServiceName(define.FetchupdateServiceName)
	zkTimeout := 10 * time.Second
	port := 10000
	zkHosts := "127.0.0.1:2181"
	zkInstance := zkutil.GetZkInstance(zkHosts, zkTimeout, RegNet)
	notifyCh, err := zkInstance.MasterSelection(define.GroupServing, serviceName, port, true, false)
	if err != nil {
		log.Panic(err)
	}

	for {
		if notify, ok := <-notifyCh; ok && notify != nil {
			jsonBytes, _ := json.Marshal(notify)
			log.Println(string(jsonBytes))
			if notify.IsMaster {
				log.Println("IsMaster")
			} else {

				log.Println("IsSlave")
			}

		}
	}
}

func main() {
	masterSlaveHA()
}
