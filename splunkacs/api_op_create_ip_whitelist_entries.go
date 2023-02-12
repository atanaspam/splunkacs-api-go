package splunkacs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// The request for creating new entries in an IP whitelist
type IPWhitelistEntriesCreateRequest struct {
	TargetFeature string
	EntriesToAdd  []string
}

// The response for creating new entries in an IP whitelist.
type IPWhitelistEntriesCreateResponse struct {
	Body string
}

func (c *SplunkAcsClient) CreateIPWhitelistEntries(ipWhitelistEntriesCreateRequest IPWhitelistEntriesCreateRequest) (*IPWhitelistEntriesCreateResponse, *SplunkACSResponse, error) {
	reqBody, err := json.Marshal(IPWhiteList{Subnets: ipWhitelistEntriesCreateRequest.EntriesToAdd})
	if err != nil {
		return nil, nil, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/adminconfig/v2/access/%s/ipallowlists", c.Url, ipWhitelistEntriesCreateRequest.TargetFeature), strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, nil, err
	}

	apiRes, err := c.doRequest(NewSplunkACSRequest(httpReq))
	if err != nil {
		return nil, apiRes, err
	}

	if apiRes.StatusCode != http.StatusOK {
		return nil, apiRes, fmt.Errorf("unexpected response while creating IP whitelist entries. status: %d, body: %s", apiRes.StatusCode, apiRes.Body)
	}

	// The remote API does not return any body so we return an empty struct
	result := IPWhitelistEntriesCreateResponse{Body: string(apiRes.Body)}

	return &result, apiRes, nil
}
