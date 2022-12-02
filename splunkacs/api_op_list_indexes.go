package splunkacs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// The result of listing multiple indexes
type IndexListResponse []Index

// Lists all Indexes
// TODO introduce pagination
func (c *SplunkAcsClient) ListIndexes() (*IndexListResponse, *http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/adminconfig/v2/indexes?count=0", c.Url), nil)
	if err != nil {
		return nil, nil, err
	}

	body, res, err := c.doRequest(req)
	if err != nil {
		return nil, res, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("unexpected response while listing indexes. status: %d, body: %s", res.StatusCode, body)
	}

	result := IndexListResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, res, err
	}

	return &result, res, nil
}
