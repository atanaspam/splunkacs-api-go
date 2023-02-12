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
	reqBody, err := json.Marshal(indexRequest)
	if err != nil {
		return nil, nil, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/adminconfig/v2/indexes", c.Url), strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, nil, err
	}

	apiRes, err := c.doRequest(NewSplunkApiRequest(httpReq))
	if err != nil {
		return nil, apiRes.HttpResponse, err
	}

	if apiRes.StatusCode != http.StatusAccepted {
		return nil, apiRes.HttpResponse, fmt.Errorf("unexpected response while creating index. status: %d, body: %s", apiRes.StatusCode, apiRes.Body)
	}

	result := IndexCreateResponse{}
	err = json.Unmarshal(apiRes.Body, &result)
	if err != nil {
		return &result, apiRes.HttpResponse, err
	}

	return &result, apiRes.HttpResponse, nil
}
