# InstancePriority

Instance priority — a relative ranking that determines which instances the system prefers to keep running when capacity is constrained. When a capacity's quota drops below its running-instance count, instances are terminated in priority order (lower first).

Ordering: `yield < normal < preferred < critical`.

## Example Usage

```go
import (
	"github.com/sfcompute/sfc-go/models/components"
)

value := components.InstancePriorityYield

// Open enum: custom values can be created with a direct type cast
custom := components.InstancePriority("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `InstancePriorityYield`     | yield                       |
| `InstancePriorityNormal`    | normal                      |
| `InstancePriorityPreferred` | preferred                   |
| `InstancePriorityCritical`  | critical                    |