package dau

import "encoding/json"

type H map[string]interface{}

func (h H) String() string {
	b, _ := json.Marshal(h)
	return string(b)
}
