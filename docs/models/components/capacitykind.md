# CapacityKind

Capacity kind determines what operations are allowed on a capacity.
- `Market`: User-created capacities. Fully manageable. - `Originating`: Provider capacities for selling compute. Cannot receive buy orders or be deleted. - `ReadOnly`: System-managed capacities. Cannot be modified through the API.

## Example Usage

```go
import (
	"github.com/sfcompute/sfc-go-sdk/models/components"
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