package pack

import (
	common2 "douyin/cmd/api/biz/model/common"
	"douyin/kitex_gen/common"
)

func Comment(comment *common.Comment) *common2.Comment {
	return &common2.Comment{
		ID:         comment.Id,
		User:       User(comment.User),
		Content:    comment.Content,
		CreateDate: comment.CreateDate,
		UserID:     comment.UserId,
	}
}

func CommentList(comments []*common.Comment) []*common2.Comment {
	res := make([]*common2.Comment, len(comments))
	for index, comment := range comments {
		res[index] = &common2.Comment{
			ID:         comment.Id,
			User:       User(comment.User),
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
			UserID:     comment.UserId,
		}
	}
	return res
}
