package helper

import (
	"net"
	"net/rpc"
	"time"
)

type RpcServer struct {
	Server *rpc.Server
}

func NewRpcServer() *RpcServer {
	rpcServer := &RpcServer{
		Server: rpc.NewServer(),
	}
	rpcServer.Server.Register(rpcServer)
	return rpcServer
}

func StartRpcServer(server *RpcServer, listener net.Listener, syncCnt uint) {

	conn, err := listener.Accept()
	if err != nil {
		Log.Error("rpc.Serve: accept:", err.Error()) // TODO(r): exit?
	}
	go func() {
		tcpConn, _ := (conn).(*net.TCPConn)
		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(20 * time.Second)
		var sh chan int
		if syncCnt > 0 {
			sh = make(chan int, syncCnt)
		}
		srv := NewUserRpcServerCodec(conn, 0, sh, 0)
		server.Server.ServeCodec(srv)
	}()
}

func (self *RpcServer) RpcGetTotal(req *RpcTotalRequest, resp *RpcTotalReply) error {
	if req == nil {
		return ErrInvalidArg
	}

	tmp := &TotalInfo{}
	totalStatus := tmp.SeqId

	Log.Info("RpcGetTotal totalStatus=%d req=%s", totalStatus, req)
	defer Log.Info("RpcGetTotal resp=%s", resp)

	if self.DCommon.VersionId != req.VersionId {
		resp.Status = STATUS_VERSIONERROR
		return nil
	}
	if len(tmp.Buff) <= 0 {
		resp.Status = STATUS_UNKNOWERROR
		return nil
	}

	if totalStatus < req.SeqId {
		resp.Status = STATUS_NOTREADY
		return nil
	} else if totalStatus == req.SeqId {
		resp.Status = STATUS_UPDATED
		resp.SeqId = totalStatus
		return nil
	} else {
		resp.Status = STATUS_OK
		resp.SeqId = tmp.SeqId
		resp.EncType = tmp.EncType
		resp.Buff = tmp.Buff
	}
	return nil
}
