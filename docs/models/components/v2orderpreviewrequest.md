# V2OrderPreviewRequest


## Supported Types

### V2OrderPreviewRequestBuy

```go
v2OrderPreviewRequest := components.CreateV2OrderPreviewRequestV2OrderPreviewRequestBuy(components.V2OrderPreviewRequestBuy{/* values here */})
```

### V2OrderPreviewRequestSell

```go
v2OrderPreviewRequest := components.CreateV2OrderPreviewRequestV2OrderPreviewRequestSell(components.V2OrderPreviewRequestSell{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2OrderPreviewRequest.Type {
	case components.V2OrderPreviewRequestTypeV2OrderPreviewRequestBuy:
		// v2OrderPreviewRequest.V2OrderPreviewRequestBuy is populated
	case components.V2OrderPreviewRequestTypeV2OrderPreviewRequestSell:
		// v2OrderPreviewRequest.V2OrderPreviewRequestSell is populated
}
```
