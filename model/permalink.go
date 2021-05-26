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
	Type      string `json:"type"`
	UserID    string `json:"user_id"`
}

func PreviewPostFromPost(post *Post) *PreviewPost {
	return &PreviewPost{
		Id:        post.Id,
		CreateAt:  post.CreateAt,
		UpdateAt:  post.UpdateAt,
		ChannelId: post.ChannelId,
		Message:   post.Message,
		Type:      post.Type,
		UserID:    post.UserId,
	}
}

func (o *Permalink) ToJson() string {
	b, _ := json.Marshal(o)
	return string(b)
}
