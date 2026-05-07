# Toolsets

## Overview

Managed toolsets for gram AI consumers.

### Available Operations

* [GetBySlug](#getbyslug) - getToolset toolsets
* [List](#list) - listToolsets toolsets

## GetBySlug

Get detailed information about a toolset including full HTTP tool definitions

### Example Usage

<!-- UsageSnippet language="go" operationID="getToolset" method="get" path="/rpc/toolsets.get" -->
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

    res, err := s.Toolsets.GetBySlug(ctx, operations.GetToolsetSecurity{
        Option1: &operations.GetToolsetSecurityOption1{
            ProjectSlugHeaderGramProject: "<YOUR_API_KEY_HERE>",
            SessionHeaderGramSession: "<YOUR_API_KEY_HERE>",
        },
    }, "<value>", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Toolset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `security`                                                                         | [operations.GetToolsetSecurity](../../pkg/models/operations/gettoolsetsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |
| `slug`                                                                             | `string`                                                                           | :heavy_check_mark:                                                                 | The slug of the toolset                                                            |
| `gramSession`                                                                      | `*string`                                                                          | :heavy_minus_sign:                                                                 | Session header                                                                     |
| `gramKey`                                                                          | `*string`                                                                          | :heavy_minus_sign:                                                                 | API Key header                                                                     |
| `gramProject`                                                                      | `*string`                                                                          | :heavy_minus_sign:                                                                 | project header                                                                     |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetToolsetResponse](../../pkg/models/operations/gettoolsetresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.Error                   | 400, 401, 403, 404, 409, 415, 422 | application/json                  |
| sdkerrors.Error                   | 500, 502                          | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## List

List all toolsets for a project

### Example Usage

<!-- UsageSnippet language="go" operationID="listToolsets" method="get" path="/rpc/toolsets.list" -->
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

    res, err := s.Toolsets.List(ctx, operations.ListToolsetsSecurity{
        Option1: &operations.ListToolsetsSecurityOption1{
            ProjectSlugHeaderGramProject: "<YOUR_API_KEY_HERE>",
            SessionHeaderGramSession: "<YOUR_API_KEY_HERE>",
        },
    }, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListToolsetsResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `security`                                                                             | [operations.ListToolsetsSecurity](../../pkg/models/operations/listtoolsetssecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |
| `gramSession`                                                                          | `*string`                                                                              | :heavy_minus_sign:                                                                     | Session header                                                                         |
| `gramKey`                                                                              | `*string`                                                                              | :heavy_minus_sign:                                                                     | API Key header                                                                         |
| `gramProject`                                                                          | `*string`                                                                              | :heavy_minus_sign:                                                                     | project header                                                                         |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListToolsetsResponse](../../pkg/models/operations/listtoolsetsresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.Error                   | 400, 401, 403, 404, 409, 415, 422 | application/json                  |
| sdkerrors.Error                   | 500, 502                          | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |