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
	httpReq, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/adminconfig/v2/indexes?count=0", c.Url), nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := c.doRequest(NewSplunkApiRequest(httpReq))
	if err != nil {
		return nil, res.HttpResponse, err
	}

	if res.HttpResponse.StatusCode != http.StatusOK {
		return nil, res.HttpResponse, fmt.Errorf("unexpected response while listing indexes. status: %d, body: %s", res.HttpResponse.StatusCode, res.Body)
	}

	result := IndexListResponse{}
	err = json.Unmarshal(res.Body, &result)
	if err != nil {
		return nil, res.HttpResponse, err
	}

	return &result, res.HttpResponse, nil
}
