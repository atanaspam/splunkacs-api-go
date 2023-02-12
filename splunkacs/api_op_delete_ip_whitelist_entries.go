package splunkacs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// The request for removing new entries in an IP whitelist
type IPWhitelistEntriesDeleteRequest struct {
	TargetFeature   string
	EntriesToDelete []string
}

// The response for creating new entries in an IP whitelist
type IPWhitelistEntriesDeleteResponse struct {
	Body string
}

func (c *SplunkAcsClient) DeleteIPWhitelistEntries(ipWhitelistEntriesCreateRequest IPWhitelistEntriesDeleteRequest) (*IPWhitelistEntriesDeleteResponse, *SplunkACSResponse, error) {
	reqBody, err := json.Marshal(IPWhiteList{Subnets: ipWhitelistEntriesCreateRequest.EntriesToDelete})
	if err != nil {
		return nil, nil, err
	}

	httpReq, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/adminconfig/v2/access/%s/ipallowlists", c.Url, ipWhitelistEntriesCreateRequest.TargetFeature), strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, nil, err
	}

	apiRes, err := c.doRequest(NewSplunkACSRequest(httpReq))
	if err != nil {
		return nil, apiRes, err
	}

	if apiRes.StatusCode != http.StatusOK {
		return nil, apiRes, fmt.Errorf("unexpected response while deleting IP whitelist entries. status: %d, body: %s", apiRes.StatusCode, apiRes.Body)
	}

	result := IPWhitelistEntriesDeleteResponse{Body: string(apiRes.Body)}

	return &result, apiRes, nil
}
