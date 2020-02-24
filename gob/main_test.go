package gob

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

type Person struct {
	name string
	password string
}

func TestGob(t *testing.T)  {
	p := Person{
		name:"你好",
		password:"haha",
	}
	buf:=bytes.Buffer{}
	enc:=gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(buf.String())
}
