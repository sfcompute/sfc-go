# V2OrderPreviewNotice


## Supported Types

### V2OrderPreviewNoticeMaintenanceWindow

```go
v2OrderPreviewNotice := components.CreateV2OrderPreviewNoticeV2OrderPreviewNoticeMaintenanceWindow(components.V2OrderPreviewNoticeMaintenanceWindow{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2OrderPreviewNotice.Type {
	case components.V2OrderPreviewNoticeTypeV2OrderPreviewNoticeMaintenanceWindow:
		// v2OrderPreviewNotice.V2OrderPreviewNoticeMaintenanceWindow is populated
	default:
		// Unknown type - use v2OrderPreviewNotice.GetUnknownRaw() for raw JSON
}
```
