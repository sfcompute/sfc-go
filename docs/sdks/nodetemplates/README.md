# NodeTemplates

## Overview

Reusable node configuration.

### Available Operations

* [List](#list) - List node templates
* [Create](#create) - Create node template
* [Fetch](#fetch) - Get node template
* [Delete](#delete) - Delete node template

## List

List all node templates.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_node_templates" method="get" path="/v2/node_templates" -->
```go
package main

import(
	"context"
	"os"
	sfc "github.com/sfcompute/sfc-go-sdk"
	"github.com/sfcompute/sfc-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sfc.New(
        sfc.WithSecurity(os.Getenv("SFC_BEARER_AUTH")),
    )

    res, err := s.NodeTemplates.List(ctx, operations.ListNodeTemplatesRequest{
        Workspace: "wksp_k3R-nX9vLm7Qp2Yw5Jd8F",
        ID: []string{
            "ntmpl_k3R-nX9vLm7Qp2Yw5Jd8F",
        },
        StartingAfter: sfc.Pointer("ntmplc_gqXR7s0Kj5mHvE2wNpLc4Q"),
        EndingBefore: sfc.Pointer("ntmplc_gqXR7s0Kj5mHvE2wNpLc4Q"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListNodeTemplatesResponse != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListNodeTemplatesRequest](../../models/operations/listnodetemplatesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListNodeTemplatesResponse](../../models/operations/listnodetemplatesresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Create

Create a reusable node configuration.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_node_template" method="post" path="/v2/node_templates" -->
```go
package main

import(
	"context"
	"os"
	sfc "github.com/sfcompute/sfc-go-sdk"
	"github.com/sfcompute/sfc-go-sdk/optionalnullable"
	"github.com/sfcompute/sfc-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sfc.New(
        sfc.WithSecurity(os.Getenv("SFC_BEARER_AUTH")),
    )

    res, err := s.NodeTemplates.Create(ctx, components.CreateNodeTemplateRequest{
        Name: optionalnullable.From(sfc.Pointer("my-resource-name")),
        Workspace: "wksp_k3R-nX9vLm7Qp2Yw5Jd8F",
        Image: "image_k3R-nX9vLm7Qp2Yw5Jd8F",
        CloudInitUserData: sfc.Pointer("IyEvYmluL2Jhc2gKZWNobyBoZWxsbyB3b3JsZAo="),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.NodeTemplateResponse != nil {
        switch res.NodeTemplateResponse.Image.Type {
            case components.ExpandableImageIDImageSummaryUnionTypeStr:
                // res.NodeTemplateResponse.Image.Str is populated
            case components.ExpandableImageIDImageSummaryUnionTypeExpandableImageIDImageSummary:
                // res.NodeTemplateResponse.Image.ExpandableImageIDImageSummary is populated
            default:
                // Unknown type - use res.NodeTemplateResponse.Image.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.CreateNodeTemplateRequest](../../models/components/createnodetemplaterequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateNodeTemplateResponse](../../models/operations/createnodetemplateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Fetch

Retrieve a node template by ID, resource path, or name.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch_node_template" method="get" path="/v2/node_templates/{id}" -->
```go
package main

import(
	"context"
	"os"
	sfc "github.com/sfcompute/sfc-go-sdk"
	"log"
	"github.com/sfcompute/sfc-go-sdk/models/components"
)

func main() {
    ctx := context.Background()

    s := sfc.New(
        sfc.WithSecurity(os.Getenv("SFC_BEARER_AUTH")),
    )

    res, err := s.NodeTemplates.Fetch(ctx, "ntmpl_k3R-nX9vLm7Qp2Yw5Jd8F", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.NodeTemplateResponse != nil {
        switch res.NodeTemplateResponse.Image.Type {
            case components.ExpandableImageIDImageSummaryUnionTypeStr:
                // res.NodeTemplateResponse.Image.Str is populated
            case components.ExpandableImageIDImageSummaryUnionTypeExpandableImageIDImageSummary:
                // res.NodeTemplateResponse.Image.ExpandableImageIDImageSummary is populated
            default:
                // Unknown type - use res.NodeTemplateResponse.Image.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | ntmpl_k3R-nX9vLm7Qp2Yw5Jd8F                              |
| `expand`                                                 | []`string`                                               | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchNodeTemplateResponse](../../models/operations/fetchnodetemplateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete a node template. The template must not be in use by any capacity.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_node_template" method="delete" path="/v2/node_templates/{id}" -->
```go
package main

import(
	"context"
	"os"
	sfc "github.com/sfcompute/sfc-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := sfc.New(
        sfc.WithSecurity(os.Getenv("SFC_BEARER_AUTH")),
    )

    res, err := s.NodeTemplates.Delete(ctx, "ntmpl_k3R-nX9vLm7Qp2Yw5Jd8F")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | ntmpl_k3R-nX9vLm7Qp2Yw5Jd8F                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteNodeTemplateResponse](../../models/operations/deletenodetemplateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |