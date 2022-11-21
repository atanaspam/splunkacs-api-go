package splunkacs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// The result of getting an individual index
type IndexGetResponse struct {
	Index
}

func (c *SplunkAcsClient) GetIndex(indexName string) (*IndexGetResponse, *http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/adminconfig/v2/indexes/%s", c.Url, indexName), nil)
	if err != nil {
		return nil, nil, err
	}

	body, res, err := c.doRequest(req)
	if err != nil {
		return nil, res, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, res, fmt.Errorf("Index not found. body: '%s'", body)
	}

	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("unexpected response while getting index. status: %d, body: %s", res.StatusCode, body)
	}

	result := IndexGetResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, res, err
	}

	return &result, res, nil
}
