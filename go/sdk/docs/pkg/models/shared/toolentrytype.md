# ToolEntryType

## Example Usage

```go
import (
	"github.com/speakeasy-api/cli/go/sdk/pkg/models/shared"
)

value := shared.ToolEntryTypeHTTP

// Open enum: custom values can be created with a direct type cast
custom := shared.ToolEntryType("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ToolEntryTypeHTTP`        | http                       |
| `ToolEntryTypePrompt`      | prompt                     |
| `ToolEntryTypeFunction`    | function                   |
| `ToolEntryTypePlatform`    | platform                   |
| `ToolEntryTypeExternalmcp` | externalmcp                |