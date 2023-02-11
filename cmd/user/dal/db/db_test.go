package db_test

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/pkg/repository"
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db.Init()
	ctx := context.Background()

	// Test case : check if user is created successfully
	user := &repository.User{
		UserName: "testuser1",
		Password: "test",
	}
	id, err := db.CreateUser(ctx, user)
	if err != nil {
		t.Errorf("CreateUser returned an error: %v", err)
	}
	fmt.Println(id)

}
