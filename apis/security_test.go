package apis

import (
	"github.com/buzzxu/yuanmai-cloud-sdk-go/sdk"
	"testing"
)

var credential = sdk.NewCredential("xwjd", "XXXX")

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
		Key:      "157811",
		ValiCode: "7852",
	}
	info, err := securityApi.Login(loginRequest)
	if err != nil {
		t.Fatalf("调用失败: %s", err.Error())
	}
	t.Logf("返回结果: %v", info)
}
