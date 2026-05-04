# ExpandableInstanceTemplateIDInstanceTemplateSummaryUnion

ID (default) or expanded summary when using expand parameter


## Supported Types

### 

```go
expandableInstanceTemplateIDInstanceTemplateSummaryUnion := components.CreateExpandableInstanceTemplateIDInstanceTemplateSummaryUnionStr(string{/* values here */})
```

### ExpandableInstanceTemplateIDInstanceTemplateSummary

```go
expandableInstanceTemplateIDInstanceTemplateSummaryUnion := components.CreateExpandableInstanceTemplateIDInstanceTemplateSummaryUnionExpandableInstanceTemplateIDInstanceTemplateSummary(components.ExpandableInstanceTemplateIDInstanceTemplateSummary{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch expandableInstanceTemplateIDInstanceTemplateSummaryUnion.Type {
	case components.ExpandableInstanceTemplateIDInstanceTemplateSummaryUnionTypeStr:
		// expandableInstanceTemplateIDInstanceTemplateSummaryUnion.Str is populated
	case components.ExpandableInstanceTemplateIDInstanceTemplateSummaryUnionTypeExpandableInstanceTemplateIDInstanceTemplateSummary:
		// expandableInstanceTemplateIDInstanceTemplateSummaryUnion.ExpandableInstanceTemplateIDInstanceTemplateSummary is populated
	default:
		// Unknown type - use expandableInstanceTemplateIDInstanceTemplateSummaryUnion.GetUnknownRaw() for raw JSON
}
```
