package rdb

import (
	"douyin/pkg/repository"
	"github.com/go-redis/redis"
	"math/rand"
	"strconv"
	"time"
)

const defaultExpireTime = 1 * time.Hour
const zsetKey = "video"

func getRandomExpireTime() time.Duration {
	return defaultExpireTime + time.Duration(rand.Intn(60))*time.Minute
}

func SetVideo(v repository.Video) error {
	if err := GetRDB().ZAdd(zsetKey, redis.Z{Score: float64(v.CreatedAt.Unix()), Member: v}).Err(); err != nil {
		return err
	}

	return GetRDB().Set(v.GetCacheKey(), v, getRandomExpireTime()).Err()
}

// GetVideoById return nil when error
func GetVideoById(id int64) repository.Video {
	v := repository.Video{}
	GetRDB().Get(v.GetCacheKey()).Scan(&v)

	return v
}

func MSetVideo(videoModels []*repository.Video) error {
	vs := make([]redis.Z, len(videoModels))

	for i := range vs {
		vs[i] = redis.Z{Score: float64(videoModels[i].CreatedAt.Unix()), Member: videoModels[i]}
	}

	return GetRDB().ZAdd(zsetKey, vs...).Err()
}

func MGetVideoIDByTime(latestTime int64) ([]int64, error) {
	opt := redis.ZRangeBy{
		Max:   strconv.FormatInt(latestTime, 10), // 最大分数
		Count: 10,                                // 一次返回多少数据
	}
	vs := make([]int64, 10)
	//根据opt范围返回集合元素
	err := GetRDB().ZRevRangeByScore(zsetKey, opt).ScanSlice(&vs)
	return vs, err
}

func MGetVideoByTime(latestTime int64) []*repository.Video {
	opt := redis.ZRangeBy{
		Max:   strconv.FormatInt(latestTime, 10), // 最大分数
		Count: 10,                                // 一次返回多少数据
	}
	vs := make([]*repository.Video, 10)
	//根据opt范围返回集合元素
	GetRDB().ZRevRangeByScore(zsetKey, opt).ScanSlice(&vs)
	return vs
}
