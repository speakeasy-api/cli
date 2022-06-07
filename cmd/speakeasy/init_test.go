package main

import (
	"flag"
	"fmt"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestBuildConfigStrings(t *testing.T) {
	var tests = []struct {
		name, generalInfo string
		expected          []string
	}{
		{"api_name", "main.go", []string{"name: api_name", "\troot: main.go"}},
		{"other_api", "controller.go", []string{"name: other_api", "\troot: controller.go"}},
	}
	expectedLength := 4

	for _, test := range tests {

		t.Run(fmt.Sprintf("%s, %s", test.name, test.generalInfo), func(t *testing.T) {
			set := flag.NewFlagSet("test", 0)
			set.String(apiNameFlag, test.name, "name")
			set.String(generalInfoFlag, test.generalInfo, "generalInfo")

			actual := buildConfigStrings(cli.NewContext(nil, set, nil))

			if len(actual) != expectedLength {
				t.Errorf("Receieved %d strings, expected %d", len(actual), expectedLength)
			}
			if actual[0] != test.expected[0] {
				t.Errorf("Received %s, expected %s", actual[0], test.expected[0])
			}
			if actual[3] != test.expected[1] {
				t.Errorf("Received %s, expected %s", actual[3], test.expected[1])
			}
		})
	}
}

func TestBuildActionStrings(t *testing.T) {
	expectedCount := 8
	t.Run("action string count", func(t *testing.T) {
		actual := buildActionStrings()

		if len(actual) != expectedCount {
			t.Errorf("Received %d strings, expected %d", len(actual), expectedCount)
		}
	})
}
