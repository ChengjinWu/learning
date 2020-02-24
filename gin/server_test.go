package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	uri := fmt.Sprintf("uid=&cv=IK7.1.76_Android&live_uid=0&channel_id=4&tab_key=%s&offset=0", "123123")
	u := &url.URL{
		Scheme:   "https",
		Host:     "service.inke.cn",
		Path:     "/api/live/theme_card_recommend",
		RawQuery: uri,
	}
	fmt.Println(u)
}
