# Keys

## Overview

Managing system api keys.

### Available Operations

* [Validate](#validate) - verifyKey keys

## Validate

Verify an api key

### Example Usage

<!-- UsageSnippet language="go" operationID="validateAPIKey" method="get" path="/rpc/keys.verify" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/gf/go/sdk"
	"github.com/speakeasy-api/gf/go/sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdk.New()

    res, err := s.Keys.Validate(ctx, operations.ValidateAPIKeySecurity{
        ApikeyHeaderGramKey: "<YOUR_API_KEY_HERE>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidateKeyResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `security`                                                                                 | [operations.ValidateAPIKeySecurity](../../pkg/models/operations/validateapikeysecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |
| `speakeasyKey`                                                                                  | `*string`                                                                                  | :heavy_minus_sign:                                                                         | API Key header                                                                             |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ValidateAPIKeyResponse](../../pkg/models/operations/validateapikeyresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.Error                   | 400, 401, 403, 404, 409, 415, 422 | application/json                  |
| sdkerrors.Error                   | 500, 502                          | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |