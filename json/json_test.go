package json

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"testing"
)

type biddingConfig struct {
	AdvWeightMap           map[int]float32
	AdvEnterpriseWeight    float32
	AdvNonEnterpriseWeight float32
	AdPlaceMinBidPriceMap  map[int32]map[int]float32
}

func TestJson222(t *testing.T) {
	data := "[{\"c\":\"我们提倡绿色直播，严禁涉政、涉恐、涉群体性事件、涉黄等直播内容。\",\"clr\":\"#FEE494\",\"typ\":1}]"
	object := []map[string]interface{}{}
	err := json.Unmarshal([]byte(data), &object)
	if err != nil {
		log.Error(err)
	}
	object22 := map[string]interface{}{
		"para": object,
	}
	jb, _ := json.Marshal(object22)
	log.Info(string(jb))
	fmt.Println(string(jb))

}
func TestJson(t *testing.T) {
	BiddingConfig := biddingConfig{}
	//BiddingConfig := biddingConfig{
	//	AdvWeightMap: map[int]float32{
	//		1: 123.23,
	//		2: 23.23,
	//		3: 656.23,
	//		4: 764.23,
	//		5: 23432.23,
	//	},
	//	AdvEnterpriseWeight:    1.5,
	//	AdvNonEnterpriseWeight: 1,
	//	AdPlaceMinBidPriceMap: map[int32]map[int]float32{
	//		0: map[int]float32{
	//			1: 0.2,
	//			2: 0.3,
	//		},
	//		3: map[int]float32{
	//			1: 0.2,
	//			2: 0.3,
	//		},
	//		5: map[int]float32{
	//			1: 0.2,
	//			2: 0.3,
	//		},
	//	},
	//}
	data := []byte(`{"AdvEnterpriseWeight":1.5,"AdvNonEnterpriseWeight":1,"AdPlaceMinBidPriceMap":{"0":{"1":0.2,"2":0.3},"3":{"1":0.2,"2":0.3},"5":{"1":0.2,"2":0.3}}}`)
	if len(data) > 0 {
		err := json.Unmarshal(data, &BiddingConfig)
		if err != nil {
			t.Errorf("ad_bidding_config data format error:%s:%s", err, data)
		}
	}
	jsonBytes, err := json.Marshal(BiddingConfig)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("get ad_bidding_config data:%s\n", jsonBytes)

}
