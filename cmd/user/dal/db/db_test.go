package db_test

import (
	"context"
	"douyin/cmd/user/dal/db"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db.Init()
	ctx := context.Background()

	// Test case : check if user is created successfully
	users := []*db.User{
		{
			UserName: "testuser1",
			Password: "test",
		},
	}
	if err := db.CreateUser(ctx, users); err != nil {
		t.Errorf("CreateUser returned an error: %v", err)
	}
}
