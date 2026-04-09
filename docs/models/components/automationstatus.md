# AutomationStatus

## Example Usage

```go
import (
	"github.com/sfcompute/sfc-go-sdk/models/components"
)

value := components.AutomationStatusInfo

// Open enum: custom values can be created with a direct type cast
custom := components.AutomationStatus("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `AutomationStatusInfo`    | info                      |
| `AutomationStatusWarning` | warning                   |
| `AutomationStatusError`   | error                     |