# Capacities

## Overview

A bucket of owned compute balance over time.

### Available Operations

* [List](#list) - List capacities
* [Create](#create) - Create capacity
* [Fetch](#fetch) - Get capacity
* [Delete](#delete) - Delete capacity
* [Update](#update) - Update capacity

## List

List all capacities.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_capacities" method="get" path="/v2/capacities" -->
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

    res, err := s.Capacities.List(ctx, operations.ListCapacitiesRequest{
        Workspace: "wksp_k3R-nX9vLm7Qp2Yw5Jd8F",
        ID: []string{
            "cap_k3R-nX9vLm7Qp2Yw5Jd8F",
        },
        StartingAfter: sfc.Pointer("capc_gqXR7s0Kj5mHvE2wNpLc4Q"),
        EndingBefore: sfc.Pointer("capc_gqXR7s0Kj5mHvE2wNpLc4Q"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCapacitiesResponse != nil {
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListCapacitiesRequest](../../models/operations/listcapacitiesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListCapacitiesResponse](../../models/operations/listcapacitiesresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Create

Create a capacity to hold compute in specified zones.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_capacity" method="post" path="/v2/capacities" -->
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

    res, err := s.Capacities.Create(ctx, components.CreateCapacityRequest{
        Name: optionalnullable.From(sfc.Pointer("my-resource-name")),
        Workspace: "wksp_k3R-nX9vLm7Qp2Yw5Jd8F",
        Zones: []string{
            "richmond",
        },
        Tags: optionalnullable.From(sfc.Pointer(map[string]string{
            "env": "prod",
            "team": "infra",
        })),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CapacityResponse != nil {
        switch res.CapacityResponse.Procurements.Type {
            case components.ExpandableListProcurementIDProcurementSummaryUnionTypeArrayOfStr:
                // res.CapacityResponse.Procurements.ArrayOfStr is populated
            case components.ExpandableListProcurementIDProcurementSummaryUnionTypeArrayOfExpandableListProcurementIDProcurementSummary:
                // res.CapacityResponse.Procurements.ArrayOfExpandableListProcurementIDProcurementSummary is populated
            default:
                // Unknown type - use res.CapacityResponse.Procurements.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.CreateCapacityRequest](../../models/components/createcapacityrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateCapacityResponse](../../models/operations/createcapacityresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Fetch

Retrieve a capacity by ID, resource path, or name, including its compute schedule.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch_capacity" method="get" path="/v2/capacities/{id}" -->
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

    res, err := s.Capacities.Fetch(ctx, "cap_k3R-nX9vLm7Qp2Yw5Jd8F", sfc.Pointer[int64](0), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CapacityResponse != nil {
        switch res.CapacityResponse.Procurements.Type {
            case components.ExpandableListProcurementIDProcurementSummaryUnionTypeArrayOfStr:
                // res.CapacityResponse.Procurements.ArrayOfStr is populated
            case components.ExpandableListProcurementIDProcurementSummaryUnionTypeArrayOfExpandableListProcurementIDProcurementSummary:
                // res.CapacityResponse.Procurements.ArrayOfExpandableListProcurementIDProcurementSummary is populated
            default:
                // Unknown type - use res.CapacityResponse.Procurements.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `id`                                                                               | `string`                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                | cap_k3R-nX9vLm7Qp2Yw5Jd8F                                                          |
| `scheduleHistoryMinutes`                                                           | `*int64`                                                                           | :heavy_minus_sign:                                                                 | How many minutes of past schedule to include.                                      |                                                                                    |
| `expand`                                                                           | [][operations.FetchCapacityExpand](../../models/operations/fetchcapacityexpand.md) | :heavy_minus_sign:                                                                 | Expand related resources inline instead of returning IDs.                          |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.FetchCapacityResponse](../../models/operations/fetchcapacityresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete a capacity. The capacity must have no active orders, future allocations, active nodes, deployments, or procurements. Remove all dependencies before deleting.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_capacity" method="delete" path="/v2/capacities/{id}" -->
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

    res, err := s.Capacities.Delete(ctx, "cap_k3R-nX9vLm7Qp2Yw5Jd8F")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | cap_k3R-nX9vLm7Qp2Yw5Jd8F                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteCapacityResponse](../../models/operations/deletecapacityresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Update

Update a capacity. Omitted fields are left unchanged.

### Example Usage

<!-- UsageSnippet language="go" operationID="update_capacity" method="patch" path="/v2/capacities/{id}" -->
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

    res, err := s.Capacities.Update(ctx, "cap_k3R-nX9vLm7Qp2Yw5Jd8F", components.PatchCapacityRequest{
        Name: optionalnullable.From(sfc.Pointer("my-resource-name")),
        Zones: optionalnullable.From(sfc.Pointer([]string{
            "richmond",
        })),
        Tags: optionalnullable.From(sfc.Pointer(map[string]string{
            "env": "prod",
            "team": "infra",
        })),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CapacityResponse != nil {
        switch res.CapacityResponse.Procurements.Type {
            case components.ExpandableListProcurementIDProcurementSummaryUnionTypeArrayOfStr:
                // res.CapacityResponse.Procurements.ArrayOfStr is populated
            case components.ExpandableListProcurementIDProcurementSummaryUnionTypeArrayOfExpandableListProcurementIDProcurementSummary:
                // res.CapacityResponse.Procurements.ArrayOfExpandableListProcurementIDProcurementSummary is populated
            default:
                // Unknown type - use res.CapacityResponse.Procurements.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `id`                                                                               | `string`                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                | cap_k3R-nX9vLm7Qp2Yw5Jd8F                                                          |
| `body`                                                                             | [components.PatchCapacityRequest](../../models/components/patchcapacityrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.UpdateCapacityResponse](../../models/operations/updatecapacityresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |