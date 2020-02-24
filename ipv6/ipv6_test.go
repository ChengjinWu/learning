package main

import (
	"fmt"
	"math/rand"
	"net"
	"sort"
	"testing"
)

var (
	v4InV6Prefix = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff}
)

func IPLess(ipx, ipy net.IP) bool {
	if len(ipx) == len(ipy) {
		return bytesLessCompare(ipx, ipy)
	}
	if len(ipx) == net.IPv4len && len(ipy) == net.IPv6len {
		return bytesLessCompare(ipy[0:12], v4InV6Prefix) && bytesLessCompare(ipy, ipy[12:])
	}
	if len(ipx) == net.IPv6len && len(ipy) == net.IPv4len {
		return bytesLessCompare(ipx[0:12], v4InV6Prefix) && bytesLessCompare(ipx[12:], ipy)
	}
	return false
}

func bytesLessCompare(x, y []byte) bool {
	if len(x) != len(y) {
		return false
	}
	for i, b := range x {
		if b < y[i] {
			return true
		} else if b > y[i] {
			return false
		}
	}
	return false
}

type Ipv6Coding struct {
	Min  net.IP
	Max  net.IP
	Area string
}
type Ipv6Array []Ipv6Coding

func (s Ipv6Array) Len() int {
	return len(s)
}
func (s Ipv6Array) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Ipv6Array) Less(i, j int) bool {
	return IPLess(s[i].Min, s[j].Min)
}
func GetArea(array Ipv6Array, ip string) string {
	currIp := net.ParseIP(ip)
	if IPLess(currIp, net.IPv6zero) {
		return ""
	}
	for left, right := 0, len(array)-1; left <= right; {
		mid := (left + right) >> 1
		record := array[mid]
		if IPLess(currIp, record.Min) {
			right = mid - 1
		} else if IPLess(record.Max, currIp) {
			left = mid + 1
		} else {
			fmt.Println(record.Min, record.Max)
			return record.Area
		}
	}
	return ""
}
func TestDemo(t *testing.T) {
	ipv6s := readFile("IPV6.csv")
	array := Ipv6Array{}
	for _, ipv6 := range ipv6s {
		if len(ipv6) != 3 {
			fmt.Println(ipv6)
		} else {
			array = append(array, Ipv6Coding{
				Min:  net.ParseIP(ipv6[0]),
				Max:  net.ParseIP(ipv6[1]),
				Area: ipv6[2],
			})
		}
	}
	sort.Sort(array)
	fmt.Println(sort.IsSorted(array))
	for i := 0; i < 100; i++ {
		ip := GetDoman()
		fmt.Println(GetArea(array, ip), ip)
	}

}
func GetDoman() string {
	array := []byte{}
	for i := 0; i < 16; i++ {
		array = append(array, byte(rand.Intn(255)))
	}
	var ip net.IP = array
	return ip.String()
}

func TestDemo333(t *testing.T) {
	fmt.Println(GetDoman())
}
