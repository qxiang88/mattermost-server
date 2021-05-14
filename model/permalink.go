package model

import "encoding/json"

type Permalink struct {
	PostMessage string `json:"post_message"`
}

func (o *Permalink) ToJson() string {
	b, _ := json.Marshal(o)
	return string(b)
}
