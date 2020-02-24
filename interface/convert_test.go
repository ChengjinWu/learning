package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	data := `{
  "atom": 12323,
  "c": "aa",
  "liveid": "1571050205720750",
  "server_ip": "10.128.0.186",
  "userid": 100240
}`
	req := make(map[string]interface{})
	json.Unmarshal([]byte(data), &req)
	atom, ok := req["atom"]
	fmt.Println(atom, ok)
	fmt.Println(atom.(float64), ok)

}
