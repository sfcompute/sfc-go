# Permissions

## Overview

Inspect what the caller is allowed to do.

### Available Operations

* [CheckPermissionHandler](#checkpermissionhandler) - Check permissions

## CheckPermissionHandler

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

Evaluate one or more `"resource:verb"` actions against the caller's grants and return a single aggregated verdict. Pass `workspace` to check workspace-scoped grants, or omit it to check org-scoped grants.

### Example Usage

<!-- UsageSnippet language="go" operationID="check_permission_handler" method="post" path="/preview/v2/permissions/check" -->
```go
package main

import(
	"context"
	sfc "github.com/sfcompute/sfc-go"
	"github.com/sfcompute/sfc-go/optionalnullable"
	"github.com/sfcompute/sfc-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sfc.New(
        sfc.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Permissions.CheckPermissionHandler(ctx, components.CheckPermissionRequest{
        Workspace: optionalnullable.From(sfc.Pointer("wksp_k3R-nX9vLm7Qp2Yw5Jd8F")),
        Actions: []string{
            "node:read",
            "order:write",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PermissionCheckResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.CheckPermissionRequest](../../models/components/checkpermissionrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CheckPermissionHandlerResponse](../../models/operations/checkpermissionhandlerresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.BadRequestError     | 400                           | application/json              |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |