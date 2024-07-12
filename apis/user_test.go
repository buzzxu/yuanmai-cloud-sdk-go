package apis

import (
	"github.com/buzzxu/yuanmai-cloud-sdk-go/sdk"
	"testing"
)

func TestOpenUserApi_Info(t *testing.T) {
	client := sdk.NewDefaultClient("xwjd-mgr.xingchenga.xyz")
	client.WithCredential(credential)
	securityApi := NewOpenUserApi(client)
	getUserInfoRequest := &GetUserInfoRequest{

		UserId: 247,
	}
	trees, err := securityApi.Info(getUserInfoRequest)
	if err != nil {
		t.Fatalf("调用失败: %s", err.Error())
	}
	t.Logf("返回结果: %v", trees)
}

func TestOpenUserApi_FindRoleMenu(t *testing.T) {
	client := sdk.NewDefaultClient("api.mgr.xwjd.xingchenga.xyz")
	client.WithCredential(credential)
	securityApi := NewOpenUserApi(client)
	findRoleMenuRequest := &FindRoleMenuRequest{

		RoleId: 42,
		Region: "",
	}
	trees, err := securityApi.FindRoleMenu(findRoleMenuRequest)
	if err != nil {
		t.Fatalf("调用失败: %s", err.Error())
	}
	t.Logf("返回结果: %v", trees)
}
