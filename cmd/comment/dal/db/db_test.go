package db_test

import (
	"context"

	"douyin/cmd/comment/dal/db"
	"douyin/pkg/repository"
	"fmt"
	"testing"
)

func TestGetCommentListByVideoId(t *testing.T) {

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
