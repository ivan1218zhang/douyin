package pack

import (
	common2 "douyin/cmd/api/biz/model/common"
	"douyin/kitex_gen/message"
)

func MessageList(messages []*message.Message) []*common2.Message {
	res := make([]*common2.Message, len(messages))
	for index, message1 := range messages {
		res[index] = &common2.Message{
			ID:         message1.Id,
			ToUserID:   message1.ToUserId,
			FromUserID: message1.FromUserId,
			Content:    message1.Content,
			CreateTime: message1.CreateTime,
		}
	}
	return res
}
