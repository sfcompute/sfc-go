# V2OrderPrincipalKind

Whether an order's `created_by` principal is a human user or an API token.

## Example Usage

```go
import (
	"github.com/sfcompute/sfc-go/models/components"
)

value := components.V2OrderPrincipalKindUser

// Open enum: custom values can be created with a direct type cast
custom := components.V2OrderPrincipalKind("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `V2OrderPrincipalKindUser`  | user                        |
| `V2OrderPrincipalKindToken` | token                       |