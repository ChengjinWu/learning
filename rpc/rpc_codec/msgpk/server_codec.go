package msgpk

import (
	"encoding/binary"
	"github.com/vmihailenco/msgpack"
	"io"
	"net"
	"net/rpc"
	"reflect"
	"strings"
)

type MessagePackServerCodec struct {
	rwc    io.ReadWriteCloser //用于读写数据，实际是一个网络连接
	req    MsgpackReq         //用于缓存解析到的请求
	closed bool               //标识codec是否关闭
}

func NewServerCodec(conn net.Conn) *MessagePackServerCodec {
	return &MessagePackServerCodec{conn, MsgpackReq{}, false}
}

func (c *MessagePackServerCodec) WriteResponse(r *rpc.Response, reply interface{}) error {
	//先判断codec是否已经关闭，如果是则直接返回
	if c.closed {
		return nil
	}
	//将r和reply组装成一个MsgpackResp并序列化
	response := &MsgpackResp{*r, reply}

	respData, err := msgpack.Marshal(response)
	if err != nil {
		panic(err)
		return err
	}
	head := make([]byte, 4)
	binary.BigEndian.PutUint32(head, uint32(len(respData)))
	_, err = c.rwc.Write(head)
	//将序列化产生的数据发送出去
	_, err = c.rwc.Write(respData)
	return err
}

func (c *MessagePackServerCodec) ReadRequestHeader(r *rpc.Request) error {
	//先判断codec是否已经关闭，如果是则直接返回
	if c.closed {
		return nil
	}
	//读取数据
	data, err := readData(c.rwc)
	if err != nil {
		//这里不能直接panic，需要处理EOF和reset的情况
		if err == io.EOF {
			return err
		}
		if strings.Contains(err.Error(), "connection reset by peer") {
			return err
		}
		panic(err) //其他异常直接panic
	}
	//将读取到的数据反序列化成一个MsgpackReq
	var request MsgpackReq
	err = msgpack.Unmarshal(data, &request)

	if err != nil {
		panic(err) //简单起见，出现异常直接panic
	}

	//根据读取到的数据设置request的各个属性
	r.ServiceMethod = request.ServiceMethod
	r.Seq = request.Seq
	//同时将解析到的数据缓存起来
	c.req = request

	return nil
}

func (c *MessagePackServerCodec) ReadRequestBody(arg interface{}) error {
	if arg != nil {
		//参数不为nil，通过反射将结果设置到arg变量
		reflect.ValueOf(arg).Elem().Set(reflect.ValueOf(c.req.Arg))
	}
	return nil
}

func (c *MessagePackServerCodec) Close() error {
	c.closed = true
	if c.rwc != nil {
		return c.rwc.Close()
	}
	return nil
}
