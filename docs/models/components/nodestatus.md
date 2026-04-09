# NodeStatus

`awaiting_allocation` when waiting for compute allocation on its capacity, `running` once assigned and physical node is running (still takes time for image to be downloaded and booted), `terminated` when stopped by user or after running out of allocation, `failed` on hardware fault.

## Example Usage

```go
import (
	"github.com/sfcompute/sfc-go-sdk/models/components"
)

value := components.NodeStatusAwaitingAllocation

// Open enum: custom values can be created with a direct type cast
custom := components.NodeStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `NodeStatusAwaitingAllocation` | awaiting_allocation            |
| `NodeStatusRunning`            | running                        |
| `NodeStatusTerminated`         | terminated                     |
| `NodeStatusFailed`             | failed                         |