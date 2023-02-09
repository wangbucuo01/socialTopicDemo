package handler

import (
	"github.com/socialTopicDemo/service"
	"strconv"
)

func PublishPost(uidStr, topicIdStr, content string) *PageData {
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	topic, _ := strconv.ParseInt(topicIdStr, 10, 64)
	postId, err := service.PublishPost(topic, uid, content)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &PageData{
		Code: 0,
		Msg:  "success",
		Data: map[string]int64{
			"post_id": postId,
		},
	}
}
