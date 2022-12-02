package splunkacs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// The request for updating an individual HEC Token
type HttpEventCollectorUpdateRequest struct {
	HecTokenSpec
}

// The result of updating an individual HEC Token
// Splunk Docs and Splunk API response seem to differ. The struct below represents whats in the docs
// but is actually not correct.
type HttpEventCollectorUpdateResponse struct {
	Code string `json:"code"`
}

func (c *SplunkAcsClient) UpdateHecToken(hecName string, hecUpdateRequest HttpEventCollectorUpdateRequest) (*HttpEventCollectorUpdateResponse, *http.Response, error) {
	rb, err := json.Marshal(hecUpdateRequest)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/adminconfig/v2/inputs/http-event-collectors/%s", c.Url, hecName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, nil, err
	}

	body, res, err := c.doRequest(req)
	if err != nil {
		return nil, res, err
	}

	if res.StatusCode != http.StatusAccepted {
		return nil, res, fmt.Errorf("unexpected response while updating HEC token. status: %d, body: %s", res.StatusCode, body)
	}

	result := HttpEventCollectorUpdateResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("failed to unmarshal response body: %s", string(body))
		return nil, res, err
	}

	return &result, res, nil
}
