# CapacityKind

Capacity kind determines what operations are allowed on a capacity.

- `Market`: User-created capacities. - `Originating`: Provider capacities for selling compute. Cannot add compute   (buy orders/procurements). - `ReadOnly`: System-managed capacities used for legacy compute, bare metal   contracts, and other. Cannot be modified through the API.

## Example Usage

```go
import (
	"github.com/sfcompute/sfc-go/models/components"
)

value := components.CapacityKindMarket

// Open enum: custom values can be created with a direct type cast
custom := components.CapacityKind("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `CapacityKindMarket`      | market                    |
| `CapacityKindOriginating` | originating               |
| `CapacityKindReadOnly`    | read_only                 |