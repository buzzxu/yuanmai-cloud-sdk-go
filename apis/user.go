package apis

import "github.com/buzzxu/yuanmai-cloud-sdk-go/sdk"

type (
	UserApi interface {
		// Info 用户信息
		Info(request *GetUserInfoRequest) (map[string]interface{}, error)
		// FindById 查询用户信息
		FindById(request *GetUserInfoRequest) ([]map[string]interface{}, error)
		// FindRoleMenu 查询角色菜单
		FindRoleMenu(request *FindRoleMenuRequest) ([]map[string]interface{}, error)
		// PasswordChange 修改密码
		PasswordChange(request *PasswordChangeRequest) (string, error)
		// AvatarUpload 头像上传
		AvatarUpload(request *AvatarUploadRequest) (bool, error)
	}

	OpenUserApi struct {
		BaseOpenApi
	}

	GetUserInfoRequest struct {
		UserId int64 `json:"userId"`
	}

	// FindRoleMenuRequest 查询角色菜单
	FindRoleMenuRequest struct {
		RoleId int    `json:"roleId"`
		Region string `json:"region"`
	}

	// PasswordChangeRequest 修改密码
	PasswordChangeRequest struct {
		UserId      int64  `json:"userId"`
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}
	AvatarUploadRequest struct {
		UserId int64 `json:"userId"`
		// base64编码
		Image string `json:"image"`
	}
)

// NewOpenUserApi 实例化OpenUserApi
func NewOpenUserApi(client *sdk.Client) *OpenUserApi {
	return &OpenUserApi{
		BaseOpenApi{
			client:  client,
			service: "user",
		},
	}
}

// Info 用户信息
func (t *OpenUserApi) Info(request *GetUserInfoRequest) (map[string]interface{}, error) {
	response := &sdk.BaseResponse[map[string]interface{}]{}
	err := t.client.Send(t.service, "info", request, response)
	if err != nil {
		return nil, err
	}
	return response.Data, err
}

// FindById 用户信息
func (t *OpenUserApi) FindById(request *GetUserInfoRequest) (map[string]interface{}, error) {
	response := &sdk.BaseResponse[map[string]interface{}]{}
	err := t.client.Send(t.service, "findById", request, response)
	if err != nil {
		return nil, err
	}
	return response.Data, err
}

// FindRoleMenu 查询角色菜单
func (t *OpenUserApi) FindRoleMenu(request *FindRoleMenuRequest) ([]map[string]interface{}, error) {
	response := &sdk.BaseResponse[[]map[string]interface{}]{}
	err := t.client.Send(t.service, "findRoleMenu", request, response)
	if err != nil {
		return nil, err
	}
	return response.Data, err
}

// PasswordChange 修改密码
func (t *OpenUserApi) PasswordChange(request *PasswordChangeRequest) (string, error) {
	response := &sdk.BaseResponse[string]{}
	err := t.client.Send(t.service, "password/change", request, response)
	if err != nil {
		return "", err
	}
	return response.Data, err
}

// AvatarUpload 头像上传
func (t *OpenUserApi) AvatarUpload(request *AvatarUploadRequest) (bool, error) {
	response := &sdk.BaseResponse[bool]{}
	err := t.client.Send(t.service, "upload/avatar", request, response)
	if err != nil {
		return false, err
	}
	return response.Data, err
}
