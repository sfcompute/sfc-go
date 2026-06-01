# V2OrderEstimateNotice


## Supported Types

### V2OrderEstimateNoticeMaintenanceWindow

```go
v2OrderEstimateNotice := components.CreateV2OrderEstimateNoticeV2OrderEstimateNoticeMaintenanceWindow(components.V2OrderEstimateNoticeMaintenanceWindow{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2OrderEstimateNotice.Type {
	case components.V2OrderEstimateNoticeTypeV2OrderEstimateNoticeMaintenanceWindow:
		// v2OrderEstimateNotice.V2OrderEstimateNoticeMaintenanceWindow is populated
	default:
		// Unknown type - use v2OrderEstimateNotice.GetUnknownRaw() for raw JSON
}
```
