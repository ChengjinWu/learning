package _map

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type People struct {
	Name string
	Age  int
}

func TestEqual(t *testing.T) {

	strMap := make(map[string]bool)
	m := make(map[People]bool)
	p1 := People{
		Name: "你好",
		Age:  78,
	}
	m[p1] = true
	strMap["你好_78"] = true
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		p2 := People{
			Name: "你好",
			Age:  78,
		}
		_ = m[p2]
	}
	fmt.Println(time.Now().Sub(start))

	start = time.Now()
	for i := 0; i < 1000000; i++ {
		key := "你好"+"_"+"78"
		_ = strMap[key]
	}
	fmt.Println(time.Now().Sub(start))
}

func PrintMapInfo(m map[int]bool) {
	fmt.Println(len(m), m)
}
func TestMapCopy(t *testing.T) {
	m1 := map[int]bool{
		1: true,
		2: false,
	}
	PrintMapInfo(m1)
	m2 := m1
	m2[3] = false
	PrintMapInfo(m2)
	PrintMapInfo(m1)

}

type Term2PostingListMap struct {
	data map[string]int
}

func TestMapNil(t *testing.T) {
	term := map[string]int{
		"123456": 123,
		"654789": 243,
	}
	data := `{"123456":123,"54886211":6548,"654789":243,"985488":95644}`
	err := json.Unmarshal([]byte(data), &term)
	if err != nil {
		t.Error(err)
	}
	jb, _ := json.Marshal(term)
	fmt.Println(string(jb))

}

func getKey(n int) int {
	return n % 4
}

func TestArray(t *testing.T) {
	m := map[int][]int{}
	for i := 0; i < 100; i++ {
		k := getKey(i)
		arr, ok := m[k]
		fmt.Printf("%p ", arr)
		if !ok {
			arr = []int{i}
			m[k] = arr
		} else {
			arr = append(arr, i)
			fmt.Printf("%p \n", arr)
		}
		fmt.Println(m)
		time.Sleep(time.Second)
	}
}

func TestMapNew(t *testing.T) {
	m := new(map[string]bool)
	fmt.Println(m)
	v, ok := (*m)["123"]
	fmt.Println(v, ok)

}
