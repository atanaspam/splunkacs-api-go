package splunkacs

import (
	"io"
	"net/http"
)

const BaseURL = "https://admin.splunk.com/"

type SplunkAcsClient struct {
	Url        string
	Token      string
	HttpClient *http.Client
}

func NewClient(deploymentName string, token string) (*SplunkAcsClient, error) {
	return &SplunkAcsClient{
		Url:        "https://admin.splunk.com/" + deploymentName,
		Token:      token,
		HttpClient: &http.Client{},
	}, nil
}

func (c *SplunkAcsClient) doRequest(req *http.Request) ([]byte, *http.Response, error) {
	token := c.Token

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, res, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, res, err
	}

	return body, res, err
}
