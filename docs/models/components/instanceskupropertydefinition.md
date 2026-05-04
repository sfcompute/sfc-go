# InstanceSkuPropertyDefinition

A catalog entry describing a property key together with its allowed enumeration values.


## Supported Types

### InstanceSkuPropertyDefinitionEnumeration

```go
instanceSkuPropertyDefinition := components.CreateInstanceSkuPropertyDefinitionInstanceSkuPropertyDefinitionEnumeration(components.InstanceSkuPropertyDefinitionEnumeration{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch instanceSkuPropertyDefinition.Type {
	case components.InstanceSkuPropertyDefinitionTypeInstanceSkuPropertyDefinitionEnumeration:
		// instanceSkuPropertyDefinition.InstanceSkuPropertyDefinitionEnumeration is populated
	default:
		// Unknown type - use instanceSkuPropertyDefinition.GetUnknownRaw() for raw JSON
}
```
