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
	httpReq, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/adminconfig/v2/inputs/http-event-collectors?count=0", c.Url), nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := c.doRequest(NewSplunkApiRequest(httpReq))
	if err != nil {
		return nil, res.HttpResponse, err
	}

	if res.HttpResponse.StatusCode != http.StatusOK {
		return nil, res.HttpResponse, fmt.Errorf("unexpected response while listing HEC Tokens. status: %d, body: %s", res.HttpResponse.StatusCode, res.Body)
	}

	result := HttpEventCollectorListResponse{}
	err = json.Unmarshal(res.Body, &result)
	if err != nil {
		return nil, res.HttpResponse, err
	}

	return &result, res.HttpResponse, nil
}
