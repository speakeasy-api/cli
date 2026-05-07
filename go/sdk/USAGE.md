<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/speakeasy-api/gf/go/sdk"
	"github.com/speakeasy-api/gf/go/sdk/pkg/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdk.New()

	example, fileErr := os.Open("example.file")
	if fileErr != nil {
		panic(fileErr)
	}

	res, err := s.Assets.UploadFunctions(ctx, operations.UploadFunctionsRequest{
		ContentLength: 858625,
		Body:          example,
	}, operations.UploadFunctionsSecurity{
		Option1: &operations.UploadFunctionsSecurityOption1{
			ApikeyHeaderGramKey:          "<YOUR_API_KEY_HERE>",
			ProjectSlugHeaderGramProject: "<YOUR_API_KEY_HERE>",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.UploadFunctionsResult != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->