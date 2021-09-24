package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Azure/go-autorest/autorest/azure/cli"
)

type DeploymentTypeValue struct {
	Type string `json:"type"`
	Value interface{} `json:"value"`
}

type DeploymentProperties struct {
	Outputs map[string]DeploymentTypeValue `json:"outputs"`
	Parameters map[string]DeploymentTypeValue `json:"parameters"`
}

type Deployment struct {
	Name 		  string `json:"name"`
	ResourceGroup string `json:"resourceGroup"`
	Properties    DeploymentProperties `json:"properties"`
}

func main() {
	if (len(os.Args) != 2) {
		fmt.Println("Please specify a resource group")
		os.Exit(1)
	}

	resourceGroup := os.Args[1]

	cliCmd := cli.GetAzureCLICommand()

	cliCmd.Args = append(cliCmd.Args, "deployment", "group", "list", "-g", resourceGroup)

	output, err := cliCmd.Output()

	if (err != nil) {
		fmt.Println(err)
	}

	var deploymentsResponse []Deployment

	err = json.Unmarshal(output, &deploymentsResponse)

	if (err != nil) {
		fmt.Println(err)
	}

	aioDeployment := getAIODeployment(deploymentsResponse)

	if aioDeployment != nil {
		fmt.Println(prettyPrint(aioDeployment))
	} else {
		fmt.Println("Failed to find any AIO deployment")
	}
}

func getAIODeployment(deployments []Deployment) *Deployment {
	for _, deployment := range deployments {
		if (strings.HasPrefix(deployment.Name, "AIODeploy")) {
			return &deployment
		}
	}

	return nil
}

func prettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "  ")
    return string(s)
}
