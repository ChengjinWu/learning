package _map

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
)

func TestBlock(t *testing.T) {
	m := sync.Map{}
	m.Store("123", "234343")
	m.Load("123")
	m.Load("1232142")
	m.Load("123")
	m.Load("sdfsfsd")
	jb, _ := json.Marshal(m)
	fmt.Println(string(jb))
}

func TestConcurrent(t *testing.T) {

}
