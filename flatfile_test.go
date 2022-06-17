package flatfile_test

import (
	"os"
	"testing"

	"github.com/be-streamlined/flatfile-go"
)

var (
	test_api_key        = "test_api_key"
	test_secret_api_key = "test_secret_key"
	test_license_key    = "test_license_key"
	test_base_url       = "https://bestreamlined.io"
)

func setEnvironmentVariables() {
	os.Setenv("FLATFILE_API_KEY", test_api_key)
	os.Setenv("FLATFILE_SECRET_API_KEY", test_secret_api_key)
	os.Setenv("FLATFILE_LICENSE_KEY", test_license_key)
	os.Setenv("FLATFILE_BASE_URL", test_base_url)
}

func unsetEnvironmentVariables() {
	os.Unsetenv("FLATFILE_API_KEY")
	os.Unsetenv("FLATFILE_SECRET_API_KEY")
	os.Unsetenv("FLATFILE_LICENSE_KEY")
	os.Unsetenv("FLATFILE_BASE_URL")
}

func Test_NewFromEnv(t *testing.T) {
	setEnvironmentVariables()
	defer unsetEnvironmentVariables()

	_, err := flatfile.NewFromEnv()

	if err != nil {
		t.Errorf("wanted success, got error: %e", err)
	}
}

func Test_BaseUrlFromEnvironment(t *testing.T) {
	setEnvironmentVariables()
	defer unsetEnvironmentVariables()

	ff, err := flatfile.NewFromEnv()

	if err != nil {
		t.Errorf("wanted success, got error: %e", err)
	}

	if ff.BaseURL != test_base_url {
		t.Errorf("wanted %s, got %s", test_base_url, ff.BaseURL)
	}
}

func Test_NewFromConfig(t *testing.T) {
	unsetEnvironmentVariables()

	ff, err := flatfile.New(test_api_key, test_secret_api_key, test_license_key)

	if err != nil {
		t.Errorf("wanted success, got error: %e", err)
	}

	if ff.ApiKey != test_api_key {
		t.Errorf("wanted %s, got %s", test_api_key, ff.ApiKey)
	}
}
