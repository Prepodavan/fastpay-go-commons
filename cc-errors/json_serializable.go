package cc_errors

import "encoding/json"

// Ошибка, строковым представлением которой яв-ся json
type JsonSerializableErr struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"data"`
}

func (jse *JsonSerializableErr) Error() string {
	b, _ := json.Marshal(jse)
	return string(b)
}
