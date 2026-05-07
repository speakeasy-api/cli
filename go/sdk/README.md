# Go SDK

Placeholder for the OAS-generated Go SDK used by `go/cli`.

This directory should stay generated-client-only. Do not import Gram server packages here.


<!-- Start Summary [summary] -->
## Summary

Gram API Description: Gram is the tools platform for AI agents
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [Go SDK](#go-sdk)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/speakeasy-api/gf/go/sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name                           | Type   | Scheme  |
| ------------------------------ | ------ | ------- |
| `ApikeyHeaderGramKey`          | apiKey | API key |
| `ProjectSlugHeaderGramProject` | apiKey | API key |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:


### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
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
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Assets](docs/sdks/assets/README.md)

* [UploadFunctions](docs/sdks/assets/README.md#uploadfunctions) - uploadFunctions assets
* [UploadOpenAPIv3](docs/sdks/assets/README.md#uploadopenapiv3) - uploadOpenAPIv3 assets

### [Deployments](docs/sdks/deployments/README.md)

* [Active](docs/sdks/deployments/README.md#active) - getActiveDeployment deployments
* [Create](docs/sdks/deployments/README.md#create) - createDeployment deployments
* [EvolveDeployment](docs/sdks/deployments/README.md#evolvedeployment) - evolve deployments
* [GetByID](docs/sdks/deployments/README.md#getbyid) - getDeployment deployments
* [Latest](docs/sdks/deployments/README.md#latest) - getLatestDeployment deployments
* [RedeployDeployment](docs/sdks/deployments/README.md#redeploydeployment) - redeploy deployments

### [Keys](docs/sdks/keys/README.md)

* [Validate](docs/sdks/keys/README.md#validate) - verifyKey keys

### [Toolsets](docs/sdks/toolsets/README.md)

* [GetBySlug](docs/sdks/toolsets/README.md#getbyslug) - getToolset toolsets
* [List](docs/sdks/toolsets/README.md#list) - listToolsets toolsets

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	"github.com/speakeasy-api/gf/go/sdk"
	"github.com/speakeasy-api/gf/go/sdk/pkg/models/operations"
	"github.com/speakeasy-api/gf/go/sdk/pkg/retry"
	"log"
	"os"
	"pkg/models/operations"
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
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.UploadFunctionsResult != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	"github.com/speakeasy-api/gf/go/sdk"
	"github.com/speakeasy-api/gf/go/sdk/pkg/models/operations"
	"github.com/speakeasy-api/gf/go/sdk/pkg/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdk.New(
		sdk.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
	)

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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `UploadFunctions` function may return the following errors:

| Error Type         | Status Code                       | Content Type     |
| ------------------ | --------------------------------- | ---------------- |
| sdkerrors.Error    | 400, 401, 403, 404, 409, 415, 422 | application/json |
| sdkerrors.Error    | 500, 502                          | application/json |
| sdkerrors.APIError | 4XX, 5XX                          | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	"github.com/speakeasy-api/gf/go/sdk"
	"github.com/speakeasy-api/gf/go/sdk/pkg/models/operations"
	"github.com/speakeasy-api/gf/go/sdk/pkg/models/sdkerrors"
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

		var e *sdkerrors.Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
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

	s := sdk.New(
		sdk.WithServerURL("https://app.getgram.ai"),
	)

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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/speakeasy-api/gf/go/sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->
