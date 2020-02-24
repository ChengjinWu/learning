package helper

import (
	"encoding/gob"
	"io"
	"net/rpc"
	"reflect"
	"strings"
	"time"
)

var (
	totalRespTypeStr = reflect.TypeOf(&RpcTotalReply{}).String()
)

//userRpcServerCodec implements the rpc.ServerCodec interface
type userRpcServerCodec struct {
	tcw     *trafficControlConn
	dec     *gob.Decoder
	enc     *gob.Encoder
	closed  bool
	totalCh chan int
	mtuTime float64
}

func NewUserRpcServerCodec(conn io.ReadWriteCloser, defMtuCostTime float64, ch chan int, length int) *userRpcServerCodec {
	tcw := newTrafficControlConn(conn, length)
	return &userRpcServerCodec{
		tcw:     tcw,
		dec:     gob.NewDecoder(conn),
		enc:     gob.NewEncoder(tcw),
		totalCh: ch,
		mtuTime: defMtuCostTime,
	}
}

func (c *userRpcServerCodec) ReadRequestHeader(r *rpc.Request) error {
	return c.dec.Decode(r)
}

func (c *userRpcServerCodec) ReadRequestBody(body interface{}) error {
	if err := c.dec.Decode(body); err != nil {
		return err
	} else {
		if req, ok := body.(*RpcTotalRequest); ok && req != nil {
			c.tcw.SetClientId(req.ClientId)
			if doTrafficControl(req) {
				// only total update for some clients need limit transmit speed
				c.tcw.SetMtuTime(c.mtuTime)
			}
		}
	}

	return nil
}

func doTrafficControl(req *RpcTotalRequest) bool {
	if strings.HasPrefix(req.ClientId, "bj") ||
		strings.HasPrefix(req.ClientId, "sh") ||
		strings.HasPrefix(req.ClientId, "SC_HOST") {
		return true
	} else {
		return false
	}
}

func (c *userRpcServerCodec) WriteResponse(r *rpc.Response, body interface{}) (err error) {
	defer c.Close()

	if reflect.TypeOf(body).String() == totalRespTypeStr {
		if resp, ok := body.(*RpcTotalReply); ok && resp != nil && resp.Status == STATUS_OK {
			//只对全量请求且需要返回全量索引的情况打印进度
			c.tcw.DoLogProgress()
			if c.totalCh != nil { //只对需要返回索引的全量请求并且需要限制并发数的加以并发限制
				c.totalCh <- 0
				defer func() {
					<-c.totalCh
				}()
			}
		}
	}

	if err = c.enc.Encode(r); err != nil {
		Log.Error("userRpcServerCodec Encode header failed! %v", err)
		return err
	} else if err = c.enc.Encode(body); err != nil {
		Log.Error("userRpcServerCodec Encode response failed! %v", err)
		return err
	}

	return nil
}

func (c *userRpcServerCodec) Close() error {
	if c.closed {
		// Only call c.conn.Close once; otherwise the semantics are undefined.
		return nil
	}
	c.closed = true
	return c.tcw.Close()
}

//connection with traffic control and progress logging
type trafficControlConn struct {
	conn           io.ReadWriteCloser
	limitSpeed     bool
	mtuTime        float64 //time cost per MTU (nanosecond)
	totalLen       int64   //length of the data to be to sent, just for total
	sent           int64   //bytes that has been sent, just for total
	lastMoment     int64
	progressTicker *time.Ticker
	closed         bool
	buf            []byte
	clientId       string
}

func newTrafficControlConn(conn io.ReadWriteCloser, length int) *trafficControlConn {
	u := &trafficControlConn{
		conn:     conn,
		totalLen: int64(length),
	}
	return u
}

func (u *trafficControlConn) SetMtuTime(t float64) {
	u.mtuTime = t
	u.limitSpeed = true
}

func (u *trafficControlConn) SetClientId(id string) {
	u.clientId = id
}

func (u *trafficControlConn) DoLogProgress() {
	u.progressTicker = time.NewTicker(time.Second)
}

func (u *trafficControlConn) Read(buf []byte) (int, error) {
	return u.conn.Read(buf)
}

func (u *trafficControlConn) Write(b []byte) (n int, err error) {
	if u.mtuTime > 0.0 && u.lastMoment == 0 {
		u.lastMoment = time.Now().UnixNano()
	}

	var pch <-chan time.Time
	if u.progressTicker != nil {
		pch = u.progressTicker.C
	}

	sendFunc := func() error {
		if n, err = u.conn.Write(u.buf); err != nil || n < len(u.buf) {
			u.sent += int64(n)
			u.buf = u.buf[n:]
			return err
		}
		u.sent += int64(len(u.buf))
		u.buf = u.buf[0:0]
		return nil
	}

	idx := 0
	for idx < len(b) {
		if idx+MTU < len(b) {
			u.buf = append(u.buf, b[idx:idx+MTU]...)
			idx += MTU
		} else {
			u.buf = append(u.buf, b[idx:]...)
			idx = len(b)
		}

		select {
		case <-pch:
			Log.Info("RpcServer send progress[%s]: %.2f%% %s/%s", u.clientId, 100.00*float64(u.sent)/float64(u.totalLen), getReadableSize(u.sent), getReadableSize(u.totalLen))
		default:
			if err = sendFunc(); err != nil {
				return n, err
			} else if u.mtuTime > 0.0 {
				now := time.Now().UnixNano()
				if left := int64(u.mtuTime) - (now - u.lastMoment); left > 0 {
					time.Sleep(time.Duration(left))
					u.lastMoment = time.Now().UnixNano()
				} else {
					u.lastMoment = now
				}
			}
		}
	}

	return n, err
}

func (u *trafficControlConn) Close() error {
	if u.closed {
		return nil
	}
	u.closed = true
	if u.progressTicker != nil {
		u.progressTicker.Stop()
		if u.sent == u.totalLen {
			Log.Info("RpcServer send progress[%s]: %d%% %s/%s", u.clientId, 100, getReadableSize(u.sent), getReadableSize(u.totalLen))
		} else {
			Log.Warn("RpcServer send failed[%s] at %.2f%% %s/%s", u.clientId, 100.00*float64(u.sent)/float64(u.totalLen), getReadableSize(u.sent), getReadableSize(u.totalLen))
		}
	}

	return u.conn.Close()
}
