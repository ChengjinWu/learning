package msgpk

import (
	"bytes"
	"encoding/binary"
	"io"
	"net/rpc"
)

type MsgpackReq struct {
	rpc.Request             //head
	Arg         interface{} //body
}

type MsgpackResp struct {
	rpc.Response             //head
	Reply        interface{} //body
}

func readData(conn io.ReadWriteCloser) (data []byte, returnError error) {
	const HeadSize = 4 //设定长度部分占4个字节
	headBuf := bytes.NewBuffer(make([]byte, 0, HeadSize))
	headData := make([]byte, HeadSize)
	for {
		readSize, err := conn.Read(headData)
		if err != nil {
			returnError = err
			return
		}
		headBuf.Write(headData[0:readSize])
		if headBuf.Len() == HeadSize {
			break
		} else {
			headData = make([]byte, HeadSize-readSize)
		}
	}
	bodyLen := int(binary.BigEndian.Uint32(headBuf.Bytes()))
	bodyBuf := bytes.NewBuffer(make([]byte, 0, bodyLen))
	bodyData := make([]byte, bodyLen)
	for {
		readSize, err := conn.Read(bodyData)
		if err != nil {
			returnError = err
			return
		}
		bodyBuf.Write(bodyData[0:readSize])
		if bodyBuf.Len() == bodyLen {
			break
		} else {
			bodyData = make([]byte, bodyLen-readSize)
		}
	}
	data = bodyBuf.Bytes()
	returnError = nil
	return
}
