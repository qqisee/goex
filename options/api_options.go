package options

type ApiOptions struct {
	Key        string
	Secret     string
	Passphrase string
	ClientId   string
	TestNet    bool
}

type ApiOption func(options *ApiOptions)

func WithApiKey(key string) ApiOption {
	return func(options *ApiOptions) {
		options.Key = key
	}
}

func WithApiSecretKey(secret string) ApiOption {
	return func(options *ApiOptions) {
		options.Secret = secret
	}
}

func WithPassphrase(passphrase string) ApiOption {
	return func(options *ApiOptions) {
		options.Passphrase = passphrase
	}
}

func WithClientId(clientId string) ApiOption {
	return func(options *ApiOptions) {
		options.ClientId = clientId
	}
}
func WithTestNet(testNet bool) ApiOption {
	//添加okx测试网选项
	return func(options *ApiOptions) {
		options.TestNet = testNet
	}
}
