package splunkacs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// The result of listing ip whitelists
type IPWhitelistGetResponse struct {
	IPWhiteList
}

// Lists all IP whitelists given a specific feature
func (c *SplunkAcsClient) GetIPWhitelist(feature string) (*IPWhitelistGetResponse, *SplunkACSResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/adminconfig/v2/access/%s/ipallowlists", c.Url, feature), nil)
	if err != nil {
		return nil, nil, err
	}

	apiRes, err := c.doRequest(NewSplunkACSRequest(httpReq))
	if err != nil {
		return nil, apiRes, err
	}

	if apiRes.StatusCode != http.StatusOK {
		return nil, apiRes, fmt.Errorf("unexpected response while getting ip whitelist. status: %d, body: %s", apiRes.StatusCode, apiRes.Body)
	}

	result := IPWhitelistGetResponse{}
	err = json.Unmarshal(apiRes.Body, &result)
	if err != nil {
		return nil, apiRes, err
	}

	return &result, apiRes, nil
}
