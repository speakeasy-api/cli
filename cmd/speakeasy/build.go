package main

import (
	"fmt"
	"strings"

	"github.com/speakeasy-api/parser/apipackage"
	"github.com/speakeasy-api/parser/services/parser"
	"github.com/urfave/cli/v2"
)

var buildFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    generalInfoFlag,
		Aliases: []string{"g"},
		Value:   "main.go",
		Usage:   "Go file path in which 'OpenAPI general API Info' is written",
	},
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
		Name:    propertyStrategyFlag,
		Aliases: []string{"p"},
		Value:   parser.CamelCase,
		Usage:   "Property Naming Strategy like " + parser.SnakeCase + "," + parser.CamelCase + "," + parser.PascalCase,
	},
	&cli.StringFlag{
		Name:    outputFlag,
		Aliases: []string{"o"},
		Value:   "./docs",
		Usage:   "Output directory for all the generated files(opeanapi.json, opeanapi.yaml)",
	},
	&cli.StringFlag{
		Name:    outputTypesFlag,
		Aliases: []string{"ot"},
		Value:   "json,yaml",
		Usage:   "Output types of generated files (opeanapi.json, opeanapi.yaml) like go,json,yaml",
	},
	&cli.BoolFlag{
		Name:  parseVendorFlag,
		Usage: "Parse go files in 'vendor' folder, disabled by default",
	},
	&cli.BoolFlag{
		Name:    parseDependencyFlag,
		Aliases: []string{"pd"},
		Usage:   "Parse go files inside dependency folder, disabled by default",
	},
	&cli.StringFlag{
		Name:    markdownFilesFlag,
		Aliases: []string{"md"},
		Value:   "",
		Usage:   "Parse folder containing markdown files to use as description, disabled by default",
	},
	&cli.StringFlag{
		Name:    codeExampleFilesFlag,
		Aliases: []string{"cef"},
		Value:   "",
		Usage:   "Parse folder containing code example files to use for the x-codeSamples extension, disabled by default",
	},
	&cli.BoolFlag{
		Name:  parseInternalFlag,
		Usage: "Parse go files in internal packages, disabled by default",
	},
	&cli.BoolFlag{
		Name:  generatedTimeFlag,
		Usage: "Generate timestamp at the top of docs.go, disabled by default",
	},
	&cli.IntFlag{
		Name:  parseDepthFlag,
		Value: 100,
		Usage: "Dependency parse depth",
	},
	&cli.StringFlag{
		Name:  instanceNameFlag,
		Value: "",
		Usage: "This parameter can be used to name different schema(openapi) document instances. It is optional.",
	},
	&cli.StringFlag{
		Name:  overridesFileFlag,
		Value: apipackage.DefaultOverridesFile,
		Usage: "File to read global type overrides from.",
	},
}

func buildAction(c *cli.Context) error {
	strategy := c.String(propertyStrategyFlag)

	switch strategy {
	case parser.CamelCase, parser.SnakeCase, parser.PascalCase:
	default:
		return fmt.Errorf("not supported %s propertyStrategy", strategy)
	}

	// apipackage library handles trimming white-space.
	outputTypes := strings.Split(c.String(outputTypesFlag), ",")
	if len(outputTypes) == 1 && len(outputTypes[0]) == 0 {
		return fmt.Errorf("no output types specified")
	}

	return apipackage.New().Build(&apipackage.Config{
		SearchDir:           c.String(searchDirFlag),
		Excludes:            c.String(excludeFlag),
		MainAPIFile:         c.String(generalInfoFlag),
		PropNamingStrategy:  strategy,
		OutputDir:           c.String(outputFlag),
		OutputTypes:         outputTypes,
		ParseVendor:         c.Bool(parseVendorFlag),
		ParseDependency:     c.Bool(parseDependencyFlag),
		MarkdownFilesDir:    c.String(markdownFilesFlag),
		ParseInternal:       c.Bool(parseInternalFlag),
		GeneratedTime:       c.Bool(generatedTimeFlag),
		CodeExampleFilesDir: c.String(codeExampleFilesFlag),
		ParseDepth:          c.Int(parseDepthFlag),
		InstanceName:        c.String(instanceNameFlag),
		OverridesFile:       c.String(overridesFileFlag),
	})
}
