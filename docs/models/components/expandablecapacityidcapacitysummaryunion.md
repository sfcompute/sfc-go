# ExpandableCapacityIDCapacitySummaryUnion

ID (default) or expanded summary when using expand parameter


## Supported Types

### 

```go
expandableCapacityIDCapacitySummaryUnion := components.CreateExpandableCapacityIDCapacitySummaryUnionStr(string{/* values here */})
```

### ExpandableCapacityIDCapacitySummary

```go
expandableCapacityIDCapacitySummaryUnion := components.CreateExpandableCapacityIDCapacitySummaryUnionExpandableCapacityIDCapacitySummary(components.ExpandableCapacityIDCapacitySummary{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch expandableCapacityIDCapacitySummaryUnion.Type {
	case components.ExpandableCapacityIDCapacitySummaryUnionTypeStr:
		// expandableCapacityIDCapacitySummaryUnion.Str is populated
	case components.ExpandableCapacityIDCapacitySummaryUnionTypeExpandableCapacityIDCapacitySummary:
		// expandableCapacityIDCapacitySummaryUnion.ExpandableCapacityIDCapacitySummary is populated
	default:
		// Unknown type - use expandableCapacityIDCapacitySummaryUnion.GetUnknownRaw() for raw JSON
}
```
