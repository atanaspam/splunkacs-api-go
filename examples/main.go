package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/atanaspam/splunkacs-api-go/splunkacs"
)

var DeploymentNameFlag = flag.String("deployment_name", "", "The Splunk deployment name")

func main() {
	flag.Parse()
	acsClient, err := splunkacs.NewClient(*DeploymentNameFlag, os.Getenv("SPLUNK_AUTH_TOKEN"))
	if err != nil {
		log.Fatal(err)
		return
	}
	// getHecs(*acsClient)
	// getHec(*acsClient, "atanas-test")
	// createHec(*acsClient, "atanas-test")
	// updateHec(*acsClient, "atanas-test")
	// deleteHec(*acsClient, "atanas-test")
	// listIndexes(*acsClient)
	// getIndex(*acsClient, "atanas-test")
	noOp(*acsClient)
}

func noOp(acsClient splunkacs.SplunkAcsClient) {
	// do nothing
}

func getHecs(acsClient splunkacs.SplunkAcsClient) {
	hecListResp, res, err := acsClient.ListHecTokens()
	if err != nil {
		fmt.Printf("%d", res.StatusCode)
		log.Fatal(err)
	}
	for _, token := range hecListResp.HttpEventCollectors {
		fmt.Printf("name: '%v' token: '%v'\n", token.Spec.Name, token.Token)
	}
}

func getHec(acsClient splunkacs.SplunkAcsClient, hecName string) {
	hecResp, res, err := acsClient.GetHecToken(hecName)
	if err != nil {
		fmt.Printf("%v\n", res.StatusCode)
		log.Fatal(err)
	}

	fmt.Printf("name: '%s' token: '%s'\n", hecResp.HttpEventCollector.Spec.Name, hecResp.HttpEventCollector.Token)
}

func createHec(acsClient splunkacs.SplunkAcsClient, tokenName string) {
	defaultIndex := "main"
	disabled := false
	hecCreateRequest := new(splunkacs.HttpEventCollectorCreateRequest)
	hecCreateRequest.AllowedIndexes = []string{"main"}
	hecCreateRequest.DefaultHost = ""
	hecCreateRequest.DefaultIndex = defaultIndex
	hecCreateRequest.DefaultSource = ""
	hecCreateRequest.DefaultSourcetype = ""
	hecCreateRequest.Disabled = disabled
	hecCreateRequest.Name = tokenName
	hecCreateRequest.UseACK = false

	hecCreateResp, res, err := acsClient.CreateHecToken(*hecCreateRequest)
	if err != nil {
		fmt.Printf("%v\n", res)
		log.Fatal(err)
	}

	fmt.Printf("result: '%v'\n", hecCreateResp.CreateResponseItem)
}

func updateHec(acsClient splunkacs.SplunkAcsClient, tokenName string) {
	defaultIndex := "main"
	disabled := false
	hecUpdateRequest := new(splunkacs.HttpEventCollectorUpdateRequest)
	hecUpdateRequest.AllowedIndexes = []string{"main"}
	hecUpdateRequest.DefaultHost = ""
	hecUpdateRequest.DefaultIndex = defaultIndex
	hecUpdateRequest.DefaultSource = "test1"
	hecUpdateRequest.DefaultSourcetype = ""
	hecUpdateRequest.Disabled = disabled
	hecUpdateRequest.Name = tokenName
	hecUpdateRequest.UseACK = true

	hecUpdateResp, res, err := acsClient.UpdateHecToken(tokenName, *hecUpdateRequest)
	if err != nil {
		fmt.Printf("%v\n", res)
		log.Fatal(err)
	}

	fmt.Printf("result: '%s'\n", hecUpdateResp)
}

func deleteHec(acsClient splunkacs.SplunkAcsClient, tokenName string) {
	hecDeleteResp, res, err := acsClient.DeleteHecToken(tokenName)
	if err != nil {
		fmt.Printf("%v\n", res)
		log.Fatal(err)
	}
	fmt.Printf("%s\n", hecDeleteResp)
}

func listIndexes(acsClient splunkacs.SplunkAcsClient) {
	listIndexResp, res, err := acsClient.ListIndexes()
	if err != nil {
		fmt.Printf("%d", res.StatusCode)
		log.Fatal(err)
	}
	for _, index := range *listIndexResp {
		fmt.Printf("name: '%s' totalEventCount: '%s'\n", *index.Name, *index.TotalEventCount)
	}
}

func getIndex(acsClient splunkacs.SplunkAcsClient, indexName string) {
	getIndexResp, res, err := acsClient.GetIndex(indexName)
	if err != nil {
		fmt.Printf("%d", res.StatusCode)
		log.Fatal(err)
	}
	fmt.Printf("name: '%s' totalEventCount: '%s'\n", getIndexResp.Name, getIndexResp.TotalEventCount)
}
