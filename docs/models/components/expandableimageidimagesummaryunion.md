# ExpandableImageIDImageSummaryUnion

ID (default) or expanded summary when using expand parameter


## Supported Types

### 

```go
expandableImageIDImageSummaryUnion := components.CreateExpandableImageIDImageSummaryUnionStr(string{/* values here */})
```

### ExpandableImageIDImageSummary

```go
expandableImageIDImageSummaryUnion := components.CreateExpandableImageIDImageSummaryUnionExpandableImageIDImageSummary(components.ExpandableImageIDImageSummary{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch expandableImageIDImageSummaryUnion.Type {
	case components.ExpandableImageIDImageSummaryUnionTypeStr:
		// expandableImageIDImageSummaryUnion.Str is populated
	case components.ExpandableImageIDImageSummaryUnionTypeExpandableImageIDImageSummary:
		// expandableImageIDImageSummaryUnion.ExpandableImageIDImageSummary is populated
	default:
		// Unknown type - use expandableImageIDImageSummaryUnion.GetUnknownRaw() for raw JSON
}
```
