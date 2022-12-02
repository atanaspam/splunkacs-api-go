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

	flag.Parse()
	acsClient, err := splunkacs.NewClient(*DeploymentNameFlag, os.Getenv("SPLUNK_AUTH_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	listIndexResp, res, err := acsClient.ListIndexes()
	if err != nil {
		fmt.Printf("encountered unexpected error. Response code: %d\n", res.StatusCode)
		log.Fatal(err)
	}
	for _, index := range *listIndexResp {
		fmt.Printf("name: '%s' totalEventCount: '%s'\n", index.Name, index.TotalEventCount)
	}
}
