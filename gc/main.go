package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

func main() {
	for {
		s := make([]int, 10)
		for i := 0; i < 1000000; i++ {
			s = append(s, i)
		}
		time.Sleep(time.Second)
		fmt.Println(len(s), unsafe.Sizeof(s))
		stats := runtime.MemStats{}
		runtime.ReadMemStats(&stats)
		jb, _ := json.Marshal(stats)
		fmt.Println(string(jb))
	}
}
