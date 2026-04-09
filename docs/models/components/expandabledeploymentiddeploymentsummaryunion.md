# ExpandableDeploymentIDDeploymentSummaryUnion

ID (default) or expanded summary when using expand parameter


## Supported Types

### 

```go
expandableDeploymentIDDeploymentSummaryUnion := components.CreateExpandableDeploymentIDDeploymentSummaryUnionStr(string{/* values here */})
```

### ExpandableDeploymentIDDeploymentSummary

```go
expandableDeploymentIDDeploymentSummaryUnion := components.CreateExpandableDeploymentIDDeploymentSummaryUnionExpandableDeploymentIDDeploymentSummary(components.ExpandableDeploymentIDDeploymentSummary{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch expandableDeploymentIDDeploymentSummaryUnion.Type {
	case components.ExpandableDeploymentIDDeploymentSummaryUnionTypeStr:
		// expandableDeploymentIDDeploymentSummaryUnion.Str is populated
	case components.ExpandableDeploymentIDDeploymentSummaryUnionTypeExpandableDeploymentIDDeploymentSummary:
		// expandableDeploymentIDDeploymentSummaryUnion.ExpandableDeploymentIDDeploymentSummary is populated
	default:
		// Unknown type - use expandableDeploymentIDDeploymentSummaryUnion.GetUnknownRaw() for raw JSON
}
```
