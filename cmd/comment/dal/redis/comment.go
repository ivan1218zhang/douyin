package redis

import (
	"context"
	"douyin/kitex_gen/common"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

func GetCommentsForVideo(ctx context.Context, videoID int64) ([]*common.Comment, error) {
	comments := []*common.Comment{}

	// Try to fetch comments from Redis cache
	cacheKey := "video_comments:" + strconv.FormatInt(videoID, 10)
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

func GetCommentCount(ctx context.Context, videoID int64) (int64, error) {
	cacheKey := "comment_count:" + strconv.FormatInt(videoID, 10)
	cacheData, err := RedisClient.Get(ctx, cacheKey).Bytes()
	if err == nil {
		count, err := strconv.Atoi(string(cacheData))
		if err != nil {
			return -1, err
		}
		return int64(count), nil
	}
	return -1, err
}

func SetComments(ctx context.Context, videoID int64, comments []*common.Comment) error {
	cacheKey := "video_comments:" + strconv.FormatInt(videoID, 10)
	// Insert comments into Redis cache with random expiry time
	cacheData, err := json.Marshal(comments)
	if err != nil {
		return err
	}
	expiryTime := time.Duration(rand.Intn(9)+1) * time.Minute
	err = RedisClient.Set(ctx, cacheKey, cacheData, expiryTime).Err()
	if err != nil {
		return err
	}
	return nil
}

func SetCommentCount(ctx context.Context, videoID int64, count int64) error {
	cacheKey := "comment_count:" + strconv.FormatInt(videoID, 10)

	expiryTime := time.Duration(rand.Intn(9)+1) * time.Minute
	err := RedisClient.Set(ctx, cacheKey, count, expiryTime).Err()
	if err != nil {
		return err
	}
	return nil
}
