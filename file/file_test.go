package main_test

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
	"testing"
)

func ReadFile(filePath string) (array []string) {
	fi, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		array = append(array, string(a))
	}
	return
}

func repeat(array []string) string {
	res := make([]string, 0)
	repeat := make(map[string]bool)
	for _, v := range array {
		if !repeat[v] {
			res = append(res, v)
			repeat[v] = true
		}
	}
	if len(res) == 0 {
		return  "_"
	} else  {
		return strings.Join(res,",")
	}
}

func TestDemo22(t *testing.T) {
	uriArray := ReadFile("/Users/wuchengjin/Desktop/nihao.log")

	for _, uri := range uriArray {
		values, err := url.ParseQuery(uri[1:])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%v\t%v\t%v\n", repeat(values["channel_id"]), repeat(values["rec_tab"]), repeat(values["keyword"]))
		}
	}
}
