package splunkacs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// The request for creating an index
type IndexCreateRequest struct {
	Name           string `json:"name,omitempty"`
	DataType       string `json:"datatype,omitempty"`
	SearchableDays int    `json:"searchableDays,omitempty"`
	MaxDataSizeMb  int    `json:"maxDataSizeMB,omitempty"`
}

// The response for creating an index
type IndexCreateResponse struct {
	Index
}

func (c *SplunkAcsClient) CreateIndex(indexRequest IndexCreateRequest) (*IndexCreateResponse, *http.Response, error) {
	rb, err := json.Marshal(indexRequest)
	if err != nil {
		return nil, nil, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/adminconfig/v2/indexes", c.Url), strings.NewReader(string(rb)))
	if err != nil {
		return nil, nil, err
	}

	res, err := c.doRequest(NewSplunkApiRequest(httpReq))
	if err != nil {
		return nil, res.HttpResponse, err
	}

	if res.HttpResponse.StatusCode != http.StatusAccepted {
		return nil, res.HttpResponse, fmt.Errorf("unexpected response while creating index. status: %d, body: %s", res.HttpResponse.StatusCode, res.Body)
	}

	result := IndexCreateResponse{}
	err = json.Unmarshal(res.Body, &result)
	if err != nil {
		return &result, res.HttpResponse, err
	}

	return &result, res.HttpResponse, nil
}
