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
	httpReq, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/adminconfig/v2/inputs/http-event-collectors/%s", c.Url, hecName), nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := c.doRequest(NewSplunkApiRequest(httpReq))
	if err != nil {
		return nil, res.HttpResponse, err
	}

	if res.HttpResponse.StatusCode == http.StatusNotFound {
		return nil, res.HttpResponse, fmt.Errorf("HEC not found. body: '%s'", res.Body)
	}

	if res.HttpResponse.StatusCode != http.StatusAccepted {
		return nil, res.HttpResponse, fmt.Errorf("unexpected response while deleting HEC Token. status: %d, body: %s", res.HttpResponse.StatusCode, res.Body)
	}

	result := HttpEventCollectorDeleteResponse{}
	result.Body = string(res.Body)

	return &result, res.HttpResponse, nil
}
