package tool

import (
	"encoding/json"
	"io"
)

type JsonParse struct {
}

//封装一个参数解析方法
func Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}
