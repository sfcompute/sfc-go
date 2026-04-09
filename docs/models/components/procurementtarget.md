# ProcurementTarget


## Supported Types

### NodeCountTag

```go
procurementTarget := components.CreateProcurementTargetNodeCountTag(string{/* values here */})
```

### 

```go
procurementTarget := components.CreateProcurementTargetInteger(int64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch procurementTarget.Type {
	case components.ProcurementTargetTypeNodeCountTag:
		// procurementTarget.NodeCountTag is populated
	case components.ProcurementTargetTypeInteger:
		// procurementTarget.Integer is populated
	default:
		// Unknown type - use procurementTarget.GetUnknownRaw() for raw JSON
}
```
