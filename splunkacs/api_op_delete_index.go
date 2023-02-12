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
	httpReq, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/adminconfig/v2/indexes/%s", c.Url, indexName), nil)
	if err != nil {
		return nil, nil, err
	}

	apiRes, err := c.doRequest(NewSplunkApiRequest(httpReq))
	if err != nil {
		return nil, apiRes.HttpResponse, err
	}

	if apiRes.StatusCode == http.StatusNotFound {
		return nil, apiRes.HttpResponse, fmt.Errorf("index not found. body: '%s'", apiRes.Body)
	}

	if apiRes.StatusCode != http.StatusAccepted {
		return nil, apiRes.HttpResponse, fmt.Errorf("unexpected response while deleting index. status: %d, body: %s", apiRes.StatusCode, apiRes.Body)
	}

	result := IndexDeleteResponse{}
	result.Body = string(apiRes.Body)

	return &result, apiRes.HttpResponse, nil
}
