package main

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"

	"github.com/speakeasy-api/parser/services/parser"
)

const (
	speakeasyConfigFileName = "speakeasy.yaml"

	// CLI flags
	configFileFlag       = "config"
	searchDirFlag        = "dir"
	excludeFlag          = "exclude"
	generalInfoFlag      = "generalInfo"
	propertyStrategyFlag = "propertyStrategy"
	outputFlag           = "output"
	outputTypesFlag      = "outputTypes"
	parseVendorFlag      = "parseVendor"
	parseDependencyFlag  = "parseDependency"
	markdownFilesFlag    = "markdownFiles"
	codeExampleFilesFlag = "codeExampleFiles"
	parseInternalFlag    = "parseInternal"
	generatedTimeFlag    = "generatedTime"
	parseDepthFlag       = "parseDepth"
	instanceNameFlag     = "instanceName"
	overridesFileFlag    = "overridesFile"
	apiNameFlag          = "apiNameFlag"
)

func exitNormally(diffEmpty bool) {
	if false && !diffEmpty {
		os.Exit(1)
	}
	os.Exit(0)
}

func printYAML(output interface{}) error {
	if reflect.ValueOf(output).IsNil() {
		return nil
	}

	bytes, err := yaml.Marshal(output)
	if err != nil {
		return err
	}
	fmt.Printf("%s", bytes)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Version = "v0.1.1-alpha-rc5"
	app.Usage = "Automatically track the state of your API, Generate artifacts like OpenAPI schemas and more."
	app.Commands = []*cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Create initial state",
			Action:  initAction,
			Flags:   initFlags,
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "Create schema artifacts",
			Action:  buildAction,
			Flags:   buildFlags,
		},
		{
			Name:    "fmt",
			Aliases: []string{"f"},
			Usage:   "format Annotations",
			Action: func(c *cli.Context) error {
				searchDir := c.String(searchDirFlag)
				excludeDir := c.String(excludeFlag)
				mainFile := c.String(generalInfoFlag)

				return parser.NewFormat().Build(&parser.Config{
					SearchDir: searchDir,
					Excludes:  excludeDir,
					MainFile:  mainFile,
				})
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    searchDirFlag,
					Aliases: []string{"d"},
					Value:   "./",
					Usage:   "Directories you want to parse,comma separated and general-info file must be in the first one",
				},
				&cli.StringFlag{
					Name:  excludeFlag,
					Usage: "Exclude directories and files when searching, comma separated",
				},
				&cli.StringFlag{
					Name:    generalInfoFlag,
					Aliases: []string{"g"},
					Value:   "main.go",
					Usage:   "Go file path in which 'openapi general API Info' is written",
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
