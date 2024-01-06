package apis

import "github.com/buzzxu/yuanmai-cloud-sdk-go/sdk"

type (
	UserApi interface {
		// Info 用户信息
		Info() (map[string]interface{}, error)
		// FindRoleMenu 查询角色菜单
		FindRoleMenu(request *FindRoleMenuRequest) ([]map[string]interface{}, error)
	}

	OpenUserApi struct {
		client  *sdk.Client
		Service string
	}

	GetUserInfoRequest struct {
		UserId int64 `json:"userId"`
	}

	// FindRoleMenuRequest 查询角色菜单
	FindRoleMenuRequest struct {
		RoleId int    `json:"roleId"`
		Region string `json:"region"`
	}
)

// Info 用户信息
func (t *OpenUserApi) Info(request *GetUserInfoRequest) (map[string]interface{}, error) {
	response := &sdk.BaseResponse[map[string]interface{}]{}
	err := t.client.Send(t.Service, "info", request, response)
	if err != nil {
		return nil, err
	}
	return response.Data, err
}

// FindRoleMenu 查询角色菜单
func (t *OpenUserApi) FindRoleMenu(request *FindRoleMenuRequest) ([]map[string]interface{}, error) {
	response := &sdk.BaseResponse[[]map[string]interface{}]{}
	err := t.client.Send(t.Service, "findRoleMenu", request, response)
	if err != nil {
		return nil, err
	}
	return response.Data, err
}
