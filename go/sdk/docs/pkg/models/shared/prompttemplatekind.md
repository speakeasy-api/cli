# PromptTemplateKind

The kind of prompt the template is used for

## Example Usage

```go
import (
	"github.com/speakeasy-api/gf/go/sdk/pkg/models/shared"
)

value := shared.PromptTemplateKindPrompt

// Open enum: custom values can be created with a direct type cast
custom := shared.PromptTemplateKind("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `PromptTemplateKindPrompt`          | prompt                              |
| `PromptTemplateKindHigherOrderTool` | higher_order_tool                   |