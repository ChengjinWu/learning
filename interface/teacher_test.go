package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowA() {
	fmt.Println("teachershowA")
	t.People.ShowB()
	//p.ShowB()
}
func (t *Teacher) ShowB() {
	fmt.Println("teachershowB")
}
func TestMainNihao(t *testing.T) {
	te := Teacher{}
	te.ShowA()
	te.ShowB()
}

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{}) // 解除注释看看！
	//ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()
		for elem, value := range set.s {
			ch <- elem
			println("Iter:", elem, value)
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}
func TestMainNihao22(t *testing.T) {
	th := threadSafeSet{
		s: []interface{}{"1", "2", "3", "4"},
	}
	v := <-th.Iter()
	fmt.Printf("%s%v", "ch", v)
	time.Sleep(time.Second)
}
