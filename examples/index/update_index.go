package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/atanaspam/splunkacs-api-go/splunkacs"
)

func main() {
	var DeploymentNameFlag = flag.String("deployment_name", "", "The Splunk deployment name")
	var IndexNameFlag = flag.String("index_name", "", "The name of the index to get")
	flag.Parse()
	acsClient, err := splunkacs.NewClient(*DeploymentNameFlag, os.Getenv("SPLUNK_AUTH_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	indexUpdateRequest := splunkacs.IndexUpdateRequest{
		SearchableDays: 35,
		MaxDataSizeMb:  0,
	}

	getIndexResp, res, err := acsClient.UpdateIndex(*IndexNameFlag, indexUpdateRequest)
	if err != nil {
		fmt.Printf("encountered unexpected error. Response code: %d\n", res.StatusCode)
		log.Fatal(err)
	}
	fmt.Printf("name: '%s' totalEventCount: '%s, searchableDays: '%s'\n", getIndexResp.Name, getIndexResp.TotalEventCount, getIndexResp.SearchableDays)
}
