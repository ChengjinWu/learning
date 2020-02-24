package helper

import (
	"bytes"
	l4g "code.google.com/p/log4go"
	"encoding/gob"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
	"tripod/define"
	"tripod/devkit"
)

var Log l4g.Logger
var RpcProtocol string = "tcp"
var ListenAddr string = "0.0.0.0:"
var CacheSize int = 100000
var ProxyConfig []string
var RealDumpPostfix = ".dat"
var TmpPostfix = ".tmp"
var DumpInterval = 5000
var indexTimeParam = "id"
var PersistMaxLength = 100000

const (
	INDEX_TIME_PATH   = "/indextime"
	UPDATE_TOTAL_PATH = "/updatetotal"
	STATUS_PATH       = "/status"
)

var ErrItemUnavailable = errors.New("dproxy: Item require is not available")
var ErrServerUnavailable = errors.New("dproxy: server is  unavailable")
var ErrItemMissing = errors.New("dproxy: item missing")
var ErrVersionNotMatch = errors.New("dproxy: version not matched!")
var ErrRetryNoSuccess = errors.New("dproxy: retried many times, no success")
var ErrInvalidMode = errors.New("invalid mode!")
var ErrTriggleTotalUpdate = errors.New("triggle total update")
var ErrInvalidArg = errors.New("invalid arguement!")
var ErrUnsupportedEncType = errors.New("unsupported encType")

var FetchupdatePorts = []int{5001, 5002, 5003, 5004}

var MTU = 1472

const TimeLayout = "2006-01-02 15:04:05"
const (
	B  = 1 << (iota * 10)
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
	PB = 1 << (iota * 10)
	EB = 1 << (iota * 10)
)

var nArr = []float64{B, KB, MB, GB, TB, PB, EB}
var sArr = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

type WorkMode int

const (
	STARTMODE     WorkMode = 0
	UPDATEMODE    WorkMode = 1
	EMERGENCYMODE WorkMode = 2
)

type RpcTotalRequest struct {
	VersionId string
	SeqId     int
	Hash      string // md5 hash value of total buf
	ClientId  string
}

func (self *RpcTotalRequest) String() string {
	return fmt.Sprintf("RpcTotalRequest[ClientId:%v, VersionId:%v, SeqId:%v]", self.ClientId, self.VersionId, self.SeqId)
}

type RpcRealRequest struct {
	VersionId string
	SeqId     int
	ClientId  string
}

func (self *RpcRealRequest) String() string {
	return fmt.Sprintf("RpcRealRequest[ClientId:%v, VersionId:%v, SeqId:%v]", self.ClientId, self.VersionId, self.SeqId)
}

type TotalInfo struct {
	SeqId int
	EncType
	Buff []byte
}

func (self *TotalInfo) String() string {
	return fmt.Sprintf("TotalInfo[SeqId:%d, EncType:%v, Buff.len:%d]", self.SeqId, self.EncType, len(self.Buff))
}

type RpcTotalReply struct {
	Status
	TotalInfo
}

func (self *RpcTotalReply) String() string {
	return fmt.Sprintf("RpcTotalReply[Status:%v, SeqId:%d, EncType:%v, Buff.len:%d]", self.Status, self.SeqId, self.EncType, len(self.Buff))
}

type RpcRealReply struct {
	Items    []*DataItem
	ErrStart int
	Status
}

func (self *RpcRealReply) String() string {
	return fmt.Sprintf("RpcRealReply[Status:%v, ErrStart:%d, Items.len:%d]", self.Status, self.ErrStart, len(self.Items))
}

type DataItem struct {
	SeqId int
	Data  []byte
	Time  time.Time
}

type DataItems []*DataItem

func (p DataItems) Len() int           { return len(p) }
func (p DataItems) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p DataItems) Less(i, j int) bool { return p[i].SeqId < p[j].SeqId }

type DataGrabHandler interface {
	TotalGetFunc(mode WorkMode, seqId int) (*TotalInfo, error)
	RealGetFunc(seqId int) ([]*DataItem, error)
}

type DataConsumer interface {
	UpdateTotal(totalInfo *TotalInfo) error
	UpdateReal(realItems []*DataItem) error
}

type indexInfo struct {
	Id int
	T  time.Time
}

func init() {
	//regist for once
	gob.Register(&DataItem{})
	gob.Register(&[]*DataItem{})
}

func GetDumpFile(DataPath, DataPrefix string, id int) string {
	if id > 0 {
		return DataPrefix + strconv.Itoa(id)
	}
	for _, file := range devkit.GetFilesInDir(DataPath) {
		if strings.HasPrefix(file, DataPrefix) {
			return file
		}
	}
	return ""
}

func GetLastTotalId(DataPath, DataPrefix string) int {
	dumpfile := GetDumpFile(DataPath, DataPrefix, -1)
	if dumpfile == "" {
		return 0
	}
	seqId := devkit.Atoi(strings.TrimPrefix(dumpfile, DataPrefix))
	return seqId
}

func WriteFileWithRetry(dumpFile string, buf []byte, tryCnt int) (e error) {
	for i := 0; i < tryCnt; i++ {
		if tmpfile, err := os.OpenFile(dumpFile, os.O_RDWR|os.O_CREATE, 0644); err == nil {
			_, err = tmpfile.Write(buf)
			tmpfile.Close()
			if err == nil {
				return nil
			} else {
				Log.Error("WriteFileWithRetry try %d times, file=%s, err=%v", i, dumpFile, err)
				e = err
			}
		} else {
			Log.Error("WriteFileWithRetry try %d times, file=%s, err=%v", i, dumpFile, err)
			e = err
		}
	}
	return fmt.Errorf("%v, %v", e, ErrRetryNoSuccess)
}

//get min and max id from real dump file
func GetSequenceFromFileName(name string) (int, int) {
	if strings.HasSuffix(name, RealDumpPostfix) {
		tmpFile := strings.TrimSuffix(name, RealDumpPostfix)
		if l := strings.Split(tmpFile, "_"); len(l) == 3 {
			return devkit.Atoi(l[0]), devkit.Atoi(l[1])
		}
	}
	return 0, 0
}

func LoadRealDumpFiles(dir string, minId int) []*DataItem {
	Log.Info("LoadRealDumpFiles minId=%d dir=%s", minId, dir)
	ret := make(DataItems, 0)
	for _, file := range devkit.GetFilesInDir(dir) {
		if !strings.HasSuffix(file, RealDumpPostfix) {
			continue
		}
		tmpfile := path.Base(file)
		min, max := GetSequenceFromFileName(tmpfile)
		if !(min > 0 && max > 0) || max <= minId {
			continue
		}
		tbuf, err := ioutil.ReadFile(file)
		if err != nil {
			Log.Error("LoadRealDumpFiles read file failed! file=%s, err=%v", file, err)
			continue
		}
		if items, err := loadDumpFile(tbuf); err != nil {
			Log.Error("LoadRealDumpFiles gob decode failed! file=%s, err=%v", file, err)
		} else {
			if len(items) > 0 {
				cnt := 0
				minId, maxId := 0, 0
				for _, item := range items {
					if item.SeqId > minId {
						ret = append(ret, item)
						cnt++
						if item.SeqId < minId {
							minId = item.SeqId
						}
						if item.SeqId > maxId {
							maxId = item.SeqId
						}
					}
				}
				Log.Info("LoadRealDumpFiles %s got %d range:[%d, %d] ", tmpfile, cnt, min, max)
			}
		}
	}

	// sort the items
	sort.Sort(ret)
	return ret
}

func loadDumpFile(tbuf []byte) ([]*DataItem, error) {
	fbuf := bytes.NewReader(tbuf)
	dec := gob.NewDecoder(fbuf)
	items := []*DataItem{}
	if err := dec.Decode(&items); err != nil {
		return nil, err
	}
	return items, nil
}

func getReadableSize(input int64) string {
	n := float64(input)
	if n < nArr[0] {
		return fmt.Sprintf("%dB", input)
	}

	for i, pn := range nArr {
		if n < pn {
			n /= float64(nArr[i-1])
			return fmt.Sprintf("%3.1f%s", n, sArr[i-1])
		}
	}
	return fmt.Sprintf("%.1f%s", float64(input)/nArr[len(nArr)-1], sArr[len(sArr)-1])
}

func GetFetchupdatePort(versionId string) int {
	length := len(FetchupdatePorts)

	//eg: 1.0.71
	if tl := strings.Split(versionId, define.StrDot); len(tl) == 3 {
		if version, err := strconv.Atoi(tl[2]); err != nil {
			Log.Error("parse versionId error! %v %v", versionId, err)
			Log.Close()
			panic("parse versionId failed")
		} else {
			return FetchupdatePorts[version%length]
		}
	} else {
		Log.Error("versionId is not right! %v", versionId)
		Log.Close()
		panic("versionId is not right!")
	}
}

func TcpListen(addr string) (net.Listener, error) {
	if l, err := net.Listen(RpcProtocol, addr); err != nil {
		return nil, err
	} else {
		return l, nil
	}
}
