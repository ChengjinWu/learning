package msg

import (
	"testing"
	"time"

	//protobuf编解码库,下面两个库是相互兼容的，可以使用其中任意一个
	"github.com/golang/protobuf/proto"
	//"github.com/gogo/protobuf/proto"
)

func Test_NilElement(t *testing.T) {
	itemlist := new(ItemList)
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			itemlist.Items = append(itemlist.Items, nil)
		} else {
			item := &Item{
				Time:  time.Now().Unix(),
				Value: int64(i),
			}
			itemlist.Items = append(itemlist.Items, item)
		}
	}

	//protobuf编码
	_, err := proto.Marshal(itemlist)
	if err != nil {
		t.Fatal(err)
	}
}
