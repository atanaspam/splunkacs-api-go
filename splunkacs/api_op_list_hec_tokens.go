package splunkacs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// The result of Listing all HEC Tokens
type HttpEventCollectorListResponse struct {
	HttpEventCollectors []HttpEventCollectorToken `json:"http-event-collectors"`
}

// Lists all HECs
// TODO check if count=0 is sufficent or we need to use pagination to go beyond 100 results
func (c *SplunkAcsClient) ListHecTokens() (*HttpEventCollectorListResponse, *http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/adminconfig/v2/inputs/http-event-collectors?count=0", c.Url), nil)
	if err != nil {
		return nil, nil, err
	}

	body, res, err := c.doRequest(req)
	if err != nil {
		return nil, res, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("unexpected response while listing HEC Tokens. status: %d, body: %s", res.StatusCode, body)
	}

	result := HttpEventCollectorListResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, res, err
	}

	return &result, res, nil
}
