package model

import "encoding/json"

type Permalink struct {
	PreviewPost *PreviewPost `json:"preview_post"`
}

type PreviewPost struct {
	Id        string `json:"id"`
	CreateAt  int64  `json:"create_at"`
	UpdateAt  int64  `json:"update_at"`
	ChannelId string `json:"channel_id"`
	Message   string `json:"message"`
}

func PreviewPostFromPost(post *Post) *PreviewPost {
	return &PreviewPost{
		Id:        post.Id,
		CreateAt:  post.CreateAt,
		UpdateAt:  post.UpdateAt,
		ChannelId: post.ChannelId,
		Message:   post.Message,
	}
}

func (o *Permalink) ToJson() string {
	b, _ := json.Marshal(o)
	return string(b)
}
