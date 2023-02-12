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

	if apiRes.StatusCode != http.StatusOK {
		log.Printf("failed to get stack status: %s", string(apiRes.Body))
		return nil, apiRes, fmt.Errorf("unexpected response while getting stack status. code: %d, body: %s", apiRes.StatusCode, apiRes.Body)
	}

	result := StackStatusResponse{}
	err = json.Unmarshal(apiRes.Body, &result)
	if err != nil {
		return nil, apiRes, err
	}

	return &result, apiRes, nil
}
