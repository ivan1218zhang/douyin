package db_test

import (
	"context"

	"douyin/cmd/comment/dal/db"
	"douyin/pkg/repository"
	"fmt"
	"testing"
)

func TestGetCommentListByVideoId(t *testing.T) {
	db.Init()
	ctx := context.Background()
	// Insert test data into database
	videoId1, videoId2 := 3, 4
	comments := []*repository.Comment{
		{
			UserId:  94053855959977984,
			VideoId: int64(videoId1),
			Content: "test1",
		},
		{
			UserId:  94053855959977984,
			VideoId: int64(videoId1),
			Content: "test2",
		},
		{
			UserId:  94053855959977984,
			VideoId: int64(videoId2),
			Content: "test3",
		},
	}
	for _, c := range comments {
		_, err := db.CreateComment(ctx, c)
		if err != nil {
			t.Errorf("CreateComment returned an error: %v", err)
		}
	}
	result, err := db.GetCommentListByVideoId(ctx, int64(videoId1))
	if err != nil {
		t.Fatalf("Error getting comment list: %v", err)
	}

	// Verify result
	if len(result) != 2 {
		t.Fatalf("Expected 2 comments, but got %d", len(result))
	}
	if result[0].Content != comments[0].Content || result[0].VideoId != comments[0].VideoId {
		t.Fatalf("Expected first comment to have text '%s' and video ID %d, but got text '%s' and video ID %d", comments[0].Content, comments[0].VideoId, result[0].Content, result[1].VideoId)
	}
	if result[1].Content != comments[1].Content || result[1].VideoId != comments[1].VideoId {
		t.Fatalf("Expected first comment to have text '%s' and video ID %d, but got text '%s' and video ID %d", comments[1].Content, comments[1].VideoId, result[1].Content, result[1].VideoId)
	}
}

func TestCreateComment(t *testing.T) {
	db.Init()
	ctx := context.Background()

	comment := &repository.Comment{
		UserId:  94053855959977984,
		VideoId: 1,
		Content: "test",
	}
	id, err := db.CreateComment(ctx, comment)
	if err != nil {
		t.Errorf("CreateComment returned an error: %v", err)
	}
	fmt.Println(id)
}

func TestDeleteComment(t *testing.T) {
	db.Init()
	ctx := context.Background()

	comment := &repository.Comment{
		UserId:  94053855959977984,
		VideoId: 2,
		Content: "test",
	}
	c, err := db.CreateComment(ctx, comment)
	if err != nil {
		t.Errorf("CreateComment returned an error: %v", err)
	}
	err = db.DeleteComment(ctx, c.ID)
	if err != nil {
		t.Errorf("DeleteComment returned an error: %v", err)
	}
}
