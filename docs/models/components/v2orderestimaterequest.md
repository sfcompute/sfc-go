# V2OrderEstimateRequest


## Supported Types

### V2OrderEstimateRequestBuy

```go
v2OrderEstimateRequest := components.CreateV2OrderEstimateRequestV2OrderEstimateRequestBuy(components.V2OrderEstimateRequestBuy{/* values here */})
```

### V2OrderEstimateRequestSell

```go
v2OrderEstimateRequest := components.CreateV2OrderEstimateRequestV2OrderEstimateRequestSell(components.V2OrderEstimateRequestSell{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2OrderEstimateRequest.Type {
	case components.V2OrderEstimateRequestTypeV2OrderEstimateRequestBuy:
		// v2OrderEstimateRequest.V2OrderEstimateRequestBuy is populated
	case components.V2OrderEstimateRequestTypeV2OrderEstimateRequestSell:
		// v2OrderEstimateRequest.V2OrderEstimateRequestSell is populated
}
```
