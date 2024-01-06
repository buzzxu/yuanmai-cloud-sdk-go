package common

import "fmt"

type YuanmaiCloudSdkError struct {
	Code    int
	Message string
}

func (e YuanmaiCloudSdkError) Error() string {
	return fmt.Sprintf("[YuanmaiCloudSdkError] Code=%d, Message=%s", e.Code, e.Message)
}

func NewYuanmaiCloudSdkError(code int, message string) error {
	return &YuanmaiCloudSdkError{
		Code:    code,
		Message: message,
	}
}

func (e *YuanmaiCloudSdkError) GetCode() int {
	return e.Code
}

func (e *YuanmaiCloudSdkError) GetMessage() string {
	return e.Message
}
