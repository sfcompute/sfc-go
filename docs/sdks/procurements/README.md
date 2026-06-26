# Procurements

## Overview

Market automations that maintain capacity by placing buy/sell orders.

### Available Operations

* [List](#list) - List procurements
* [Create](#create) - Create procurement
* [GetProcurement](#getprocurement) - Get procurement
* [Delete](#delete) - Delete procurement
* [PatchProcurement](#patchprocurement) - Update procurement

## List

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

List all procurements.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_procurements" method="get" path="/preview/v2/procurements" -->
```go
package main

import(
	"context"
	sfc "github.com/sfcompute/sfc-go"
	"github.com/sfcompute/sfc-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sfc.New(
        sfc.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Procurements.List(ctx, operations.ListProcurementsRequest{
        Workspace: "wksp_k3R-nX9vLm7Qp2Yw5Jd8F",
        Pool: sfc.Pointer("pool_k3R-nX9vLm7Qp2Yw5Jd8F"),
        StartingAfter: sfc.Pointer("procc_gqXR7s0Kj5mHvE2wNpLc4Q"),
        EndingBefore: sfc.Pointer("procc_gqXR7s0Kj5mHvE2wNpLc4Q"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListProcurementsResponse != nil {
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListProcurementsRequest](../../models/operations/listprocurementsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListProcurementsResponse](../../models/operations/listprocurementsresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Create

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

Create a market automation that maintains capacity by placing buy/sell orders.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_procurement" method="post" path="/preview/v2/procurements" -->
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

    res, err := s.Procurements.Create(ctx, components.CreateProcurementRequest{
        Name: optionalnullable.From(sfc.Pointer("my-resource-name")),
        Target: components.CreateProcurementTargetNodeCountTag(
            "<value>",
        ),
        Pool: "<value>",
        InstanceSku: "isku_k3R-nX9vLm7Qp2Yw5Jd8F",
        MinSellPriceDollarsPerNodeHour: "2.500000",
        MaxBuyPriceDollarsPerNodeHour: "2.500000",
        ManagedWindowMinutes: 423588,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ProcurementResponse != nil {
        switch res.ProcurementResponse.Target.Type {
            case components.ProcurementTargetTypeNodeCountTag:
                // res.ProcurementResponse.Target.NodeCountTag is populated
            case components.ProcurementTargetTypeInteger:
                // res.ProcurementResponse.Target.Integer is populated
            default:
                // Unknown type - use res.ProcurementResponse.Target.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.CreateProcurementRequest](../../models/components/createprocurementrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateProcurementResponse](../../models/operations/createprocurementresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.BadRequestError          | 400                                | application/json                   |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.ConflictError            | 409                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## GetProcurement

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

Retrieve a procurement by ID or name.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_procurement" method="get" path="/preview/v2/procurements/{id}" -->
```go
package main

import(
	"context"
	sfc "github.com/sfcompute/sfc-go"
	"log"
	"github.com/sfcompute/sfc-go/models/components"
)

func main() {
    ctx := context.Background()

    s := sfc.New(
        sfc.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Procurements.GetProcurement(ctx, "proc_k3R-nX9vLm7Qp2Yw5Jd8F")
    if err != nil {
        log.Fatal(err)
    }
    if res.ProcurementResponse != nil {
        switch res.ProcurementResponse.Target.Type {
            case components.ProcurementTargetTypeNodeCountTag:
                // res.ProcurementResponse.Target.NodeCountTag is populated
            case components.ProcurementTargetTypeInteger:
                // res.ProcurementResponse.Target.Integer is populated
            default:
                // Unknown type - use res.ProcurementResponse.Target.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | proc_k3R-nX9vLm7Qp2Yw5Jd8F                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetProcurementResponse](../../models/operations/getprocurementresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.ForbiddenError      | 403                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

Delete a procurement. Standing orders are cancelled automatically.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_procurement" method="delete" path="/preview/v2/procurements/{id}" -->
```go
package main

import(
	"context"
	sfc "github.com/sfcompute/sfc-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sfc.New(
        sfc.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Procurements.Delete(ctx, "proc_k3R-nX9vLm7Qp2Yw5Jd8F")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | proc_k3R-nX9vLm7Qp2Yw5Jd8F                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteProcurementResponse](../../models/operations/deleteprocurementresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.ForbiddenError      | 403                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## PatchProcurement

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

Update a procurement's configuration.

### Example Usage

<!-- UsageSnippet language="go" operationID="patch_procurement" method="patch" path="/preview/v2/procurements/{id}" -->
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

    res, err := s.Procurements.PatchProcurement(ctx, "proc_k3R-nX9vLm7Qp2Yw5Jd8F", components.PatchProcurementRequest{
        Name: optionalnullable.From(sfc.Pointer("my-resource-name")),
        InstanceSku: optionalnullable.From(sfc.Pointer("isku_k3R-nX9vLm7Qp2Yw5Jd8F")),
        MinSellPriceDollarsPerNodeHour: optionalnullable.From(sfc.Pointer("2.500000")),
        MaxBuyPriceDollarsPerNodeHour: optionalnullable.From(sfc.Pointer("2.500000")),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ProcurementResponse != nil {
        switch res.ProcurementResponse.Target.Type {
            case components.ProcurementTargetTypeNodeCountTag:
                // res.ProcurementResponse.Target.NodeCountTag is populated
            case components.ProcurementTargetTypeInteger:
                // res.ProcurementResponse.Target.Integer is populated
            default:
                // Unknown type - use res.ProcurementResponse.Target.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `id`                                                                                     | `string`                                                                                 | :heavy_check_mark:                                                                       | N/A                                                                                      | proc_k3R-nX9vLm7Qp2Yw5Jd8F                                                               |
| `body`                                                                                   | [components.PatchProcurementRequest](../../models/components/patchprocurementrequest.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.PatchProcurementResponse](../../models/operations/patchprocurementresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.BadRequestError          | 400                                | application/json                   |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |