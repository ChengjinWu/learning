package test

import (
	"learning/reflect/produce"
	"testing"
)

type OtherInfo struct {
	Age int
}

func Test_json(t *testing.T) {
	data := `{"id":524042,"name":"酷旅-mob-otv-2","male":true,"other":null}`
	stu := struct {
		Id    int
		Name  string
		Male  bool
		Other *OtherInfo
	}{}
	gojson.Unmarshal([]byte(data), &stu)
}
