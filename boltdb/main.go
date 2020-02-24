package main

import "learning/boltdb/util"

func main() {
	//now := time.Now()
	//for i := 0; i < 1000; i++ {
	//	csMap := make(map[int]map[int]int)
	//	n := rand.Intn(5)
	//	for k := 0; k < n; k++ {
	//		cid := rand.Intn(5)
	//		statusCountMap := make(map[int]int)
	//		nn := rand.Intn(10)
	//		for j := 0; j < nn; j++ {
	//			status := rand.Intn(5)
	//			statusCountMap[status] = rand.Intn(100)
	//		}
	//		csMap[cid] = statusCountMap
	//	}
	//
	//	go util.AddCreativeStatus(csMap, now.Add(time.Duration(rand.Intn(20))*time.Second))
	//}
	//time.Sleep(10 * time.Second)
	//for i := 0; i < 20; i++ {
	//	currTime := now.Add(time.Duration(i) * time.Second)
	//	fmt.Println(currTime.Format(util.TimeFormatStr))
	//	fmt.Println(util.QueryCreativeStatus(currTime))
	//}
	util.PrintBoltDb()
	//startTime, _ := time.ParseInLocation(util.TimeFormatStr, "2019-06-11 16:10:09", time.Local)
	//endTime, _ := time.ParseInLocation(util.TimeFormatStr, "2019-06-11 16:19:27", time.Local)
	//log.Info(startTime, endTime)
	//log.Info(util.QueryCreativeStatus(startTime, endTime))
}
