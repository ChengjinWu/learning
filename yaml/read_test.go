package test

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
	"os"
	"testing"
	"tripod/zconf"
	"util/conf"
)

var ServiceConfig struct {
	conf.BaseConf        `yaml:",inline"`
	BidserverRedis       []string
	Port                 int
	Timeout              int
	LogConfig            string
	InstantMasterCluster []string
	ItemRedisCluster     []string
	CbRedisCluster       []string
	OnlyTransfor         bool
	//拉模式， 上游节点地址
	PullAddr []string
	//推模式,  下有节点地址
	PushAddr  []string
	MemSize   int // 内存事件最大个数
	BatchSize int // 每次批量获取的事件个数
	Mode      int // 获取事件的模式，均衡、随机等
}

func init() {
	fmt.Println(os.Getwd())
	configPath := "conf/bj/sync.yaml"
	zconf.ParseYaml(configPath, &ServiceConfig)
	jsonBytes, _ := json.Marshal(ServiceConfig)
	fmt.Println(string(jsonBytes))
}

func Test_111(t *testing.T) {
	fmt.Println(111)
}
