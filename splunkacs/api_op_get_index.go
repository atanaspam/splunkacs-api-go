package splunkacs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// The result of getting an individual index
type IndexGetResponse struct {
	Index
}

func (c *SplunkAcsClient) GetIndex(indexName string) (*IndexGetResponse, *SplunkACSResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/adminconfig/v2/indexes/%s", c.Url, indexName), nil)
	if err != nil {
		return nil, nil, err
	}

	apiRes, err := c.doRequest(NewSplunkACSRequest(httpReq))
	if err != nil {
		return nil, apiRes, err
	}

	if apiRes.StatusCode == http.StatusNotFound {
		return nil, apiRes, fmt.Errorf("Index not found. body: '%s'", apiRes.Body)
	}

	if apiRes.StatusCode != http.StatusOK {
		log.Printf("failed to unmarshal response body: %s", string(apiRes.Body))
		return nil, apiRes, fmt.Errorf("unexpected response while getting index. status: %d, body: %s", apiRes.StatusCode, apiRes.Body)
	}

	result := IndexGetResponse{}
	err = json.Unmarshal(apiRes.Body, &result)
	if err != nil {
		return nil, apiRes, err
	}

	return &result, apiRes, nil
}
