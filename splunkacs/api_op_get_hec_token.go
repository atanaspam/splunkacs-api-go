package splunkacs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// The result of getting an individual HEC Token
type HttpEventCollectorGetResponse struct {
	HttpEventCollector HttpEventCollectorToken `json:"http-event-collector"`
}

func (c *SplunkAcsClient) GetHecToken(hecName string) (*HttpEventCollectorGetResponse, *http.Response, error) {
	httpReq, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/adminconfig/v2/inputs/http-event-collectors/%s", c.Url, hecName), nil)
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

	if res.HttpResponse.StatusCode != http.StatusOK {
		return nil, res.HttpResponse, fmt.Errorf("unexpected response while getting HEC Token. status: %d, body: %s", res.HttpResponse.StatusCode, res.Body)
	}

	result := HttpEventCollectorGetResponse{}
	err = json.Unmarshal(res.Body, &result)
	if err != nil {
		return nil, res.HttpResponse, err
	}

	return &result, res.HttpResponse, nil
}
