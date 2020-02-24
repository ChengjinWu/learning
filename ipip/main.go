package main

import (
	"fmt"
	"github.com/ipipdotnet/ipdb-go"
	"log"
)

func main() {
	db, err := ipdb.NewCity("src/learning/ipip/ipipfree.ipdb")
	if err != nil {
		log.Fatal(err)
	}

	//db.Reload("/path/to/city.ipv4.ipdb") // 更新 ipdb 文件后可调用 Reload 方法重新加载内容

	fmt.Println(db.IsIPv4())    // check database support ip type
	fmt.Println(db.IsIPv6())    // check database support ip type
	fmt.Println(db.BuildTime()) // database build time
	fmt.Println(db.Languages()) // database support language
	fmt.Println(db.Fields())    // database support fields

	fmt.Println(db.FindInfo("2001:250:200::", "CN")) // return CityInfo
	fmt.Println(db.Find("1.1.1.1", "CN"))            // return []string
	fmt.Println(db.FindMap("118.28.8.8", "CN"))      // return map[string]string
	fmt.Println(db.FindMap("223.104.246.114", "CN")) // return map[string]string
	fmt.Println(db.FindInfo("127.0.0.1", "CN"))      // return CityInfo

	fmt.Println()
}
