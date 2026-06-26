# Value

The RDMA technology available on the node.

## Example Usage

```go
import (
	"github.com/sfcompute/sfc-go/models/components"
)

value := components.ValueInfiniband

// Open enum: custom values can be created with a direct type cast
custom := components.Value("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `ValueInfiniband` | infiniband        |
| `ValueRoce`       | roce              |