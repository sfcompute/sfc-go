# InstanceSkuProperty

A property set on an instance SKU. Tagged by `type` so new value kinds can be added without breaking clients.


## Supported Types

### InstanceSkuPropertyEnumeration

```go
instanceSkuProperty := components.CreateInstanceSkuPropertyInstanceSkuPropertyEnumeration(components.InstanceSkuPropertyEnumeration{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch instanceSkuProperty.Type {
	case components.InstanceSkuPropertyTypeInstanceSkuPropertyEnumeration:
		// instanceSkuProperty.InstanceSkuPropertyEnumeration is populated
	default:
		// Unknown type - use instanceSkuProperty.GetUnknownRaw() for raw JSON
}
```
