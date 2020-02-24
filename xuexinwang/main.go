package main

import (
	"bytes"
	"fmt"
	"git.inke.cn/BackendPlatform/golang/logging"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"regexp"
)

var (
	Slice     = []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	codeRe, _ = regexp.Compile(`/capachaimg.jpg\?cap=(\d+)`)
)

func HttpPostForm(url string, vs url.Values) error {
	client := &http.Client{}
	body := bytes.NewReader([]byte(vs.Encode()))
	req, err := http.NewRequest("Post", url, body)
	if err != nil {
		logging.Error(err)
		return err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.9(0x17000929) NetType/WIFI Language/zh_CN miniProgram")
	req.Header.Set("Host", "www.chsi.com.cn")
	req.Header.Set("Cookie", "_ga=GA1.3.682038847.1579664545; _gid=GA1.3.326016423.1579664545; CHSICC01=!tC5N+FchywFlrG0GGYWrKFjgWJfD//uat1iGzWJlVOL4/5pVeOwrO0VKUCWPsJNYMFzc8YKsmwD5iRg=; CHSICC_CLIENTFLAG3=d6d1556d12784a270db64720aa7652f7; JSESSIONID=3F26BC01E7BEB5B14EA3880AD1251B2F")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	return nil
}

func HttpGet(url string) (*goquery.Document, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logging.Error(err)
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.9(0x17000929) NetType/WIFI Language/zh_CN miniProgram")
	req.Header.Set("Host", "www.chsi.com.cn")
	req.Header.Set("Cookie", "_ga=GA1.3.682038847.1579664545; _gid=GA1.3.326016423.1579664545; CHSICC01=!tC5N+FchywFlrG0GGYWrKFjgWJfD//uat1iGzWJlVOL4/5pVeOwrO0VKUCWPsJNYMFzc8YKsmwD5iRg=; CHSICC_CLIENTFLAG3=d6d1556d12784a270db64720aa7652f7; JSESSIONID=3F26BC01E7BEB5B14EA3880AD1251B2F")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	return doc, nil
}

func main() {
	doc, err := HttpGet("http://www.chsi.com.cn/xlcx/bg.do?vcode=AR77CXNEU1M04YEU&from=wxxcx&nickName=%E6%88%90%E8%BF%9B&gender=1&language=zh_CN&city=Changping&province=Beijing&country=China&avatarUrl=https%3A%2F%2Fwx.qlogo.cn%2Fmmopen%2Fvi_32%2FTsibxjuGFwBY6N37ozpCuThkDTnqXmV6ia5AYlLvxkIAOagnTNV0KC1LkS7RG0uo7p1VkG57O5mabmaLYhnWsCmQ%2F132")
	if err != nil {
		return
	}

	//fmt.Println(doc.Html())
	codeHtml, err := doc.Find(".input-field > .yzCode").Html()
	if len(codeHtml) > 0 {
		codes := codeRe.FindStringSubmatch(codeHtml)
		if len(codes) == 2 {
			vs := url.Values{}
			vs.Set("cap", codes[1])
			vs.Set("capachatok", "0015796663877874eef5b4fd12e55b4cb8337a9b27a2e01&")
			vs.Set("Submit", "继续&")
			_ = HttpPostForm("http://www.chsi.com.cn/xlcx/yzm.do", vs)
			doc, err = HttpGet("http://www.chsi.com.cn/xlcx/bg.do?vcode=AR77CXNEU1M04YEU&from=wxxcx&nickName=%E6%88%90%E8%BF%9B&gender=1&language=zh_CN&city=Changping&province=Beijing&country=China&avatarUrl=https%3A%2F%2Fwx.qlogo.cn%2Fmmopen%2Fvi_32%2FTsibxjuGFwBY6N37ozpCuThkDTnqXmV6ia5AYlLvxkIAOagnTNV0KC1LkS7RG0uo7p1VkG57O5mabmaLYhnWsCmQ%2F132")
			if err != nil {
				return
			}
		}
	}
	fmt.Println(doc.Html())
	//
	//fmt.Println(doc.Find("#getXueLi > table > tbody > tr:nth-child(1) > td:nth-child(3) > img").Html())
	//
	//fmt.Println(doc.Find("#resultTable > div.div1 > div > table").Html())

}
