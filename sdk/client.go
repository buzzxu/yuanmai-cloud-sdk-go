package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/buzzxu/boys/common/strs"
	"github.com/buzzxu/go-openapi"
	"github.com/buzzxu/yuanmai-cloud-sdk-go/common"
	"io"
	"net/http"
)

type Client struct {
	Scheme string
	Domain string
	Path   string
	// 凭证
	credential CredentialIface
}

func NewDefaultClient(domain string) *Client {
	return &Client{
		Scheme: "https",
		Domain: domain,
		Path:   "open",
	}
}
func (c *Client) Init() *Client {
	return c
}

func (c *Client) WithScheme(scheme string) *Client {
	c.Scheme = scheme
	return c
}

func (c *Client) WithPath(path string) *Client {
	c.Path = path
	return c
}

func (c *Client) WithSecretId(secretId, secretKey string) *Client {
	c.credential = NewCredential(secretId, secretKey)
	return c
}

func (c *Client) WithCredential(cred CredentialIface) *Client {
	c.credential = cred
	return c
}

func (c *Client) GetCredential() CredentialIface {
	return c.credential
}

func (c *Client) Send(service, method string, params interface{}, resp Response) error {
	return c.sendWithSignature(fmt.Sprintf("%s://%s/%s/%s/%s", c.Scheme, strs.Trim(c.Domain), strs.Trim(c.Path), service, method), params, resp)
}

func (c *Client) sendWithSignature(url string, params interface{}, resp Response) (err error) {
	err = openapi.Call(url, c.credential.GetSecretId(), c.credential.GetSecretKey(), params, nil, func(response *http.Response) error {
		if response.StatusCode == 200 {
			bytes, err := io.ReadAll(response.Body)
			if err != nil {
				return common.NewYuanmaiCloudSdkError(500, "")
			}
			err = json.Unmarshal(bytes, resp)
			if resp.error() {
				return common.NewYuanmaiCloudSdkError(resp.code(), resp.message())
			}

			return err
		}
		return common.NewYuanmaiCloudSdkError(response.StatusCode, response.Status)
	})
	if err != nil {
		return common.NewYuanmaiCloudSdkError(500, err.Error())
	}
	return nil
}
