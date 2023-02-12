package splunkacs

import (
	"fmt"
	"net/http"
)

// The result of deleting a HEC Token
type HttpEventCollectorDeleteResponse struct {
	Body string
}

func (c *SplunkAcsClient) DeleteHecToken(hecName string) (*HttpEventCollectorDeleteResponse, *SplunkACSResponse, error) {
	httpReq, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/adminconfig/v2/inputs/http-event-collectors/%s", c.Url, hecName), nil)
	if err != nil {
		return nil, nil, err
	}

	apiRes, err := c.doRequest(NewSplunkACSRequest(httpReq))
	if err != nil {
		return nil, apiRes, err
	}

	if apiRes.StatusCode == http.StatusNotFound {
		return nil, apiRes, fmt.Errorf("HEC not found. body: '%s'", apiRes.Body)
	}

	if apiRes.StatusCode != http.StatusAccepted {
		return nil, apiRes, fmt.Errorf("unexpected response while deleting HEC Token. status: %d, body: %s", apiRes.StatusCode, apiRes.Body)
	}

	result := HttpEventCollectorDeleteResponse{}
	result.Body = string(apiRes.Body)

	return &result, apiRes, nil
}
