package splunkacs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// The request for updating an individual index
type IndexUpdateRequest struct {
	SearchableDays int `json:"searchableDays,omitempty"`
	MaxDataSizeMb  int `json:"maxDataSizeMB,omitempty"`
}

// The result of updating an individual index
type IndexUpdateResponse struct {
	Index
}

func (c *SplunkAcsClient) UpdateIndex(indexName string, indexUpdateRequest IndexUpdateRequest) (*IndexUpdateResponse, *http.Response, error) {
	rb, err := json.Marshal(indexUpdateRequest)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/adminconfig/v2/indexes/%s", c.Url, indexName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, nil, err
	}

	body, res, err := c.doRequest(req)
	if err != nil {
		return nil, res, err
	}

	fmt.Println(res.StatusCode)
	if res.StatusCode != http.StatusAccepted {
		return nil, res, fmt.Errorf("unexpected response while updating index. status: %d, body: %s", res.StatusCode, body)
	}

	result := IndexUpdateResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("failed to unmarshal response body: %s", string(body))
		return nil, res, err
	}

	return &result, res, nil
}
