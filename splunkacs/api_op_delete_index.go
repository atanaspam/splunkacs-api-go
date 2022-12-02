package splunkacs

import (
	"fmt"
	"net/http"
)

// The result of deleting an index
// The Splunk api does not return a response body even though the documentation claims so
// TODO: figure out what to do here
type IndexDeleteResponse struct {
	Body string
}

func (c *SplunkAcsClient) DeleteIndex(indexName string) (*IndexDeleteResponse, *http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/adminconfig/v2/indexes/%s", c.Url, indexName), nil)
	if err != nil {
		return nil, nil, err
	}

	body, res, err := c.doRequest(req)
	if err != nil {
		return nil, res, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, res, fmt.Errorf("index not found. body: '%s'", body)
	}

	if res.StatusCode != http.StatusAccepted {
		return nil, res, fmt.Errorf("unexpected response while deleting index. status: %d, body: %s", res.StatusCode, body)
	}

	result := IndexDeleteResponse{}
	result.Body = string(body)

	return &result, res, nil
}
