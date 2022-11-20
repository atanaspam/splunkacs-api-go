package splunkacs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// The result of listing multiple indexes
type IndexListResponse []struct {
	Name            *string `json:"name"`
	DataType        *string `json:"datatype"`
	SearchableDays  *int    `json:"searchableDays"`
	MaxDataSizeMb   *int    `json:"maxDataSizeMB"`
	TotalEventCount *string `json:"totalEventCount"`
	TotalRawSizeMb  *string `json:"totalRawSizeMB"`
}

// Lists all Indexes
// TODO check if count=0 is sufficent or we need to use pagination to go beyond 100 results
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
		return nil, res, fmt.Errorf("Unexpected response while listing indexes. status: %d, body: %s", res.StatusCode, body)
	}

	result := IndexListResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, res, err
	}

	return &result, res, nil
}
