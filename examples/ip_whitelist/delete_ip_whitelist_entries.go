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
	var TargetfeatureFlag = flag.String("target_feature", "", "The feature which will be targeted")
	var SubnetToDeleteFlag = flag.String("subnet", "", "The subnet to be added (in CIDR notation). Ex: 123.123.123.123/32")
	flag.Parse()
	acsClient, err := splunkacs.NewClient(*DeploymentNameFlag, os.Getenv("SPLUNK_AUTH_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	ipWhitelistEntryDeleteRequest := splunkacs.IPWhitelistEntriesDeleteRequest{
		TargetFeature:   *TargetfeatureFlag,
		EntriesToDelete: []string{*SubnetToDeleteFlag},
	}

	_, apiRes, err := acsClient.DeleteIPWhitelistEntries(ipWhitelistEntryDeleteRequest)
	if err != nil {
		fmt.Printf("encountered unexpected error. Response code: %d\n", apiRes.StatusCode)
		log.Fatal(err)
	}
	fmt.Printf("status: %d\n", apiRes.StatusCode)
}
