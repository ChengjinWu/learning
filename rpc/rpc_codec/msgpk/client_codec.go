package msgpk

import (
	"encoding/binary"
	"errors"
	"github.com/vmihailenco/msgpack"
	"io"
	"net"
	"net/rpc"
	"reflect"
	"strings"
)

func NewClientCodec(conn net.Conn) *MessagePackClientCodec {
	return &MessagePackClientCodec{conn, MsgpackResp{}, false}
}

type MessagePackClientCodec struct {
	rwc    io.ReadWriteCloser
	resp   MsgpackResp //用于缓存解析到的请求
	closed bool
}

func (c *MessagePackClientCodec) WriteRequest(r *rpc.Request, arg interface{}) error {
	//先判断codec是否已经关闭，如果是则直接返回
	if c.closed {
		return nil
	}
	//将r和arg组装成一个MsgpackReq并序列化
	request := &MsgpackReq{*r, arg}
	reqData, err := msgpack.Marshal(request)
	if err != nil {
		panic(err)
		return err
	}
	//先发送数据长度
	head := make([]byte, 4)
	binary.BigEndian.PutUint32(head, uint32(len(reqData)))
	_, err = c.rwc.Write(head)
	//再将序列化产生的数据发送出去
	_, err = c.rwc.Write(reqData)
	return err
}

func (c *MessagePackClientCodec) ReadResponseHeader(r *rpc.Response) error {
	//先判断codec是否已经关闭，如果是则直接返回
	if c.closed {
		return nil
	}
	//读取数据
	data, err := readData(c.rwc)
	if err != nil {
		//client一旦初始化就会开始轮询数据，所以要处理连接close的情况
		if strings.Contains(err.Error(), "use of closed network connection") {
			return nil
		}
		panic(err) //简单起见，出现异常直接panic
	}

	//将读取到的数据反序列化成一个MsgpackResp
	var response MsgpackResp
	err = msgpack.Unmarshal(data, &response)

	if err != nil {
		panic(err) //简单起见，出现异常直接panic
	}

	//根据读取到的数据设置request的各个属性
	r.ServiceMethod = response.ServiceMethod
	r.Seq = response.Seq
	//同时将读取到的数据缓存起来
	c.resp = response

	return nil
}

func (c *MessagePackClientCodec) ReadResponseBody(reply interface{}) error {
	//这里直接用缓存的数据返回即可

	if "" != c.resp.Error { //如果返回的是异常
		return errors.New(c.resp.Error)
	}
	if reply != nil {
		//正常返回，通过反射将结果设置到reply变量，因为reply一定是指针类型，所以不必检查CanSet
		reflect.ValueOf(reply).Elem().Set(reflect.ValueOf(c.resp.Reply))
	}
	return nil
}

func (c *MessagePackClientCodec) Close() error {
	c.closed = true //关闭时将closed设置为true
	if c.rwc != nil {
		return c.rwc.Close()
	}
	return nil
}
