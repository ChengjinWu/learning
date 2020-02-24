package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/lunny/log"
	"testing"
	"time"
)

func TestRedisTraversing(t *testing.T) {
	log.SetOutputLevel(1)
	client := redis.NewClient(&redis.Options{
		Addr:     "10.100.13.158:6379",
		Password: "m1GwbBzf6uvm", // no password set
		DB:       0,              // use default DB
	})
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "10.100.128.16:6379",
	//	Password: "jgoi4874UODTH", // no password set
	//	DB:       0,               // use default DB
	//})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	var cursor uint64 = 0
	for {
		var keys []string
		var err error
		keys, cursor, err = client.Scan(cursor, "freq-*", 1000).Result()
		if err != nil {
			log.Error(err)
			break
		}
		log.Debugf("%10d %v", cursor, keys)

		pipeline := client.Pipeline()
		for _, key := range keys {
			pipeline.Do("ttl", key)
		}
		cmders, err := pipeline.Exec()
		if err != nil {
			log.Error(err)
			break
		}
		for _, cmder := range cmders {
			cmd := cmder.(*redis.Cmd)
			ttl, err := cmd.Int64()
			if err != nil {
				log.Error(err)
				continue
			}
			if ttl == -1 {
				fmt.Println(fmt.Sprintf("%10d %10d %s", cursor, ttl, cmd.Args()[1]))
				//pipeline.Expire(cmd.Args()[1].(string), 24*time.Hour)
			}
		}
		cmders, err = pipeline.Exec()
		if err != nil {
			log.Error(err)
			break
		}
		for _, cmder := range cmders {
			boolCmd := cmder.(*redis.BoolCmd)
			ok, err := boolCmd.Result()
			if err != nil || !ok {
				log.Error(ok, err)
				continue
			}
		}
		time.Sleep(10 * time.Millisecond)
		if cursor == 0 {
			break
		}
	}
}
