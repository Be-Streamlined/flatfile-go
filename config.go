package flatfile

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

type FlatfileConfig struct {
	ApiKey       string
	SecretApiKey string
	LicenseKey   string
	BaseURL      string
	HttpClient   *http.Client
}

func (f *FlatfileConfig) FillDefaults() {
	if f.ApiKey == "" {
		f.ApiKey = getEnv("FLATFILE_API_KEY", "")
	}

	if f.SecretApiKey == "" {
		f.SecretApiKey = getEnv("FLATFILE_SECRET_API_KEY", "")
	}

	if f.LicenseKey == "" {
		f.LicenseKey = getEnv("FLATFILE_LICENSE_KEY", "")
	}

	if f.BaseURL == "" {
		f.BaseURL = getEnv("FLATFILE_BASE_URL", "https://api.us.flatfile.io/")
	}

	if f.HttpClient == nil {
		f.HttpClient = http.DefaultClient
	}
}

func (f *FlatfileConfig) buildURL(p string, query map[string]string) (string, error) {
	parsedURL, err := url.Parse(f.BaseURL)
	if err != nil {
		return "", err
	}

	q := parsedURL.Query()

	for key, value := range query {
		q.Add(key, value)
	}

	parsedURL.Path = path.Join(parsedURL.Path, p)

	parsedURL.RawQuery = q.Encode()

	return parsedURL.String(), nil
}

func (f *FlatfileConfig) buildRequest(m, p string, b io.Reader, q map[string]string) (*http.Request, error) {
	if q == nil {
		q = map[string]string{
			"licenseKey": f.LicenseKey,
		}
	} else {
		q["licenseKey"] = f.LicenseKey
	}

	u, err := f.buildURL(p, q)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(m, u, b)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Api-Key", fmt.Sprintf("%s+%s", f.ApiKey, f.SecretApiKey))

	return req, nil
}
