# V2OrderEstimateResponse


## Supported Types

### V2OrderEstimateResponseSuccess

```go
v2OrderEstimateResponse := components.CreateV2OrderEstimateResponseV2OrderEstimateResponseSuccess(components.V2OrderEstimateResponseSuccess{/* values here */})
```

### V2OrderEstimateResponseUnavailable

```go
v2OrderEstimateResponse := components.CreateV2OrderEstimateResponseV2OrderEstimateResponseUnavailable(components.V2OrderEstimateResponseUnavailable{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2OrderEstimateResponse.Type {
	case components.V2OrderEstimateResponseTypeV2OrderEstimateResponseSuccess:
		// v2OrderEstimateResponse.V2OrderEstimateResponseSuccess is populated
	case components.V2OrderEstimateResponseTypeV2OrderEstimateResponseUnavailable:
		// v2OrderEstimateResponse.V2OrderEstimateResponseUnavailable is populated
	default:
		// Unknown type - use v2OrderEstimateResponse.GetUnknownRaw() for raw JSON
}
```
