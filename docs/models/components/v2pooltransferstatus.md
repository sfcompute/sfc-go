# V2PoolTransferStatus

Lifecycle status of a pool transfer.

## Example Usage

```go
import (
	"github.com/sfcompute/sfc-go/models/components"
)

value := components.V2PoolTransferStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.V2PoolTransferStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `V2PoolTransferStatusPending`  | pending                        |
| `V2PoolTransferStatusExecuted` | executed                       |
| `V2PoolTransferStatusRejected` | rejected                       |