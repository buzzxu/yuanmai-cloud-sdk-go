package apis

import "github.com/buzzxu/yuanmai-cloud-sdk-go/sdk"

type (
	OpenApi interface {
		Client() *sdk.Client
	}
	BaseOpenApi struct {
		client  *sdk.Client
		service string
	}
)

func (t *BaseOpenApi) Client() *sdk.Client {
	return t.client
}
