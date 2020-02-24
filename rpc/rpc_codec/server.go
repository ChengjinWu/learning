package main

import "learning/rpc/rpc_codec/server"

func main() {
	server.RegisterAndServeOnTcp() //先启动服务端
}
