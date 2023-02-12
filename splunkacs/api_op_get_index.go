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

func (c *SplunkAcsClient) GetIndex(indexName string) (*IndexGetResponse, *http.Response, error) {
	httpReq, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/adminconfig/v2/indexes/%s", c.Url, indexName), nil)
	if err != nil {
		return nil, nil, err
	}

	// body, res, err := c.doRequest(NewSplunkApiRequest(httpReq))
	// if err != nil {
	// 	return nil, res, err
	// }

	apiResp, err := c.doRequest(NewSplunkApiRequest(httpReq))
	if err != nil {
		return nil, apiResp.HttpResponse, err
	}

	if apiResp.HttpResponse.StatusCode == http.StatusNotFound {
		return nil, apiResp.HttpResponse, fmt.Errorf("Index not found. body: '%s'", apiResp.Body)
	}

	if apiResp.HttpResponse.StatusCode != http.StatusOK {
		log.Printf("failed to unmarshal response body: %s", string(apiResp.Body))
		return nil, apiResp.HttpResponse, fmt.Errorf("unexpected response while getting index. status: %d, body: %s", apiResp.HttpResponse.StatusCode, apiResp.Body)
	}

	result := IndexGetResponse{}
	err = json.Unmarshal(apiResp.Body, &result)
	if err != nil {
		return nil, apiResp.HttpResponse, err
	}

	return &result, apiResp.HttpResponse, nil
}
