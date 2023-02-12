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

	getStackStatusRes, apiRes, err := acsClient.GetStackStatus()
	if err != nil {
		fmt.Printf("encountered unexpected error. Response code: %d\n", apiRes.StatusCode)
		log.Fatal(err)
	}
	fmt.Printf("status: '%s' \n", getStackStatusRes.StackStatus.Infrastructure.Status)
}
