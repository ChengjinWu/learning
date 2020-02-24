package new

import (
	"fmt"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	var arr *[]int
	fmt.Println(arr)
	fmt.Printf("%s\n", "22")
	fmt.Println("#####################")
	arr = new([]int)
	fmt.Println(arr)
	fmt.Println(*arr)
	fmt.Printf("%s", "22")
}

func TestAlloc(t *testing.T) {
	type T struct {
		n  string
		i  int
		f  float64
		fd *os.File
		b  []byte
		s  bool
	}

	var t1 *T
	t1 = new(T)
	fmt.Println(t1)

	t2 := T{}
	fmt.Println(t2)

	t3 := T{"hello", 1, 3.1415926, nil, make([]byte, 2), true}
	fmt.Println(t3)
}
