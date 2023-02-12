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
func (c *SplunkAcsClient) ListIndexes() (*IndexListResponse, *SplunkACSResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/adminconfig/v2/indexes?count=0", c.Url), nil)
	if err != nil {
		return nil, nil, err
	}

	apiRes, err := c.doRequest(NewSplunkACSRequest(httpReq))
	if err != nil {
		return nil, apiRes, err
	}

	if apiRes.StatusCode != http.StatusOK {
		return nil, apiRes, fmt.Errorf("unexpected response while listing indexes. status: %d, body: %s", apiRes.StatusCode, apiRes.Body)
	}

	result := IndexListResponse{}
	err = json.Unmarshal(apiRes.Body, &result)
	if err != nil {
		return nil, apiRes, err
	}

	return &result, apiRes, nil
}
