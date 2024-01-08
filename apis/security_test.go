package apis

import (
	"github.com/buzzxu/yuanmai-cloud-sdk-go/sdk"
	"testing"
)

var credential = sdk.NewCredential("xwjd", "XWJD010230441222212312313V")

func TestSecurityCaptcha(t *testing.T) {
	client := sdk.NewDefaultClient("api.mgr.xwjd.xingchenga.xyz")
	client.WithCredential(credential)
	securityApi := NewOpenSecurityApi(client)
	// 图形验证码
	captchaRequest := &CaptchaRequest{
		Key: "",
	}
	captchaResponse, err := securityApi.Captcha(captchaRequest)
	if err != nil {
		t.Fatalf("调用失败: %s", err.Error())
	}
	t.Logf("返回结果: %v", captchaResponse)
}

func TestOpenSecurityApi_Login(t *testing.T) {
	client := sdk.NewDefaultClient("api.mgr.xwjd.xingchenga.xyz")
	client.WithCredential(credential)
	securityApi := NewOpenSecurityApi(client)
	// 图形验证码
	loginRequest := &LoginRequest{
		UserName: "15666666660",
		Password: "111111",
		Region:   0,
		Key:      "643132",
		ValiCode: "0532",
	}
	info, err := securityApi.Login(loginRequest)
	if err != nil {
		t.Fatalf("调用失败: %s", err.Error())
	}
	t.Logf("返回结果: %v", info)
}

func TestOpenSecurityApi_Authorize(t *testing.T) {
	client := sdk.NewDefaultClient("api.mgr.xwjd.xingchenga.xyz")
	client.WithCredential(credential)
	securityApi := NewOpenSecurityApi(client)
	authorizeTokenRequest := &AuthorizeTokenRequest{
		Token: "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ4d2pkLnhpbmdjaGVuZ2EueHl6Iiwic3ViIjoiMjAzIiwiaWF0IjoxNzA0NzAyODIyLCJleHAiOjE3MDUzMDc2MjIsInR5cGUiOjQyLCJ1c2VyTmFtZSI6IjE1NjY2NjY2NjYwIiwicmVnaW9uIjowLCJtb2JpbGUiOiIxNTYqKioqNjY2MCIsImF2YXRhciI6IiIsInNleCI6MCwibmFtZSI6IiVFNSVCRSU5MCVFNyVCRiU5NCJ9.kgFnjqFhl-EN76uhTzQB5sjgELdsUlaXBu3Ot4tIeE1EcSzTKTP70nFaSjJ75Ah5XeUfASiN7TpnqMOwh-EszQ",
	}
	authorizeUserResponse, err := securityApi.Authorize(authorizeTokenRequest)
	if err != nil {
		t.Fatalf("调用失败: %s", err.Error())
	}
	t.Logf("返回结果: %v", authorizeUserResponse)
}
