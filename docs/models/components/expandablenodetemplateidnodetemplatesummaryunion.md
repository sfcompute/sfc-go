# ExpandableNodeTemplateIDNodeTemplateSummaryUnion

ID (default) or expanded summary when using expand parameter


## Supported Types

### 

```go
expandableNodeTemplateIDNodeTemplateSummaryUnion := components.CreateExpandableNodeTemplateIDNodeTemplateSummaryUnionStr(string{/* values here */})
```

### ExpandableNodeTemplateIDNodeTemplateSummary

```go
expandableNodeTemplateIDNodeTemplateSummaryUnion := components.CreateExpandableNodeTemplateIDNodeTemplateSummaryUnionExpandableNodeTemplateIDNodeTemplateSummary(components.ExpandableNodeTemplateIDNodeTemplateSummary{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch expandableNodeTemplateIDNodeTemplateSummaryUnion.Type {
	case components.ExpandableNodeTemplateIDNodeTemplateSummaryUnionTypeStr:
		// expandableNodeTemplateIDNodeTemplateSummaryUnion.Str is populated
	case components.ExpandableNodeTemplateIDNodeTemplateSummaryUnionTypeExpandableNodeTemplateIDNodeTemplateSummary:
		// expandableNodeTemplateIDNodeTemplateSummaryUnion.ExpandableNodeTemplateIDNodeTemplateSummary is populated
	default:
		// Unknown type - use expandableNodeTemplateIDNodeTemplateSummaryUnion.GetUnknownRaw() for raw JSON
}
```
