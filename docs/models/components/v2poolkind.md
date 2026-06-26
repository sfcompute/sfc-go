# V2PoolKind

Pool kind determines what operations are allowed on a pool.

- `Market`: User-created pools. - `Originating`: Provider pools for selling compute. Cannot add compute   (buy orders/procurements). - `ReadOnly`: System-managed pools used for legacy compute, bare metal   contracts, and other. Cannot be modified through the API.

## Example Usage

```go
import (
	"github.com/sfcompute/sfc-go/models/components"
)

value := components.V2PoolKindMarket

// Open enum: custom values can be created with a direct type cast
custom := components.V2PoolKind("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `V2PoolKindMarket`      | market                  |
| `V2PoolKindOriginating` | originating             |
| `V2PoolKindReadOnly`    | read_only               |