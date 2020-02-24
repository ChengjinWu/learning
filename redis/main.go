package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"runtime"
	"sync"
	"time"
)

var w sync.WaitGroup

func newRdsPool(server, auth string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     100,
		MaxActive:   30,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if auth == "" {
				return c, err
			}
			if _, err := c.Do("AUTH", auth); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
func g1(r redis.Conn) {
	for i := 0; i < 2; i++ {
		if _, err := redis.String(r.Do("set", "hello", "1")); err != nil {
			fmt.Println(err)
		}
		time.Sleep(10 * time.Millisecond)
	}
	w.Done()
}
func g2(r redis.Conn) {
	for i := 0; i < 2; i++ {
		if _, err := redis.String(r.Do("set", "hello", "2")); err != nil {
			fmt.Println(err)
		}
		time.Sleep(10 * time.Millisecond)
	}
	w.Done()
}
func main() {
	w.Add(2)
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rc1 redis.Conn = newRdsPool(`127.0.0.1:6379`, ``).Get()
	var rc2 redis.Conn = newRdsPool(`127.0.0.1:6379`, ``).Get()
	defer rc1.Close()
	defer rc2.Close()
	go g1(rc1)
	go g2(rc2)
	w.Wait()
}
