package rdb

import (
	"strconv"
)

func CountFavorite(videoID int64) (int, error) {
	rdb := GetRDB()
	numStr, err := rdb.Get("video_favorite_" + strconv.FormatInt(videoID, 10)).Result()
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(numStr)

	return num, err
}

func MCountFavorite(videoIDs []int64) ([]int64, error) {
	rdb := GetRDB()
	keys := make([]string, len(videoIDs))
	for i := range keys {
		keys[i] = "video_favorite_" + strconv.FormatInt(videoIDs[i], 10)
	}
	numsIn, err := rdb.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}

	nums := make([]int64, len(videoIDs))
	for i := range nums {
		nums[i] = numsIn[i].(int64)
	}

	return nums, err
}

// SetFavorite set video into user favorite list , incr video favorite count
func SetFavorite(userID int64, videoID int64) error {
	rdb := GetRDB()

	if err := rdb.SAdd("user_favorite_"+strconv.FormatInt(userID, 10), videoID).Err(); err != nil {
		return err
	}
	rdb.Incr("video_favorite_" + strconv.FormatInt(videoID, 10))

	return nil
}

// DelFavorite delete the video in user favorite list , decr video favorite count
func DelFavorite(userID int64, videoID int64) error {
	rdb := GetRDB()

	if err := rdb.SRem("user_favorite_"+strconv.FormatInt(userID, 10), videoID).Err(); err != nil {
		return err
	}
	rdb.Decr("video_favorite_" + strconv.FormatInt(videoID, 10))

	return nil
}

func IsFavorite(userID int64, videoID int64) (bool, error) {
	rdb := GetRDB()

	isLiked, err := rdb.SIsMember("user_favorite_"+strconv.FormatInt(userID, 10), userID).Result()
	return isLiked, err
}

func MIsFavorite(userID int64, videoIDs []int64) ([]bool, error) {
	favMap, err := GetRDB().SMembersMap("user_favorite_" + strconv.FormatInt(userID, 10)).Result()
	if err != nil {
		return nil, err
	}

	isFav := make([]bool, len(videoIDs))
	for i, id := range videoIDs {
		if _, ok := favMap[strconv.FormatInt(id, 10)]; ok {
			isFav[i] = true
		}
	}

	return isFav, nil
}

func MGetFavoriteVideoID(userID int64) ([]int64, error) {
	rdb := GetRDB()

	vIDStrings, err := rdb.SMembers("user_favorite_" + strconv.FormatInt(userID, 10)).Result()
	vID := make([]int64, len(vIDStrings))

	for i := range vIDStrings {
		vID[i], _ = strconv.ParseInt(vIDStrings[i], 10, 64)
	}

	return vID, err
}
