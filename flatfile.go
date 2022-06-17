// Copyright © 2022 Streamlined Holdings, Inc <devrel@bestreamlined.io>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Flatfile-go is a Golang package that wraps the Flatfile APIs
package flatfile

func New(apiKey, secretApiKey, licenseKey string) (*FlatfileConfig, error) {
	config := FlatfileConfig{
		ApiKey:       apiKey,
		SecretApiKey: secretApiKey,
		LicenseKey:   licenseKey,
	}

	config.FillDefaults()

	err := config.VerifyEnvironment()

	return &config, err
}

func NewFromEnv() (*FlatfileConfig, error) {
	config := FlatfileConfig{}

	config.FillDefaults()

	err := config.VerifyEnvironment()

	return &config, err
}
