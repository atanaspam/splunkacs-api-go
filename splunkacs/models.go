package splunkacs

import (
	"io"
	"net/http"
)

// https://docs.splunk.com/Documentation/SplunkCloud/9.0.2208/Config/ManageHECtokens#View_existing_HEC_tokens
type HttpEventCollectorToken struct {
	Spec  HecTokenSpec `json:"spec"`
	Token string       `json:"token,omitempty"`
}

type HecTokenSpec struct {
	AllowedIndexes    []string `json:"allowedIndexes,omitempty"`
	DefaultHost       string   `json:"defaultHost,omitempty"`
	DefaultIndex      string   `json:"defaultIndex,omitempty"`
	DefaultSource     string   `json:"defaultSource,omitempty"`
	DefaultSourcetype string   `json:"defaultSourcetype,omitempty"`
	Disabled          bool     `json:"disabled,omitempty"`
	Name              string   `json:"name,omitempty"`
	UseACK            bool     `json:"useACK,omitempty"`
}

type Index struct {
	Name            string `json:"name,omitempty"`
	DataType        string `json:"datatype,omitempty"`
	SearchableDays  int    `json:"searchableDays,omitempty"`
	MaxDataSizeMb   int    `json:"maxDataSizeMB,omitempty"`
	TotalEventCount string `json:"totalEventCount,omitempty"`
	TotalRawSizeMb  string `json:"totalRawSizeMB,omitempty"`
}

type SplunkApiRequest struct {
	HttpRequest *http.Request
	RetryLimit  int
}

func NewSplunkApiRequest(httpRequest *http.Request) *SplunkApiRequest {
	return &SplunkApiRequest{
		HttpRequest: httpRequest,
		RetryLimit:  4, // Arbitrary magic value
	}
}

type SplunkApiResponse struct {
	HttpResponse *http.Response
	Body         []byte
	StatusCode   int
}

func NewSplunkApiResponse(httpResponse *http.Response) (*SplunkApiResponse, error) {
	defer httpResponse.Body.Close()
	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return &SplunkApiResponse{
			HttpResponse: httpResponse,
		}, nil
	}

	return &SplunkApiResponse{
		HttpResponse: httpResponse,
		Body:         body,
		StatusCode:   httpResponse.StatusCode,
	}, nil
}
