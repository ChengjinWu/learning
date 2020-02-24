package time

import (
	"fmt"
	"testing"
	"time"
)

func Test_Ticker(t *testing.T) {
	count := 0
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), count)
		default:
			count++
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func Test_Timer(t *testing.T) {
	d := time.Duration(time.Second * 2)

	timer := time.NewTimer(d)
	defer timer.Stop()

	for {
		<-timer.C

		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "timeout...")
		// need reset
		timer.Reset(time.Second * 2)
	}
}
func Test_Ticker1(t *testing.T) {
	d := time.Duration(time.Second * 2)

	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for {
		<-ticker.C

		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "timeout...")
	}
}

func TestMonths(t *testing.T) {

	firstLoginTime, err := time.Parse("2006-01-02", "2020-01-02")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(firstLoginTime)
}

func Test_TickerAfter(t *testing.T) {
	count := 0
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), count)
		default:
			count++
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func TestMonths222(t *testing.T) {
	totalTimer := time.NewTicker(4 * time.Second)
	realTimer := time.NewTicker(1 * time.Second)
	for true {
		select {
		case <-totalTimer.C:
			fmt.Println("total time")
		case <-realTimer.C:
			fmt.Println("real time")
		}
	}
}
