package server

import (
	"github.com/lunny/log"
	"learning/rpc/rpc_codec/msgpk"
	"net"
	"net/rpc"
)

//声明接口类
type EchoService struct{}

//定义方法Echo
func (service *EchoService) Echo(arg string, result *string) error {
	*result = arg
	return nil
}

//服务端启动逻辑
func RegisterAndServeOnTcp() {
	err := rpc.Register(&EchoService{}) //注册并不是注册方法，而是注册EchoService的一个实例
	if err != nil {
		log.Fatal("error registering", err)
		return
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		log.Fatal("error resolving tcp", err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error accepting", err)
		} else {
			//这里先通过NewServerCodec获得一个实例，然后调用rpc.ServeCodec来启动服务
			rpc.ServeCodec(msgpk.NewServerCodec(conn))
		}
	}
}
