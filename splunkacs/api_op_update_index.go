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

	httpReq, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/adminconfig/v2/indexes/%s", c.Url, indexName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, nil, err
	}

	res, err := c.doRequest(NewSplunkApiRequest(httpReq))
	if err != nil {
		return nil, res.HttpResponse, err
	}

	if res.HttpResponse.StatusCode != http.StatusAccepted {
		return nil, res.HttpResponse, fmt.Errorf("unexpected response while updating index. status: %d, body: %s", res.HttpResponse.StatusCode, res.Body)
	}

	result := IndexUpdateResponse{}
	err = json.Unmarshal(res.Body, &result)
	if err != nil {
		log.Printf("failed to unmarshal response body: %s", string(res.Body))
		return nil, res.HttpResponse, err
	}

	return &result, res.HttpResponse, nil
}
