# V2OrderResponseDeltaUnion

The desired change in capacity. Will be added if side is `buy`, subtracted if `side` is sell if the order fills.


## Supported Types

### V2OrderResponseDeltaRectangle

```go
v2OrderResponseDeltaUnion := components.CreateV2OrderResponseDeltaUnionV2OrderResponseDeltaRectangle(components.V2OrderResponseDeltaRectangle{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2OrderResponseDeltaUnion.Type {
	case components.V2OrderResponseDeltaUnionTypeV2OrderResponseDeltaRectangle:
		// v2OrderResponseDeltaUnion.V2OrderResponseDeltaRectangle is populated
	default:
		// Unknown type - use v2OrderResponseDeltaUnion.GetUnknownRaw() for raw JSON
}
```
