package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var initFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    generalInfoFlag,
		Aliases: []string{"g"},
		Value:   "main.go",
		Usage:   "Go file path in which 'OpenAPI general API Info' is written",
	},
	&cli.StringFlag{
		Name:    apiNameFlag,
		Aliases: []string{"n"},
		Value:   "main_api",
		Usage:   "Name of the api",
	},
}

const (
	speakeasyFileName = "speakeasy.yaml"
	actionFileName    = ".github/workflows/speakeasy.yaml"
	apiNameVariable   = "SPEAKEASY_API_NAME"
	apiRootVariable   = "SPEAKEASY_API_ROOT"
)

func writeSliceToFile(stringsToWrite []string, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, s := range stringsToWrite {
		_, err = file.WriteString(fmt.Sprintf("%s\n", s))
		if err != nil {
			return err
		}
	}
	return nil
}

func buildConfigStrings(c *cli.Context) []string {
	nameString := fmt.Sprintf("name: %s", c.String(apiNameFlag))
	versionString := "\tversion: v1"
	openApiString := "\t\tOpenAPI3.0\n\t\tversion: 1.0.0"
	rootString := fmt.Sprintf("\troot: %s", c.String(generalInfoFlag))
	return []string{nameString, versionString, openApiString, rootString}
}

func buildActionStrings() []string {
	// \n adds line of white-space
	nameString := "name: Run Speakeasy CLI\n"
	jobsString := "jobs:\n\tsetup_and_run_speakeasy:\n"
	containerString := "\t\truns-on: ubuntu-latest\n\n\t\tpermissions:\n\t\t\tcontents: 'read'\n\t\t\tid-token: 'write'\n"
	stepsString := "\t\tsteps:"
	checkoutString := "\t\t\t- name: Checkout\n\t\t\t\tuses: actions/checkout@v3\n\t\t\t\twith:\n\t\t\t\t\tref: ${{ github.head_ref }}\n"
	downloadString := "\t\t\t- name: Download Speakeasy\n\t\t\t\tuses: speakeasy-api/speakeasy-github-action\n"

	// This builds and executes the speakeasy build command
	runString := "\t\t- name: Setup and Update API state\n\t\t\trun: |"
	// grep pulls the line, cut removes everything but the desired string, head -1 selects only the first result
	apiNameString := fmt.Sprintf("\t\t\t\texport %s=$(grep -E -i '^name: ([^\n].*)' speakeasy.yaml | cut -d \" \" -f 2 | head -1)", apiNameVariable)
	rootFileString := fmt.Sprintf("\t\t\t\texport %s=$(grep -E -i 'root: ([^\n].*)' speakeasy.yaml | cut -d \" \" -f 3 | head -1)", apiRootVariable)
	commandString := fmt.Sprintf("\t\t\t\tspeakeasy build -n $%s -g $%s -o schemas/$%s\n", apiNameVariable, apiRootVariable, apiNameVariable)

	// The changes should be committed
	commitString := "\t\t- name: Commit API state\n\t\t\trun: git add schemas; git commit -m \"[no ci]Add schema files\"; git push\n"

	return []string{nameString, jobsString, containerString, stepsString, checkoutString, downloadString, runString, apiNameString, rootFileString, commandString, commitString}
}

func initAction(c *cli.Context) error {
	configStrings := buildConfigStrings(c)
	err := writeSliceToFile(configStrings, speakeasyFileName)
	if err != nil {
		return err
	}

	buildActionStrings := buildActionStrings()
	err = writeSliceToFile(buildActionStrings, actionFileName)
	if err != nil {
		return err
	}
	return nil
}
