package main

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/speakeasy-api/parser/services/parser"
	"github.com/urfave/cli/v2"
)

const testOutputDirectory = ".buildTestOutput"

func TestBuild(t *testing.T) {
	defer func() {
		// Clean up temporary output directories.
		err := os.RemoveAll("schemas")
		if err != nil {
			panic(err)
		}
	}()

	validConfigFile := "test_fixtures/speakeasy.yaml"
	invalidConfigFile := "test_fixtures/invalidConfigFile.yaml"

	tests := []struct {
		configFile, strategy string
		expected             string
	}{
		{validConfigFile, parser.CamelCase, ""},
		{validConfigFile, parser.SnakeCase, ""},
		{validConfigFile, parser.PascalCase, ""},
		{validConfigFile, parser.PascalCase, ""},
		{validConfigFile, "invalidStrategy", "not supported invalidStrategy propertyStrategy"},
		{invalidConfigFile, parser.PascalCase, fmt.Sprintf("open %s: no such file or directory", invalidConfigFile)},
		{"test_fixtures/speakeasy_with_invalid_output.yaml", parser.PascalCase, "no valid output types specified"},
	}

	for _, test := range tests {

		// Create temporary output directory.
		if _, err := os.Stat(testOutputDirectory); os.IsNotExist(err) {
			err := os.Mkdir(testOutputDirectory, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}

		t.Run(fmt.Sprintf("%s, %s", test.configFile, test.strategy), func(t *testing.T) {
			set := flag.NewFlagSet("test", 0)
			set.String(configFileFlag, test.configFile, "yaml")
			set.String(propertyStrategyFlag, test.strategy, "strategy")
			set.String(searchDirFlag, "test_fixtures", "search")
			set.String(generalInfoFlag, "fixture.go", "generalInfo")

			actual := buildAction(cli.NewContext(nil, set, nil))

			if actual == nil {
				if test.expected != "" {
					t.Errorf("Got nil, expected error with message '%s'", test.expected)
				}
			} else if actual.Error() != test.expected {
				t.Errorf("Got error with message %s, expected message %s", actual.Error(), test.expected)
			}
		})
	}
}
