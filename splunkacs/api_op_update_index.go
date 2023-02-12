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
	reqBody, err := json.Marshal(indexUpdateRequest)
	if err != nil {
		return nil, nil, err
	}

	httpReq, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/adminconfig/v2/indexes/%s", c.Url, indexName), strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, nil, err
	}

	apiRes, err := c.doRequest(NewSplunkApiRequest(httpReq))
	if err != nil {
		return nil, apiRes.HttpResponse, err
	}

	if apiRes.StatusCode != http.StatusAccepted {
		return nil, apiRes.HttpResponse, fmt.Errorf("unexpected response while updating index. status: %d, body: %s", apiRes.StatusCode, apiRes.Body)
	}

	result := IndexUpdateResponse{}
	err = json.Unmarshal(apiRes.Body, &result)
	if err != nil {
		log.Printf("failed to unmarshal response body: %s", string(apiRes.Body))
		return nil, apiRes.HttpResponse, err
	}

	return &result, apiRes.HttpResponse, nil
}
