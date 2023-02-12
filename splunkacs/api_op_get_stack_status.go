package splunkacs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// The result of getting the stack status
type StackStatusResponse struct {
	StackStatus
}

func (c *SplunkAcsClient) GetStackStatus() (*StackStatusResponse, *SplunkACSResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/adminconfig/v2/status", c.Url), nil)
	if err != nil {
		return nil, nil, err
	}

	apiRes, err := c.doRequest(NewSplunkACSRequest(httpReq))
	if err != nil {
		return nil, apiRes, err
	}

	// fmt.Printf("Status response: %d, %s", apiRes.StatusCode, apiRes.Body)

	if apiRes.StatusCode == http.StatusNotFound {
		return nil, apiRes, fmt.Errorf("Index not found. body: '%s'", apiRes.Body)
	}

	if apiRes.StatusCode != http.StatusOK {
		log.Printf("failed to unmarshal response body: %s", string(apiRes.Body))
		return nil, apiRes, fmt.Errorf("unexpected response while getting index. status: %d, body: %s", apiRes.StatusCode, apiRes.Body)
	}

	result := StackStatusResponse{}
	err = json.Unmarshal(apiRes.Body, &result)
	if err != nil {
		return nil, apiRes, err
	}

	return &result, apiRes, nil
}
