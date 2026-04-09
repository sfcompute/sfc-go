# ExpandableListDeploymentIDDeploymentSummaryUnion

Array of IDs (default) or expanded summaries when using expand parameter


## Supported Types

### 

```go
expandableListDeploymentIDDeploymentSummaryUnion := components.CreateExpandableListDeploymentIDDeploymentSummaryUnionArrayOfStr([]string{/* values here */})
```

### 

```go
expandableListDeploymentIDDeploymentSummaryUnion := components.CreateExpandableListDeploymentIDDeploymentSummaryUnionArrayOfExpandableListDeploymentIDDeploymentSummary([]components.ExpandableListDeploymentIDDeploymentSummary{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch expandableListDeploymentIDDeploymentSummaryUnion.Type {
	case components.ExpandableListDeploymentIDDeploymentSummaryUnionTypeArrayOfStr:
		// expandableListDeploymentIDDeploymentSummaryUnion.ArrayOfStr is populated
	case components.ExpandableListDeploymentIDDeploymentSummaryUnionTypeArrayOfExpandableListDeploymentIDDeploymentSummary:
		// expandableListDeploymentIDDeploymentSummaryUnion.ArrayOfExpandableListDeploymentIDDeploymentSummary is populated
	default:
		// Unknown type - use expandableListDeploymentIDDeploymentSummaryUnion.GetUnknownRaw() for raw JSON
}
```
