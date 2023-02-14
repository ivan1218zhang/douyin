package redis

import (
	"context"
	"douyin/cmd/comment/dal/db"
	"douyin/pkg/repository"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

func GetCommentsForVideo(ctx context.Context, videoID int64) ([]*repository.Comment, error) {
	comments := []*repository.Comment{}

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

	// Comments were not found in cache, fetch from database
	comments, err = db.GetCommentListByVideoId(ctx, videoID)
	if err != nil {
		return nil, err
	}

	// Insert comments into Redis cache with random expiry time
	cacheData, err = json.Marshal(comments)
	if err != nil {
		return nil, err
	}
	expiryTime := time.Duration(rand.Intn(9)+1) * time.Minute
	err = RedisClient.Set(context.Background(), cacheKey, cacheData, expiryTime).Err()
	if err != nil {
		return nil, err
	}

	return comments, nil
}
