# ExpandableListProcurementIDProcurementSummaryUnion

Array of IDs (default) or expanded summaries when using expand parameter


## Supported Types

### 

```go
expandableListProcurementIDProcurementSummaryUnion := components.CreateExpandableListProcurementIDProcurementSummaryUnionArrayOfStr([]string{/* values here */})
```

### 

```go
expandableListProcurementIDProcurementSummaryUnion := components.CreateExpandableListProcurementIDProcurementSummaryUnionArrayOfExpandableListProcurementIDProcurementSummary([]components.ExpandableListProcurementIDProcurementSummary{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch expandableListProcurementIDProcurementSummaryUnion.Type {
	case components.ExpandableListProcurementIDProcurementSummaryUnionTypeArrayOfStr:
		// expandableListProcurementIDProcurementSummaryUnion.ArrayOfStr is populated
	case components.ExpandableListProcurementIDProcurementSummaryUnionTypeArrayOfExpandableListProcurementIDProcurementSummary:
		// expandableListProcurementIDProcurementSummaryUnion.ArrayOfExpandableListProcurementIDProcurementSummary is populated
	default:
		// Unknown type - use expandableListProcurementIDProcurementSummaryUnion.GetUnknownRaw() for raw JSON
}
```
