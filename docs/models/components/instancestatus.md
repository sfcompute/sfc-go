# InstanceStatus

`awaiting_allocation` when waiting for compute allocation on its capacity, `running` once assigned and the physical machine is running (still takes time for the image to be downloaded and booted), `terminated` when stopped by the user or after running out of allocation, `failed` on hardware fault.

## Example Usage

```go
import (
	"github.com/sfcompute/sfc-go/models/components"
)

value := components.InstanceStatusAwaitingAllocation

// Open enum: custom values can be created with a direct type cast
custom := components.InstanceStatus("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `InstanceStatusAwaitingAllocation` | awaiting_allocation                |
| `InstanceStatusRunning`            | running                            |
| `InstanceStatusTerminated`         | terminated                         |
| `InstanceStatusFailed`             | failed                             |