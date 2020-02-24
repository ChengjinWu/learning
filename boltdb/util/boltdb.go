package util

import (
	"bytes"
	"encoding/json"
	"github.com/boltdb/bolt"
	"github.com/lunny/log"
	"strconv"
	"time"
)

var (
	boltDb *bolt.DB
)

func init() {
	dbFilePath := "./advertising.bolt"
	var err error
	boltDb, err = bolt.Open(dbFilePath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
}

type CreativeStatusInfo struct {
	CreativeId  int
	AdId        int
	CampId      int
	AdvId       int
	StatusCount map[int]int
	At          time.Time
}

func AddCreativeStatus(reqCsMap map[int]map[int]int, currTime time.Time) error {
	jb, _ := json.Marshal(reqCsMap)
	log.Info(currTime.Format(TimeFormatStr), string(jb))
	if len(reqCsMap) == 0 {
		return nil
	}
	bucketName := currTime.Format(DateFormatStr)
	err := boltDb.Update(func(tx *bolt.Tx) error {
		log.Infof("%d tx start", tx.ID())
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			log.Error(err)
			return err
		}

		// 存值
		key := []byte(strconv.Itoa(int(currTime.Unix())))
		value := bucket.Get(key)
		creativeStatus := make(map[int]map[int]int)
		if value != nil {
			err = json.Unmarshal(value, &creativeStatus)
			if err != nil {
				log.Error(err)
				return err
			}
			for cid, reqStatusCountMap := range reqCsMap {
				statusCount, ok := creativeStatus[cid]
				if !ok {
					creativeStatus[cid] = reqStatusCountMap
				} else {
					for status, count := range reqStatusCountMap {
						statusCount[status] = statusCount[status] + count
					}
				}
			}
		} else {
			creativeStatus = reqCsMap
		}
		value, err = json.Marshal(creativeStatus)
		if err != nil {
			log.Error(err)
			return err
		}
		err = bucket.Put(key, value)
		if err != nil {
			log.Error(err)
			return err
		}
		log.Infof("%d tx end", tx.ID())
		return nil
	})
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func QueryCreativeStatusByCurrTime(currTime time.Time) (map[int]map[int]int, error) {
	creativeStatus := make(map[int]map[int]int)
	bucketName := currTime.Format(DateFormatStr)
	err := boltDb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return nil
		}
		key := []byte(strconv.Itoa(int(currTime.Unix())))
		value := bucket.Get(key)
		if value != nil {
			err := json.Unmarshal(value, &creativeStatus)
			if err != nil {
				log.Error(err)
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return creativeStatus, nil
}

func QueryCreativeStatus(startTime, endTime time.Time) (map[int]map[int]int, error) {
	res := make(map[int]map[int]int)
	startBucketName := startTime.Format(DateFormatStr)
	endBucketName := endTime.Format(DateFormatStr)
	err := boltDb.View(func(tx *bolt.Tx) error {
		err := tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			if startBucketName > string(name) || string(name) > endBucketName {
				return nil
			}
			cursor := b.Cursor()
			min := []byte(strconv.Itoa(int(startTime.Unix())))
			max := []byte(strconv.Itoa(int(endTime.Unix())))
			for k, v := cursor.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = cursor.Next() {
				log.Infof("%s: %s\n", k, v)
				if len(v) == 0 {
					continue
				}
				creativeStatusMap := make(map[int]map[int]int)
				err := json.Unmarshal(v, &creativeStatusMap)
				if err != nil {
					log.Error(err)
					continue
				}
				for cid, reqStatusCountMap := range creativeStatusMap {
					resStatusCount, ok := res[cid]
					if !ok {
						res[cid] = reqStatusCountMap
					} else {
						for status, count := range reqStatusCountMap {
							resStatusCount[status] = resStatusCount[status] + count
						}
					}
				}
			}
			return nil
		})
		return err
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return res, nil
}

func PrintBoltDb() {
	err := boltDb.View(func(tx *bolt.Tx) error {
		err := tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			err := b.ForEach(func(k, v []byte) error {
				sec, _ := strconv.Atoi(string(k))
				currTime := time.Unix(int64(sec), 0)
				log.Infof("bucket:%s key:%s,%s value:%s", name, k, currTime.Format(TimeFormatStr), v)
				return nil
			})
			return err
		})
		return err
	})
	if err != nil {
		log.Error(err)
	}
}
