package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"tripod/define"
)

func readFile(fileName string) [][]string {
	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	defer fi.Close()
	result := [][]string{}
	br := bufio.NewReader(fi)
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		lineStr := strings.Trim(string(line), " ")
		result = append(result, strings.Split(lineStr, ","))
	}
	return result
}
func main() {
	minArray := define.NewSet()
	maxArray := define.NewSet()
	ipv6s := readFile("ipv6.csv")
	count := 0
	for _, ipv6 := range ipv6s {
		if len(ipv6) != 3 {
			fmt.Println(ipv6)
		} else {
			//minPre := strings.Split(ipv6[0], ":")[:4]
			minPos := strings.Join(strings.Split(ipv6[0], ":")[4:], "")

			minIp := net.ParseIP(ipv6[0])
			maxIp := net.ParseIP(ipv6[1])
			//maxPre := strings.Split(ipv6[1], ":")[:4]
			maxPos := strings.Join(strings.Split(ipv6[1], ":")[4:], "")

			minArray.Add(minPos)
			maxArray.Add(minPos)
			count++
			if minPos != "0000" || maxPos != "ffffffffffffffff" {
				//fmt.Println(ipv6)
				fmt.Println(minIp, maxIp)
				fmt.Printf("%4s %4s %s\n", strings.Split(ipv6[0], ":"), strings.Split(ipv6[1], ":"), ipv6[2])
				//fmt.Println(minPos, maxPos, ipv6[2])
			}
		}
	}
	fmt.Println(minArray.String())
	fmt.Println(maxArray.String())
}
