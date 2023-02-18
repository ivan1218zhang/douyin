package pack

import (
	"douyin/kitex_gen/common"
	"douyin/pkg/repository"
)

func Comment(c *repository.Comment) *common.Comment {
	if c == nil {
		return nil
	}
	return &common.Comment{
		Id:         c.ID,
		UserId:     c.UserId,
		Content:    c.Content,
		CreateDate: c.CreatedAt.Format("01-02"),
	}
}

func Comments(cs []*repository.Comment) []*common.Comment {
	comments := make([]*common.Comment, 0)
	for _, c := range cs {
		if comment := Comment(c); comment != nil {
			comments = append(comments, comment)
		}
	}
	return comments
}
