# Deployments

## Overview

Manages deployments of tools from upstream sources.

### Available Operations

* [Active](#active) - getActiveDeployment deployments
* [Create](#create) - createDeployment deployments
* [EvolveDeployment](#evolvedeployment) - evolve deployments
* [GetByID](#getbyid) - getDeployment deployments
* [Latest](#latest) - getLatestDeployment deployments
* [RedeployDeployment](#redeploydeployment) - redeploy deployments

## Active

Get the active deployment for a project.

### Example Usage

<!-- UsageSnippet language="go" operationID="getActiveDeployment" method="get" path="/rpc/deployments.active" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/cli/go/sdk"
	"github.com/speakeasy-api/cli/go/sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdk.New()

    res, err := s.Deployments.Active(ctx, operations.GetActiveDeploymentSecurity{
        Option1: &operations.GetActiveDeploymentSecurityOption1{
            ApikeyHeaderGramKey: "<YOUR_API_KEY_HERE>",
            ProjectSlugHeaderGramProject: "<YOUR_API_KEY_HERE>",
        },
    }, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetActiveDeploymentResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `security`                                                                                           | [operations.GetActiveDeploymentSecurity](../../pkg/models/operations/getactivedeploymentsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |
| `speakeasyKey`                                                                                            | `*string`                                                                                            | :heavy_minus_sign:                                                                                   | API Key header                                                                                       |
| `speakeasySession`                                                                                        | `*string`                                                                                            | :heavy_minus_sign:                                                                                   | Session header                                                                                       |
| `speakeasyProject`                                                                                        | `*string`                                                                                            | :heavy_minus_sign:                                                                                   | project header                                                                                       |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetActiveDeploymentResponse](../../pkg/models/operations/getactivedeploymentresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.Error                   | 400, 401, 403, 404, 409, 415, 422 | application/json                  |
| sdkerrors.Error                   | 500, 502                          | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Create

Create a deployment to load tool definitions.

### Example Usage

<!-- UsageSnippet language="go" operationID="createDeployment" method="post" path="/rpc/deployments.create" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/cli/go/sdk"
	"github.com/speakeasy-api/cli/go/sdk/pkg/models/shared"
	"github.com/speakeasy-api/cli/go/sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdk.New()

    res, err := s.Deployments.Create(ctx, operations.CreateDeploymentRequest{
        IdempotencyKey: "01jqq0ajmb4qh9eppz48dejr2m",
        Body: shared.CreateDeploymentRequestBody{
            ExternalID: sdk.Pointer("bc5f4a555e933e6861d12edba4c2d87ef6caf8e6"),
            ExternalMcps: []shared.AddExternalMCPForm{
                shared.AddExternalMCPForm{
                    Name: "My Slack Integration",
                    RegistryServerSpecifier: "slack",
                    SelectedRemotes: []string{
                        "https://mcp.example.com/sse",
                    },
                    Slug: "<value>",
                },
            },
            GithubPr: sdk.Pointer("1234"),
            GithubRepo: sdk.Pointer("speakeasyapi/speakeasy"),
            GithubSha: sdk.Pointer("f33e693e9e12552043bc0ec5c37f1b8a9e076161"),
            NonBlocking: sdk.Pointer(false),
        },
    }, operations.CreateDeploymentSecurity{
        Option1: &operations.CreateDeploymentSecurityOption1{
            ApikeyHeaderGramKey: "<YOUR_API_KEY_HERE>",
            ProjectSlugHeaderGramProject: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDeploymentResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateDeploymentRequest](../../pkg/models/operations/createdeploymentrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.CreateDeploymentSecurity](../../pkg/models/operations/createdeploymentsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateDeploymentResponse](../../pkg/models/operations/createdeploymentresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.Error                   | 400, 401, 403, 404, 409, 415, 422 | application/json                  |
| sdkerrors.Error                   | 500, 502                          | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## EvolveDeployment

Create a new deployment with additional or updated tool sources.

### Example Usage

<!-- UsageSnippet language="go" operationID="evolveDeployment" method="post" path="/rpc/deployments.evolve" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/cli/go/sdk"
	"github.com/speakeasy-api/cli/go/sdk/pkg/models/operations"
	"github.com/speakeasy-api/cli/go/sdk/pkg/models/shared"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdk.New()

    res, err := s.Deployments.EvolveDeployment(ctx, operations.EvolveDeploymentSecurity{
        Option1: &operations.EvolveDeploymentSecurityOption1{
            ApikeyHeaderGramKey: "<YOUR_API_KEY_HERE>",
            ProjectSlugHeaderGramProject: "<YOUR_API_KEY_HERE>",
        },
    }, shared.EvolveForm{
        NonBlocking: sdk.Pointer(false),
        UpsertExternalMcps: []shared.AddExternalMCPForm{
            shared.AddExternalMCPForm{
                Name: "My Slack Integration",
                RegistryServerSpecifier: "slack",
                SelectedRemotes: []string{
                    "https://mcp.example.com/sse",
                },
                Slug: "<value>",
            },
        },
    }, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.EvolveResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `security`                                                                                     | [operations.EvolveDeploymentSecurity](../../pkg/models/operations/evolvedeploymentsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `body`                                                                                         | [shared.EvolveForm](../../pkg/models/shared/evolveform.md)                                     | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `speakeasyKey`                                                                                      | `*string`                                                                                      | :heavy_minus_sign:                                                                             | API Key header                                                                                 |
| `speakeasySession`                                                                                  | `*string`                                                                                      | :heavy_minus_sign:                                                                             | Session header                                                                                 |
| `speakeasyProject`                                                                                  | `*string`                                                                                      | :heavy_minus_sign:                                                                             | project header                                                                                 |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.EvolveDeploymentResponse](../../pkg/models/operations/evolvedeploymentresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.Error                   | 400, 401, 403, 404, 409, 415, 422 | application/json                  |
| sdkerrors.Error                   | 500, 502                          | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## GetByID

Get a deployment by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDeployment" method="get" path="/rpc/deployments.get" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/cli/go/sdk"
	"github.com/speakeasy-api/cli/go/sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdk.New()

    res, err := s.Deployments.GetByID(ctx, operations.GetDeploymentSecurity{
        Option1: &operations.GetDeploymentSecurityOption1{
            ApikeyHeaderGramKey: "<YOUR_API_KEY_HERE>",
            ProjectSlugHeaderGramProject: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetDeploymentResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `security`                                                                               | [operations.GetDeploymentSecurity](../../pkg/models/operations/getdeploymentsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |
| `id`                                                                                     | `string`                                                                                 | :heavy_check_mark:                                                                       | The ID of the deployment                                                                 |
| `speakeasyKey`                                                                                | `*string`                                                                                | :heavy_minus_sign:                                                                       | API Key header                                                                           |
| `speakeasySession`                                                                            | `*string`                                                                                | :heavy_minus_sign:                                                                       | Session header                                                                           |
| `speakeasyProject`                                                                            | `*string`                                                                                | :heavy_minus_sign:                                                                       | project header                                                                           |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetDeploymentResponse](../../pkg/models/operations/getdeploymentresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.Error                   | 400, 401, 403, 404, 409, 415, 422 | application/json                  |
| sdkerrors.Error                   | 500, 502                          | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Latest

Get the latest deployment for a project.

### Example Usage

<!-- UsageSnippet language="go" operationID="getLatestDeployment" method="get" path="/rpc/deployments.latest" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/cli/go/sdk"
	"github.com/speakeasy-api/cli/go/sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdk.New()

    res, err := s.Deployments.Latest(ctx, operations.GetLatestDeploymentSecurity{
        Option1: &operations.GetLatestDeploymentSecurityOption1{
            ApikeyHeaderGramKey: "<YOUR_API_KEY_HERE>",
            ProjectSlugHeaderGramProject: "<YOUR_API_KEY_HERE>",
        },
    }, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetLatestDeploymentResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `security`                                                                                           | [operations.GetLatestDeploymentSecurity](../../pkg/models/operations/getlatestdeploymentsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |
| `speakeasyKey`                                                                                            | `*string`                                                                                            | :heavy_minus_sign:                                                                                   | API Key header                                                                                       |
| `speakeasySession`                                                                                        | `*string`                                                                                            | :heavy_minus_sign:                                                                                   | Session header                                                                                       |
| `speakeasyProject`                                                                                        | `*string`                                                                                            | :heavy_minus_sign:                                                                                   | project header                                                                                       |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetLatestDeploymentResponse](../../pkg/models/operations/getlatestdeploymentresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.Error                   | 400, 401, 403, 404, 409, 415, 422 | application/json                  |
| sdkerrors.Error                   | 500, 502                          | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## RedeployDeployment

Redeploys an existing deployment.

### Example Usage

<!-- UsageSnippet language="go" operationID="redeployDeployment" method="post" path="/rpc/deployments.redeploy" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/cli/go/sdk"
	"github.com/speakeasy-api/cli/go/sdk/pkg/models/operations"
	"github.com/speakeasy-api/cli/go/sdk/pkg/models/shared"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdk.New()

    res, err := s.Deployments.RedeployDeployment(ctx, operations.RedeployDeploymentSecurity{
        Option1: &operations.RedeployDeploymentSecurityOption1{
            ApikeyHeaderGramKey: "<YOUR_API_KEY_HERE>",
            ProjectSlugHeaderGramProject: "<YOUR_API_KEY_HERE>",
        },
    }, shared.RedeployRequestBody{
        DeploymentID: "<id>",
    }, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RedeployResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `security`                                                                                         | [operations.RedeployDeploymentSecurity](../../pkg/models/operations/redeploydeploymentsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |
| `body`                                                                                             | [shared.RedeployRequestBody](../../pkg/models/shared/redeployrequestbody.md)                       | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `speakeasyKey`                                                                                          | `*string`                                                                                          | :heavy_minus_sign:                                                                                 | API Key header                                                                                     |
| `speakeasySession`                                                                                      | `*string`                                                                                          | :heavy_minus_sign:                                                                                 | Session header                                                                                     |
| `speakeasyProject`                                                                                      | `*string`                                                                                          | :heavy_minus_sign:                                                                                 | project header                                                                                     |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RedeployDeploymentResponse](../../pkg/models/operations/redeploydeploymentresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.Error                   | 400, 401, 403, 404, 409, 415, 422 | application/json                  |
| sdkerrors.Error                   | 500, 502                          | application/json                  |
| sdkerrors.APIError                | 4XX, 5XX                          | \*/\*                             |