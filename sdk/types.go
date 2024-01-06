package sdk

import (
	"encoding/json"
	"io"
)

type (
	Response interface {
		ToJsonString() string
		FromJsonString(s string) error
		FromJsonReader(reader io.Reader) error
		error() bool
		code() int
		message() string
		key() string
	}
	BaseResponse[T any] struct {
		Code    int    `json:"code"`
		Error   bool   `json:"error"`
		Key     string `json:"key,omitempty"`
		Message string `json:"message"`
		Data    T      `json:"data,omitempty"`
	}
)

func (r *BaseResponse[T]) error() bool {
	return r.Error
}
func (r *BaseResponse[T]) code() int {
	return r.Code
}
func (r *BaseResponse[T]) message() string {
	return r.Message
}
func (r *BaseResponse[T]) key() string {
	return r.Key
}

func (r *BaseResponse[T]) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}
func (r *BaseResponse[T]) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
func (r *BaseResponse[T]) FromJsonReader(reader io.Reader) error {
	return json.NewDecoder(reader).Decode(&r)
}
