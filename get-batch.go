package flatfile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/be-streamlined/flatfile-go/flatfile_types"
)

func (c *FlatfileConfig) GetBatch(id string, offset int) (*flatfile_types.BatchResponse, error) {
	urlPath := fmt.Sprintf("/rest/batch/%s", id)

	q := map[string]string{
		"skip": fmt.Sprintf("%d", offset),
	}

	req, err := c.buildRequest("GET", urlPath, nil, q)

	if err != nil {
		return nil, err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 299 || resp.StatusCode < 200 {
		return nil, fmt.Errorf("invalid request: %s", resp.Status)
	}

	responseBody := flatfile_types.BatchResponse{}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &responseBody); err != nil {
		return nil, err
	}

	return &responseBody, nil
}
