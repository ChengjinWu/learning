package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://weibo.com/tv/v/FxV0biitY?from=vfun", nil)
	if err != nil {
		// handle error
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.157 Safari/537.36")
	req.Header.Set("Referer", "https://weibo.com/tv/discovery")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}
