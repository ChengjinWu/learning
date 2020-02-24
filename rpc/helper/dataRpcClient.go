package helper

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
	"tripod/devkit"
)

type GetServerAddr func() string

type RpcClient struct {
	getServerAddr GetServerAddr
	versionId     string
}

var dialer = &net.Dialer{KeepAlive: 20 * time.Second}

func NewRpcClient(getServerAddr GetServerAddr, version string) *RpcClient {
	if getServerAddr == nil {
		Log.Error("NewRpcClient getServerAddr func is nil!")
		Log.Close()
		panic("NewRpcClient getServerAddr func is nil!")
	}

	return &RpcClient{
		getServerAddr: getServerAddr,
		versionId:     version,
	}
}

func (self *RpcClient) GetNewClient() (*rpc.Client, string, error) {
	serverAddr := self.getServerAddr()
	conn, err := dialer.Dial(RpcProtocol, serverAddr)
	if err != nil {
		return nil, "", err
	}
	return rpc.NewClient(conn), serverAddr, nil
}

//fetch total index from rpc server
func (self *RpcClient) TotalGetFunc(mode WorkMode, dumpId int) (*TotalInfo, error) {

	client, serverAddr, err := self.GetNewClient()
	if client == nil || err != nil {
		Log.Error("TotalGetFunc GetNewClient failed! %s err=%v", serverAddr, err)
		return nil, err
	}

	if total, err := self.requestTotal(client, dumpId); err == nil {
		t := &TotalInfo{
			SeqId:   total.SeqId,
			EncType: total.EncType,
			Buff:    total.Buff,
		}
		return t, nil
	} else {
		Log.Error("requestTotal failed! server=%s err=%v", serverAddr, err)
		return nil, err
	}
}

func (self *RpcClient) requestTotal(client *rpc.Client, dumpId int) (*RpcTotalReply, error) {
	defer client.Close()
	req := &RpcTotalRequest{
		VersionId: self.versionId,
		SeqId:     dumpId,
		ClientId:  devkit.GetHostName(),
	}
	Log.Info("requestTotal req=%s", req)

	resp := &RpcTotalReply{Status: STATUS_UNKNOWERROR}
	if err := client.Call("RpcServer.RpcGetTotal", req, resp); err != nil {
		Log.Error("rpc.Call failed! err=%v", err)
		return nil, err
	} else {
		Log.Info("requestTotal resp=%s", resp)
	}

	var err error
	switch resp.Status {
	case STATUS_UPDATED:
		return resp, nil
	case STATUS_VERSIONERROR:
		err = fmt.Errorf("resp.Status=%v try another...", resp.Status)
	case STATUS_NOTREADY:
		Log.Warn("resp.Status=%v, retry later, seqId=%d", resp.Status, dumpId)
		resp.SeqId = dumpId
		return resp, nil
	case STATUS_UNKNOWERROR:
		err = fmt.Errorf("resp.Status=%v retry later, seqId=%d", resp.Status, dumpId)
		time.Sleep(time.Second)
	case STATUS_OK:
		return resp, nil
	default:
		err = fmt.Errorf("unknown error,status=%v retry later, seqId=%d", resp.Status, dumpId)
	}
	return nil, fmt.Errorf("%s %s", err, ErrRetryNoSuccess)
}
