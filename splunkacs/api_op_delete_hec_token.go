package splunkacs

import (
	"fmt"
	"net/http"
)

// The result of deleting a HEC Token
type HttpEventCollectorDeleteResponse struct {
	Body string
}

func (c *SplunkAcsClient) DeleteHecToken(hecName string) (*HttpEventCollectorDeleteResponse, *http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/adminconfig/v2/inputs/http-event-collectors/%s", c.Url, hecName), nil)
	if err != nil {
		return nil, nil, err
	}

	body, res, err := c.doRequest(req)
	if err != nil {
		return nil, res, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, res, fmt.Errorf("HEC not found. body: '%s'", body)
	}

	if res.StatusCode != http.StatusAccepted {
		return nil, res, fmt.Errorf("unexpected response while deleting HEC Token. status: %d, body: %s", res.StatusCode, body)
	}

	result := HttpEventCollectorDeleteResponse{}
	result.Body = string(body)

	return &result, res, nil
}
