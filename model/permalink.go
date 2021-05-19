package model

import "encoding/json"

type Permalink struct {
	LinkedPost *Post `json:"linked_post"`
}

func (o *Permalink) ToJson() string {
	b, _ := json.Marshal(o)
	return string(b)
}
