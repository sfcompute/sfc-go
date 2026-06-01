# Capacities

## Overview

A bucket of owned compute balance over time.

### Available Operations

* [List](#list) - List capacities
* [Create](#create) - Create capacity
* [Fetch](#fetch) - Get capacity
* [Delete](#delete) - Delete capacity
* [Update](#update) - Update capacity
* [ListCapacityTransfers](#listcapacitytransfers) - List capacity transfers
* [CreateCapacityTransfer](#createcapacitytransfer) - Create capacity transfer
* [FetchCapacityTransfer](#fetchcapacitytransfer) - Get capacity transfer

## List

List all capacities.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_capacities" method="get" path="/preview/v2/capacities" -->
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

Create a capacity to hold compute.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_capacity" method="post" path="/preview/v2/capacities" -->
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

    res, err := s.Capacities.Create(ctx, components.CreateCapacityRequest{
        Name: optionalnullable.From(sfc.Pointer("my-resource-name")),
        Workspace: "wksp_k3R-nX9vLm7Qp2Yw5Jd8F",
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

<!-- UsageSnippet language="go" operationID="fetch_capacity" method="get" path="/preview/v2/capacities/{id}" -->
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

<!-- UsageSnippet language="go" operationID="delete_capacity" method="delete" path="/preview/v2/capacities/{id}" -->
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

<!-- UsageSnippet language="go" operationID="update_capacity" method="patch" path="/preview/v2/capacities/{id}" -->
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

    res, err := s.Capacities.Update(ctx, "cap_k3R-nX9vLm7Qp2Yw5Jd8F", components.PatchCapacityRequest{
        Name: optionalnullable.From(sfc.Pointer("my-resource-name")),
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

## ListCapacityTransfers

List capacity transfers for the caller's organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_capacity_transfers" method="get" path="/preview/v2/capacity_transfers" -->
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

    res, err := s.Capacities.ListCapacityTransfers(ctx, operations.ListCapacityTransfersRequest{
        Capacity: []string{
            "cap_k3R-nX9vLm7Qp2Yw5Jd8F",
        },
        StartingAfter: sfc.Pointer("ctfrc_gqXR7s0Kj5mHvE2wNpLc4Q"),
        EndingBefore: sfc.Pointer("ctfrc_gqXR7s0Kj5mHvE2wNpLc4Q"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V2ListCapacityTransfersResponse != nil {
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListCapacityTransfersRequest](../../models/operations/listcapacitytransfersrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListCapacityTransfersResponse](../../models/operations/listcapacitytransfersresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## CreateCapacityTransfer

Transfer some or all of one capacity into another

### Example Usage

<!-- UsageSnippet language="go" operationID="create_capacity_transfer" method="post" path="/preview/v2/capacity_transfers" -->
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

    res, err := s.Capacities.CreateCapacityTransfer(ctx, components.V2CreateCapacityTransferRequest{
        FromCapacity: "cap_k3R-nX9vLm7Qp2Yw5Jd8F",
        ToCapacity: "cap_k3R-nX9vLm7Qp2Yw5Jd8F",
        AllocationScheduleDelta: []components.ScheduleEntry{
            components.ScheduleEntry{
                StartAt: 1738972800,
                EndAt: optionalnullable.From(sfc.Pointer[int64](1738972800)),
                NodeCount: 569495,
            },
        },
        InstanceSku: "isku_k3R-nX9vLm7Qp2Yw5Jd8F",
    }, optionalnullable.From[string](nil))
    if err != nil {
        log.Fatal(err)
    }
    if res.V2CapacityTransferResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `body`                                                                                                   | [components.V2CreateCapacityTransferRequest](../../models/components/v2createcapacitytransferrequest.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `idempotencyKey`                                                                                         | optionalnullable.OptionalNullable[`string`]                                                              | :heavy_minus_sign:                                                                                       | Unique key for idempotent transfer creation.                                                             |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreateCapacityTransferResponse](../../models/operations/createcapacitytransferresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.ServiceUnavailableError  | 503                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## FetchCapacityTransfer

Retrieve a capacity transfer by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch_capacity_transfer" method="get" path="/preview/v2/capacity_transfers/{id}" -->
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

    res, err := s.Capacities.FetchCapacityTransfer(ctx, "cxfr_abc123")
    if err != nil {
        log.Fatal(err)
    }
    if res.V2CapacityTransferResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | Capacity transfer ID                                     | cxfr_k3R-nX9vLm7Qp2Yw5Jd8F                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchCapacityTransferResponse](../../models/operations/fetchcapacitytransferresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |