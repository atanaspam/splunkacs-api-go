package splunkacs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// The request for creating a HEC Token
type HttpEventCollectorCreateRequest struct {
	HecTokenSpec
}

// The response for creating a HEC Token
type HttpEventCollectorCreateResponse struct {
	CreateResponseItem HttpEventCollectorCreateResponseItem `json:"http-event-collector"`
}

// The  HEC Token spec in the creation response
type HttpEventCollectorCreateResponseItem struct {
	Spec HttpEventCollectorCreateResponseSpec `json:"spec"`
}

// The 'short' HEC Token definition in the creation response
type HttpEventCollectorCreateResponseSpec struct {
	Name string `json:"name"`
}

func (c *SplunkAcsClient) CreateHecToken(hecCreateRequest HttpEventCollectorCreateRequest) (*HttpEventCollectorCreateResponse, *http.Response, error) {
	rb, err := json.Marshal(hecCreateRequest)
	if err != nil {
		return nil, nil, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/adminconfig/v2/inputs/http-event-collectors", c.Url), strings.NewReader(string(rb)))
	if err != nil {
		return nil, nil, err
	}

	res, err := c.doRequest(NewSplunkApiRequest(httpReq))
	if err != nil {
		return nil, res.HttpResponse, err
	}

	if res.HttpResponse.StatusCode != http.StatusAccepted {
		return nil, res.HttpResponse, fmt.Errorf("unexpected response while creating HEC Token. status: %d, body: %s", res.HttpResponse.StatusCode, res.Body)
	}

	result := HttpEventCollectorCreateResponse{}
	err = json.Unmarshal(res.Body, &result)
	if err != nil {
		return &result, res.HttpResponse, err
	}

	return &result, res.HttpResponse, nil
}
