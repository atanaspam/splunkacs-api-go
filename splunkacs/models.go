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

type StackStatus struct {
	Infrastructure      StackStatusInfrastructure `json:"infrastructure,omitempty"`
	StackStatusMessages StackStatusMessages       `json:"messages,omitempty"`
}

type StackStatusInfrastructure struct {
	StackType    string `json:"stackType,omitempty"`
	StackVersion string `json:"stackVersion,omitempty"`
	Status       string `json:"status,omitempty"`
}

type StackStatusMessages struct {
	RestartRequired bool `json:"restartRequired,omitempty"`
}

type IPWhiteList struct {
	Subnets []string `json:"subnets,omitempty"`
}

type SplunkACSRequest struct {
	HttpRequest *http.Request
	RetryLimit  int
}

func NewSplunkACSRequest(httpRequest *http.Request) *SplunkACSRequest {
	return &SplunkACSRequest{
		HttpRequest: httpRequest,
		RetryLimit:  4, // Arbitrary magic value
	}
}

type SplunkACSResponse struct {
	HttpResponse *http.Response
	Body         []byte
	StatusCode   int
}

func NewSplunkACSResponse(httpResponse *http.Response) (*SplunkACSResponse, error) {
	defer httpResponse.Body.Close()
	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return &SplunkACSResponse{
			HttpResponse: httpResponse,
		}, nil
	}

	return &SplunkACSResponse{
		HttpResponse: httpResponse,
		Body:         body,
		StatusCode:   httpResponse.StatusCode,
	}, nil
}
