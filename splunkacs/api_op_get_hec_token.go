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

	apiRes, err := c.doRequest(NewSplunkApiRequest(httpReq))
	if err != nil {
		return nil, apiRes.HttpResponse, err
	}

	if apiRes.StatusCode == http.StatusNotFound {
		return nil, apiRes.HttpResponse, fmt.Errorf("HEC not found. body: '%s'", apiRes.Body)
	}

	if apiRes.StatusCode != http.StatusOK {
		return nil, apiRes.HttpResponse, fmt.Errorf("unexpected response while getting HEC Token. status: %d, body: %s", apiRes.StatusCode, apiRes.Body)
	}

	result := HttpEventCollectorGetResponse{}
	err = json.Unmarshal(apiRes.Body, &result)
	if err != nil {
		return nil, apiRes.HttpResponse, err
	}

	return &result, apiRes.HttpResponse, nil
}
