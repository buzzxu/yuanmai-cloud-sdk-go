package sdk

type CredentialIface interface {
	GetSecretId() string
	GetToken() string
	GetSecretKey() string
	GetCredential() (string, string, string)
}

type Credential struct {
	SecretId  string
	SecretKey string
	Token     string
}

func (c *Credential) needRefresh() bool {
	return false
}

func (c *Credential) refresh() {
}

func NewCredential(secretId, secretKey string) *Credential {
	return &Credential{
		SecretId:  secretId,
		SecretKey: secretKey,
	}
}

func NewTokenCredential(secretId, secretKey, token string) *Credential {
	return &Credential{
		SecretId:  secretId,
		SecretKey: secretKey,
		Token:     token,
	}
}

func (c *Credential) GetSecretKey() string {
	return c.SecretKey
}

func (c *Credential) GetSecretId() string {
	return c.SecretId
}

func (c *Credential) GetToken() string {
	return c.Token
}

func (c *Credential) GetCredential() (string, string, string) {
	return c.SecretId, c.SecretKey, c.Token
}
