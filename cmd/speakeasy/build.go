package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"

	"github.com/speakeasy-api/parser/apipackage"
	"github.com/speakeasy-api/parser/services/parser"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

type SpeakeasyConfig struct {
	Name string
	Spec struct {
		Version string
		Schemas struct {
			Type    string
			Version string
			Output  string
		}
	}
	Root string
}

const configOutputRegexp = "(yaml|json)"

var buildFlags = []cli.Flag{
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
	&cli.StringFlag{
		Name:    configFileFlag,
		Aliases: []string{"c"},
		Value:   speakeasyConfigFileName,
		Usage:   "Yaml file to load speakeasy configuration from",
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

func readConfig(fileName string) ([]SpeakeasyConfig, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return []SpeakeasyConfig{}, err
	}

	reader := bytes.NewReader(content)
	decoder := yaml.NewDecoder(reader)

	var configs []SpeakeasyConfig
	config := SpeakeasyConfig{}
	for decoder.Decode(&config) == nil {
		configs = append(configs, config)
	}
	fmt.Printf("configs: %d\n\n", len(configs))
	return configs, nil
}

func buildAction(c *cli.Context) error {
	strategy := c.String(propertyStrategyFlag)

	switch strategy {
	case parser.CamelCase, parser.SnakeCase, parser.PascalCase:
	default:
		return fmt.Errorf("not supported %s propertyStrategy", strategy)
	}

	var configs, err = readConfig(c.String(configFileFlag))
	if err != nil {
		return err
	}

	if len(configs) == 0 {
		return fmt.Errorf("no valid configurations found")
	}

	for _, config := range configs {
		r, err := regexp.Compile(configOutputRegexp)
		if err != nil {
			return err
		}

		outputTypes := r.FindAllString(config.Spec.Schemas.Output, -1)
		if len(outputTypes) == 0 {
			return fmt.Errorf("no valid output types specified")
		}

		err = apipackage.New().Build(&apipackage.Config{
			SearchDir:          c.String(searchDirFlag),
			Excludes:           c.String(excludeFlag),
			MainAPIFile:        config.Root,
			PropNamingStrategy: strategy,
			// TODO: use the -o flag for the output dir and add property "OutputFile" to
			// parser config to determine the output-file name.
			OutputDir:           config.Name,
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
		if err != nil {
			return err
		}
	}
	return nil
}
