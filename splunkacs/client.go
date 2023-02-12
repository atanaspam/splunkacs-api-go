package splunkacs

import (
	"fmt"
	"log"
	"net/http"
	"time"
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

func (c *SplunkAcsClient) executeHttpRequest(req *http.Request) (*http.Response, error) {
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *SplunkAcsClient) doRequest(apiRequest *SplunkApiRequest) (*SplunkApiResponse, error) {

	token := c.Token
	apiRequest.HttpRequest.Header.Set("Authorization", "Bearer "+token)
	apiRequest.HttpRequest.Header.Set("Content-Type", "application/json")

	for i := 0; i < apiRequest.RetryLimit; i++ {
		res, err := c.executeHttpRequest(apiRequest.HttpRequest)
		if err != nil {
			return nil, err
		}
		/*
			A throttle conditon occurs when:
			According to the docs: (https://docs.splunk.com/Documentation/SplunkCloud/9.0.2209/Config/ACSerrormessages)
				* Response contains message "403-forbidden"
				* Header contains "awselb"
			According to my findings:
				* Response contains "code":"429-too-many-request" in json payload
				* Response code is 429
				* Headers contain Server:[awselb/2.0]

			For simplicity for now I will assume a response code 492 means we are being throttled.
			In the future we might need to add or change this condition.
		*/
		if res.StatusCode == http.StatusTooManyRequests {
			log.Printf("WARNING: Detected throttling during http request. Retry %d \n", i)
			/*
				I have hardcoded the retries to 4 because it gives a reasonable tradeoff between simply waiting 5 or 10 mins
				and handling throttling that might span two rolling windows (2*5 min).
				In the future we might want to refactor this to a standard exponential backoff and make the retries user configurable.
			*/
			switch i {
			case 0:
				time.Sleep(25 * time.Second)
				continue
			case 1:
				time.Sleep(75 * time.Second)
				continue
			case 2:
				time.Sleep(225 * time.Second)
				continue
			case 3:
				time.Sleep(300 * time.Second)
				continue
			default:
				time.Sleep(15 * time.Second)
				continue
			}
		}
		return NewSplunkApiResponse(res)
	}
	return nil, fmt.Errorf("failed to get a valid response after %d retries", apiRequest.RetryLimit)
}
