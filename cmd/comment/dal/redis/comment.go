package redis

import (
	"context"
	"douyin/kitex_gen/common"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GetCommentsForVideo(ctx context.Context, videoID int64) ([]*common.Comment, error) {
	comments := []*common.Comment{}

	// Try to fetch comments from Redis cache
	cacheKey := "video_comments:" + strconv.FormatInt(videoID, 10)
	fmt.Println(cacheKey)
	cacheData, err := RedisClient.Get(ctx, cacheKey).Bytes()
	if err == nil {
		// Comments were found in cache, return them
		err := json.Unmarshal(cacheData, &comments)
		if err != nil {
			return nil, err
		}
		return comments, nil
	}
	return nil, err
}

func SetComments(ctx context.Context, videoID int64, comments []*common.Comment) error {
	cacheKey := "video_comments:" + strconv.FormatInt(videoID, 10)
	// Insert comments into Redis cache with random expiry time
	cacheData, err := json.Marshal(comments)
	if err != nil {
		return err
	}
	expiryTime := time.Duration(rand.Intn(9)+1) * time.Minute
	err = RedisClient.Set(context.Background(), cacheKey, cacheData, expiryTime).Err()
	if err != nil {
		return err
	}
	return nil
}
