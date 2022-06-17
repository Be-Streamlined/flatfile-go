package flatfile

func New(apiKey, secretApiKey, licenseKey string) *FlatfileConfig {
	config := FlatfileConfig{
		ApiKey:       apiKey,
		SecretApiKey: secretApiKey,
		LicenseKey:   licenseKey,
	}

	config.FillDefaults()

	return &config
}
