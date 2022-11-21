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
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/adminconfig/v2/inputs/http-event-collectors/%s", c.Url, hecName), nil)
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

	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("unexpected response while getting HEC Token. status: %d, body: %s", res.StatusCode, body)
	}

	result := HttpEventCollectorGetResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, res, err
	}

	return &result, res, nil
}
