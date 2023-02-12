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
	reqBody, err := json.Marshal(hecUpdateRequest)
	if err != nil {
		return nil, nil, err
	}

	httpReq, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/adminconfig/v2/inputs/http-event-collectors/%s", c.Url, hecName), strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, nil, err
	}

	apiRes, err := c.doRequest(NewSplunkApiRequest(httpReq))
	if err != nil {
		return nil, apiRes.HttpResponse, err
	}

	if apiRes.StatusCode != http.StatusAccepted {
		return nil, apiRes.HttpResponse, fmt.Errorf("unexpected response while updating HEC token. status: %d, body: %s", apiRes.StatusCode, apiRes.Body)
	}

	result := HttpEventCollectorUpdateResponse{}
	err = json.Unmarshal(apiRes.Body, &result)
	if err != nil {
		log.Printf("failed to unmarshal response body: %s", string(apiRes.Body))
		return nil, apiRes.HttpResponse, err
	}

	return &result, apiRes.HttpResponse, nil
}
