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
func (c *SplunkAcsClient) ListHecTokens() (*HttpEventCollectorListResponse, *SplunkACSResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/adminconfig/v2/inputs/http-event-collectors?count=0", c.Url), nil)
	if err != nil {
		return nil, nil, err
	}

	apiRes, err := c.doRequest(NewSplunkACSRequest(httpReq))
	if err != nil {
		return nil, apiRes, err
	}

	if apiRes.StatusCode != http.StatusOK {
		return nil, apiRes, fmt.Errorf("unexpected response while listing HEC Tokens. status: %d, body: %s", apiRes.StatusCode, apiRes.Body)
	}

	result := HttpEventCollectorListResponse{}
	err = json.Unmarshal(apiRes.Body, &result)
	if err != nil {
		return nil, apiRes, err
	}

	return &result, apiRes, nil
}
