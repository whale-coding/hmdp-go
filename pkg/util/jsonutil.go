package util

import (
	"encoding/json"
)

// Marshal 序列化对象为 JSON 字符串
func Marshal(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Unmarshal 反序列化 JSON 字符串为对象
func Unmarshal(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}
