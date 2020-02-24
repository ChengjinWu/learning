package iota

import (
	"fmt"
	"testing"
)

const (

	maxQueueSize1 = 10
	maxQueueSize = 10
	TotalDataUpdate int = iota
	RealDataUpdate
)

func TestIota(t *testing.T)  {
	fmt.Println(TotalDataUpdate)
}
