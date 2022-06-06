package main

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/speakeasy-api/parser/services/parser"
	"github.com/urfave/cli/v2"
)

const testOutputDirectory = ".testOutput"

func TestInit(t *testing.T) {
	defer func() {
		// Clean up temporary output directory.
		err := os.RemoveAll(testOutputDirectory)
		if err != nil {
			panic(err)
		}
	}()

	var tests = []struct {
		strategy, outputTypes string
		expected              string
	}{
		{parser.CamelCase, "json,yaml", ""},
		{parser.SnakeCase, "json,yaml", ""},
		{parser.PascalCase, "json,yaml", ""},
		{parser.PascalCase, "json, yaml", ""},
		{"invalidStrategy", "json,yaml", "not supported invalidStrategy propertyStrategy"},
		{parser.CamelCase, "", "no output types specified"},
	}

	for _, test := range tests {

		// Create temporary output directory.
		if _, err := os.Stat(testOutputDirectory); os.IsNotExist(err) {
			err := os.Mkdir(testOutputDirectory, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}

		t.Run(fmt.Sprintf("%s, %s", test.strategy, test.outputTypes), func(t *testing.T) {
			set := flag.NewFlagSet("test", 0)
			set.String(propertyStrategyFlag, test.strategy, "strategy")
			set.String(outputTypesFlag, test.outputTypes, "outputTypes")
			set.String(searchDirFlag, "test_fixtures", "search")
			set.String(generalInfoFlag, "fixture.go", "generalInfo")
			set.String(outputFlag, testOutputDirectory, "output")

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
