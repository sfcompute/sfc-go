# V2CreateOrderRequestDeltaUnion

The desired change in capacity. Will be added if side is `buy`, subtracted if `side` is sell if the order fills.


## Supported Types

### V2CreateOrderRequestDeltaRectangle

```go
v2CreateOrderRequestDeltaUnion := components.CreateV2CreateOrderRequestDeltaUnionV2CreateOrderRequestDeltaRectangle(components.V2CreateOrderRequestDeltaRectangle{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2CreateOrderRequestDeltaUnion.Type {
	case components.V2CreateOrderRequestDeltaUnionTypeV2CreateOrderRequestDeltaRectangle:
		// v2CreateOrderRequestDeltaUnion.V2CreateOrderRequestDeltaRectangle is populated
}
```
