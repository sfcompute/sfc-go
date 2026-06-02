# V2OrderPreviewResponse


## Supported Types

### V2OrderPreviewResponseQuoted

```go
v2OrderPreviewResponse := components.CreateV2OrderPreviewResponseV2OrderPreviewResponseQuoted(components.V2OrderPreviewResponseQuoted{/* values here */})
```

### V2OrderPreviewResponseUnavailable

```go
v2OrderPreviewResponse := components.CreateV2OrderPreviewResponseV2OrderPreviewResponseUnavailable(components.V2OrderPreviewResponseUnavailable{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2OrderPreviewResponse.Type {
	case components.V2OrderPreviewResponseTypeV2OrderPreviewResponseQuoted:
		// v2OrderPreviewResponse.V2OrderPreviewResponseQuoted is populated
	case components.V2OrderPreviewResponseTypeV2OrderPreviewResponseUnavailable:
		// v2OrderPreviewResponse.V2OrderPreviewResponseUnavailable is populated
	default:
		// Unknown type - use v2OrderPreviewResponse.GetUnknownRaw() for raw JSON
}
```
