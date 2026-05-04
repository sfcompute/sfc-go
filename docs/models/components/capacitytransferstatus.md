# CapacityTransferStatus

## Example Usage

```go
import (
	"github.com/sfcompute/sfc-go/models/components"
)

value := components.CapacityTransferStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.CapacityTransferStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `CapacityTransferStatusPending`  | pending                          |
| `CapacityTransferStatusExecuted` | executed                         |
| `CapacityTransferStatusRejected` | rejected                         |