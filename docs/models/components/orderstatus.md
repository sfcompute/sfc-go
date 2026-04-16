# OrderStatus

The status of an order in the system.
`pending` = not resolved/processed yet.
`filled` = order executed.
`standing` = the order is waiting for a match.
`cancelled` = the order was cancelled either automatically (not a standing order and didn't immediately fill, or current time past `end_at`) or by explicit cancellation.
`rejected` = validation/system error occurred.

## Example Usage

```go
import (
	"github.com/sfcompute/sfc-go/models/components"
)

value := components.OrderStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.OrderStatus("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `OrderStatusPending`   | pending                |
| `OrderStatusFilled`    | filled                 |
| `OrderStatusRejected`  | rejected               |
| `OrderStatusCancelled` | cancelled              |
| `OrderStatusStanding`  | standing               |