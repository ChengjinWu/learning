package main

import (
	"fmt"
	"github.com/lunny/log"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	phoneRe, _ = regexp.Compile(`\d{11}`)
)

func main() {
	urlFormat := "https://m.10010.com/NumApp/NumberCenter/qryNum?callback=jsonp_queryMoreNums&provinceCode=11&cityCode=110&monthFeeLimit=0&groupKey=7200310618&searchCategory=3&net=01&amounts=200&codeTypeCode=&searchValue=&qryType=02&goodsNet=4&_=%d"
	phoneSet := make(map[string]bool)
	page := 1564539359595
	for i := 0; i < 1000; i++ {
		resp, err := http.Get(fmt.Sprintf(urlFormat, page))
		if err != nil {
			log.Error(err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error(err)
		}
		phones := phoneRe.FindAllString(string(body), -1)
		for _, phone := range phones {
			phoneSet[phone] = true
		}
		page++
	}
	for key, _ := range phoneSet {
		fmt.Println(key)
	}
}
