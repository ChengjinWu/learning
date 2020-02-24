package client

import (
	"learning/rpc/rpc_codec/msgpk"
	"net"
	"net/rpc"
)

//客户端调用逻辑
func Echo(arg string) (result string, err error) {
	var client *rpc.Client
	conn, err := net.Dial("tcp", ":1234")
	client = rpc.NewClientWithCodec(msgpk.NewClientCodec(conn))

	defer client.Close()

	if err != nil {
		return "", err
	}
	err = client.Call("EchoService.Echo", arg, &result) //通过类型加方法名指定要调用的方法
	if err != nil {
		return "", err
	}
	return result, err
}
