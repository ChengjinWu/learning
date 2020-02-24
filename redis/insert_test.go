package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

var (
	uids      = []int{12406400, 136163, 24880001, 14882031, 3187468, 25360527, 26048507, 2597345, 5271280, 19812314, 26237126, 25010657, 18109505, 25326144, 25771475, 25874948, 23412751, 25676929, 25771618, 4138415, 4138378, 4138214, 7217308, 7217284}
	geos      = []string{"北京", "上海", "广东", "南阳", "商丘", "商丘", "深圳", "龙城"}
	interests = []string{"arder", "chess", "manufacture", "square_dance", "body", "game", "constellation", "cute", "dance", "farmer", "news", "opera", "car", "film", "painting", "science", "music", "pe", "tech", "anecdote", "education", "fashion", "variety", "amusement", "motherbaby", "special", "finance", "funny", "comic", "emotion", "tourism", "animal", "beauty", "crosstalk", "health", "life", "military", "social", "career", "cartoon", "collection", "culture", "lang", "cate"}
)

/*
  addr="10.100.128.16:6379"
  password="jgoi4874UODTH"
*/
func TestInsertData(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.100.128.16:6379",
		Password: "jgoi4874UODTH", // no password set
		DB:       0,               // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	for _, uid := range uids {
		key := fmt.Sprintf("user_info_%d", uid)
		value := getRandKeyValue()
		fmt.Println(key, value)
		err := client.Set(key, value, 0).Err()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getRandKeyValue() string {
	ret := make([]string, 4)
	ret[0] = strconv.Itoa(rand.Intn(3) + 1)
	ret[1] = fmt.Sprintf("%d-01-01", time.Now().Year()-rand.Intn(80))
	ret[2] = geos[rand.Intn(len(geos))]
	ret[3] = interests[rand.Intn(len(interests))]
	return strings.Join(ret, ";")

}
