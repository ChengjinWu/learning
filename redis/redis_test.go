package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"regexp"
	"strings"
	"testing"
	"time"
)

var (
	lineSpilt, _ = regexp.Compile(`\s+`)
)

func TestAddData(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	for i := 0; i < 10000000; i++ {
		key := fmt.Sprintf("adfrequence-%d", 1000000000+rand.Intn(1000000000))
		//value := fmt.Sprintf("3;1900-01-01;火星:%d", 1000000000+rand.Intn(1000000000))
		err := client.Set(key, 1000000000+rand.Intn(1000000000), 2*12*30*24*time.Hour).Err()
		if err != nil {
			fmt.Println(err)
		}
		if i%2000 == 0 {
			ret := client.Info("Keyspace")
			lineRet := lineSpilt.Split(ret.String(), -1)
			fmt.Print(strings.Split(strings.Split(lineRet[4], ":")[1], ",")[0])
			ret = client.Info("memory")
			lineRet = lineSpilt.Split(ret.String(), -1)
			fmt.Println(lineRet[6])
		}
		time.Sleep(50 * time.Microsecond)
	}
}

func TestIncrby(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	for i := 0; i < 1000; i++ {
		key := "longxingtianxia"
		ret := client.IncrBy(key, 1)
		if ret.Err() != nil {
			fmt.Println("SetAdShowCount err:", err)
		} else {
			fmt.Println(ret.Result())
			if ret.Val() == 1 {
				cmd := client.Expire(key, 24*time.Hour)
				if cmd.Err() != nil {
					fmt.Println("Set AdShow ttl err:", err)
				}
			}
		}
	}
}
