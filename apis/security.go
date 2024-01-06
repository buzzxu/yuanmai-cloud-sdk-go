package apis

import (
	"github.com/buzzxu/yuanmai-cloud-sdk-go/sdk"
)

type (
	SecurityApi interface {
		// Captcha 图形验证码
		Captcha(request *CaptchaRequest) (*CaptchaResponse, error)
		// Login 登录
		Login(request *LoginRequest) (map[string]interface{}, error)
		// Logout 登出
		Logout(request *LogoutRequest) (bool, error)
	}

	OpenSecurityApi struct {
		client  *sdk.Client
		Service string
	}

	// CaptchaRequest 图形验证码
	CaptchaRequest struct {
		Key string `json:"key"`
	}
	// LoginRequest 登录
	LoginRequest struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
		Key      string `json:"key"`
		ValiCode string `json:"valiCode"`
		Region   int    `json:"region"`
		Lang     string `json:"lang"`
	}

	// LogoutRequest 登出
	LogoutRequest struct {
		UserId        int64  `json:"userId"`
		Authorization string `json:"authorization"`
	}

	// CaptchaResponse 图像验证码
	CaptchaResponse struct {
		Key    string `json:"key"`
		Base64 string `json:"base64"`
	}
)

func (t *OpenSecurityApi) Captcha(request *CaptchaRequest) (*CaptchaResponse, error) {
	response := &sdk.BaseResponse[CaptchaResponse]{}
	err := t.client.Send(t.Service, "captcha", request, response)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func (t *OpenSecurityApi) Login(request *LoginRequest) (map[string]interface{}, error) {
	response := &sdk.BaseResponse[map[string]interface{}]{}
	err := t.client.Send(t.Service, "login", request, response)
	if err != nil {
		return nil, err
	}
	return response.Data, err
}

func (t *OpenSecurityApi) Logout(request *LogoutRequest) (bool, error) {
	response := &sdk.BaseResponse[bool]{}
	err := t.client.Send(t.Service, "logout", request, response)
	if err != nil {
		return false, err
	}
	return response.Data, err
}
