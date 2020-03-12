package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {

	client := &http.Client{}
	data := `{
  "appname":"chatright",
  "time":"2019-10-11 00:01:02.903",
  "caller":"",
  "topic":"room_call_new", 
  "service_name":"room.paly.call",  
  "atom": {
    "cc": "TG0001",
    "conn": "WiFi",
    "cv": "IK7.0.50_Iphone",
    "devi": "1f54b5facba591ffdef8d8d41e747e385612e72d",
    "devicetoken": "",
    "idfa": "3A1ADC93-7030-4D9C-9B6F-DC65B76D6261",
    "idfv": "09D6ACF6-907B-4697-8BCD-0F5064EBC2EA",
    "imei": "",
    "imsi": "",
    "lc": "0000000000000134",
    "live_owner": 12036000,
    "logid": "",
    "mtid": "50ed7b1a38fa8a3aa78a59a64182a86d",
    "mtxid": "6696cd6adb5",
    "ndid": "",
    "osversion": "ios_11.100000",
    "pass-through": {},
    "proto": "13",
    "pub_stat": 0,
    "sid": "20V5xd4xblsPWln2JzeHmc5OKwi1GhYBPLcCVnzc9nxIRmotcgbRAi3i3",
    "smid": "D2yz6M7vou0qzHdIg2pJylLVEcETp1dIKGUG7BadHTvZ4Xc9",
    "ua": "iPhone9_1",
    "uid": "12036000",
    "userip": "61.135.45.226"
  }, 
  "info":{
    "liveid": "1552722680851560",
    "room_id": 100127,
    "section": 0,
    "song_info": {
      "author": "张翰|郑爽",
      "drc_path": "http://m4a.inke.cn/MTU1Mjg5NDY4ODMyMiMxODkjZHJj.drc",
      "duration": 242,
      "id": "1580",
      "name": "极限爱恋",
      "type": 0,
      "zip_path": "http://m4a.inke.cn/MTU1Mjg5NDY5NDY1MCM2NiN6aXA=.zip"
    },
    "type": 1,
    "uid": 12036000,
    "userid": 12036000
  }
}`
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	defer w.Close()
	w.Write([]byte(data))
	w.Flush()
	fmt.Println(buf.Bytes())
	body := bytes.NewReader(buf.Bytes())

	req, err := http.NewRequest("POST", "http://maidian.fengyuhn.cn/log/upload?cv=DL1.1.00_Android", body)
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	if resp != nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(string(body))
	}

}
