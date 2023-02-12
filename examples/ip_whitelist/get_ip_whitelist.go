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
	var TargetFeatureFlag = flag.String("target_feature", "", "The feature you want to request IP Whitelists for. See https://docs.splunk.com/Documentation/SplunkCloud/9.0.2208/Config/ConfigureIPAllowList#Determine_IP_allow_list_use_case")

	flag.Parse()
	acsClient, err := splunkacs.NewClient(*DeploymentNameFlag, os.Getenv("SPLUNK_AUTH_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	getIpWhitelistResp, apiRes, err := acsClient.GetIPWhitelist(*TargetFeatureFlag)
	if err != nil {
		fmt.Printf("encountered unexpected error. Response code: %d\n", apiRes.StatusCode)
		log.Fatal(err)
	}
	for _, subnet := range getIpWhitelistResp.Subnets {
		fmt.Printf("subnet: '%s'\n", subnet)
	}
}
