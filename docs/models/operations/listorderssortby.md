# ListOrdersSortBy

Prefix with `-` for descending.


## Supported Types

### V2OrderSortBy

```go
listOrdersSortBy := operations.CreateListOrdersSortByV2OrderSortBy(components.V2OrderSortBy{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listOrdersSortBy.Type {
	case operations.ListOrdersSortByTypeV2OrderSortBy:
		// listOrdersSortBy.V2OrderSortBy is populated
}
```
