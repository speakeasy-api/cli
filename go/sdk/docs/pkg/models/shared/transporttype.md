# TransportType

The transport type used to connect to the MCP server

## Example Usage

```go
import (
	"github.com/speakeasy-api/cli/go/sdk/pkg/models/shared"
)

value := shared.TransportTypeStreamableHTTP

// Open enum: custom values can be created with a direct type cast
custom := shared.TransportType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `TransportTypeStreamableHTTP` | streamable-http               |
| `TransportTypeSse`            | sse                           |