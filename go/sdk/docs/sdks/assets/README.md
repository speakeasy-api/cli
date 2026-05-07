# Assets

## Overview

Manages assets used by Gram projects.

### Available Operations

* [UploadFunctions](#uploadfunctions) - uploadFunctions assets
* [UploadOpenAPIv3](#uploadopenapiv3) - uploadOpenAPIv3 assets

## UploadFunctions

Upload functions to Gram.

### Example Usage

<!-- UsageSnippet language="go" operationID="uploadFunctions" method="post" path="/rpc/assets.uploadFunctions" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/gf/go/sdk"
	"os"
	"github.com/speakeasy-api/gf/go/sdk/pkg/models/operations"
	"log"
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
        Body: example,
    }, operations.UploadFunctionsSecurity{
        Option1: &operations.UploadFunctionsSecurityOption1{
            ApikeyHeaderGramKey: "<YOUR_API_KEY_HERE>",
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

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UploadFunctionsRequest](../../pkg/models/operations/uploadfunctionsrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.UploadFunctionsSecurity](../../pkg/models/operations/uploadfunctionssecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UploadFunctionsResponse](../../pkg/models/operations/uploadfunctionsresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.Error                   | 400, 401, 403, 404, 409, 415, 422 | application/json                  |
| sdkerrors.Error                   | 500, 502                          | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## UploadOpenAPIv3

Upload an OpenAPI v3 document to Gram.

### Example Usage

<!-- UsageSnippet language="go" operationID="uploadOpenAPIv3Asset" method="post" path="/rpc/assets.uploadOpenAPIv3" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/gf/go/sdk"
	"os"
	"github.com/speakeasy-api/gf/go/sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdk.New()

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.Assets.UploadOpenAPIv3(ctx, operations.UploadOpenAPIv3AssetRequest{
        ContentLength: 513080,
        Body: example,
    }, operations.UploadOpenAPIv3AssetSecurity{
        Option1: &operations.UploadOpenAPIv3AssetSecurityOption1{
            ApikeyHeaderGramKey: "<YOUR_API_KEY_HERE>",
            ProjectSlugHeaderGramProject: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UploadOpenAPIv3Result != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UploadOpenAPIv3AssetRequest](../../pkg/models/operations/uploadopenapiv3assetrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.UploadOpenAPIv3AssetSecurity](../../pkg/models/operations/uploadopenapiv3assetsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UploadOpenAPIv3AssetResponse](../../pkg/models/operations/uploadopenapiv3assetresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.Error                   | 400, 401, 403, 404, 409, 415, 422 | application/json                  |
| sdkerrors.Error                   | 500, 502                          | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |