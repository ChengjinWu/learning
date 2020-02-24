package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "inke_monitor:inke_monitor@tcp(10.128.0.2:3306)/task_activity?charset=utf8mb4&parseTime=true&loc=Local&readTimeout=1s")
	if err != nil {
		panic(err)
	}
	tx := db.Begin()
	tx.Where("id = 100229"
	defer db.Close()
}
