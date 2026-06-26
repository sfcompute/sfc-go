# Pools

## Overview

A bucket of owned compute balance over time.

### Available Operations

* [List](#list) - List pools
* [Create](#create) - Create pool
* [Fetch](#fetch) - Get pool
* [Delete](#delete) - Delete pool
* [Update](#update) - Update pool
* [ListPoolTransfers](#listpooltransfers) - List pool transfers
* [CreatePoolTransfer](#createpooltransfer) - Create pool transfer
* [FetchPoolTransfer](#fetchpooltransfer) - Get pool transfer

## List

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

List all pools.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_pools" method="get" path="/preview/v2/pools" -->
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

    res, err := s.Pools.List(ctx, operations.ListPoolsRequest{
        Workspace: "wksp_k3R-nX9vLm7Qp2Yw5Jd8F",
        ID: []string{
            "pool_k3R-nX9vLm7Qp2Yw5Jd8F",
        },
        StartingAfter: sfc.Pointer("poolc_gqXR7s0Kj5mHvE2wNpLc4Q"),
        EndingBefore: sfc.Pointer("poolc_gqXR7s0Kj5mHvE2wNpLc4Q"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPoolsResponse != nil {
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ListPoolsRequest](../../models/operations/listpoolsrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.ListPoolsResponse](../../models/operations/listpoolsresponse.md), error**

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

Create a pool to hold compute.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_pool" method="post" path="/preview/v2/pools" -->
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

    res, err := s.Pools.Create(ctx, components.V2CreatePoolRequest{
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
    if res.PoolResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.V2CreatePoolRequest](../../models/components/v2createpoolrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.CreatePoolResponse](../../models/operations/createpoolresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Fetch

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

Retrieve a pool by ID, resource path, or name, including its compute schedule.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch_pool" method="get" path="/preview/v2/pools/{id}" -->
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

    res, err := s.Pools.Fetch(ctx, "pool_k3R-nX9vLm7Qp2Yw5Jd8F", sfc.Pointer[int64](0))
    if err != nil {
        log.Fatal(err)
    }
    if res.PoolResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | pool_k3R-nX9vLm7Qp2Yw5Jd8F                               |
| `scheduleHistoryMinutes`                                 | `*int64`                                                 | :heavy_minus_sign:                                       | How many minutes of past schedule to include.            |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchPoolResponse](../../models/operations/fetchpoolresponse.md), error**

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

Delete a pool. The pool must have no active orders, future allocations, active nodes, deployments, or procurements. Remove all dependencies before deleting.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_pool" method="delete" path="/preview/v2/pools/{id}" -->
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

    res, err := s.Pools.Delete(ctx, "pool_k3R-nX9vLm7Qp2Yw5Jd8F")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | pool_k3R-nX9vLm7Qp2Yw5Jd8F                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePoolResponse](../../models/operations/deletepoolresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Update

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

Update a pool. Omitted fields are left unchanged.

### Example Usage

<!-- UsageSnippet language="go" operationID="update_pool" method="patch" path="/preview/v2/pools/{id}" -->
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

    res, err := s.Pools.Update(ctx, "pool_k3R-nX9vLm7Qp2Yw5Jd8F", components.V2PatchPoolRequest{
        Name: optionalnullable.From(sfc.Pointer("my-resource-name")),
        Tags: optionalnullable.From(sfc.Pointer(map[string]string{
            "env": "prod",
            "team": "infra",
        })),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PoolResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `id`                                                                           | `string`                                                                       | :heavy_check_mark:                                                             | N/A                                                                            | pool_k3R-nX9vLm7Qp2Yw5Jd8F                                                     |
| `body`                                                                         | [components.V2PatchPoolRequest](../../models/components/v2patchpoolrequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.UpdatePoolResponse](../../models/operations/updatepoolresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## ListPoolTransfers

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

List pool transfers for the caller's organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_pool_transfers" method="get" path="/preview/v2/pool_transfers" -->
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

    res, err := s.Pools.ListPoolTransfers(ctx, operations.ListPoolTransfersRequest{
        Pool: []string{
            "pool_k3R-nX9vLm7Qp2Yw5Jd8F",
        },
        StartingAfter: sfc.Pointer("ptfrc_gqXR7s0Kj5mHvE2wNpLc4Q"),
        EndingBefore: sfc.Pointer("ptfrc_gqXR7s0Kj5mHvE2wNpLc4Q"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V2ListPoolTransfersResponse != nil {
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
| `request`                                                                                  | [operations.ListPoolTransfersRequest](../../models/operations/listpooltransfersrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListPoolTransfersResponse](../../models/operations/listpooltransfersresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## CreatePoolTransfer

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

Transfer some or all of one pool into another

### Example Usage

<!-- UsageSnippet language="go" operationID="create_pool_transfer" method="post" path="/preview/v2/pool_transfers" -->
```go
package main

import(
	"context"
	sfc "github.com/sfcompute/sfc-go"
	"github.com/sfcompute/sfc-go/models/components"
	"github.com/sfcompute/sfc-go/optionalnullable"
	"log"
)

func main() {
    ctx := context.Background()

    s := sfc.New(
        sfc.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Pools.CreatePoolTransfer(ctx, components.V2CreatePoolTransferRequest{
        FromPool: "pool_k3R-nX9vLm7Qp2Yw5Jd8F",
        ToPool: "pool_k3R-nX9vLm7Qp2Yw5Jd8F",
        AllocationScheduleDelta: []components.ScheduleEntry{},
        InstanceSku: "isku_k3R-nX9vLm7Qp2Yw5Jd8F",
    }, optionalnullable.From[string](nil))
    if err != nil {
        log.Fatal(err)
    }
    if res.V2PoolTransferResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `body`                                                                                           | [components.V2CreatePoolTransferRequest](../../models/components/v2createpooltransferrequest.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `idempotencyKey`                                                                                 | optionalnullable.OptionalNullable[`string`]                                                      | :heavy_minus_sign:                                                                               | Unique key for idempotent transfer creation.                                                     |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreatePoolTransferResponse](../../models/operations/createpooltransferresponse.md), error**

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

## FetchPoolTransfer

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

Retrieve a pool transfer by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch_pool_transfer" method="get" path="/preview/v2/pool_transfers/{id}" -->
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

    res, err := s.Pools.FetchPoolTransfer(ctx, "pxfr_abc123")
    if err != nil {
        log.Fatal(err)
    }
    if res.V2PoolTransferResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | Pool transfer ID                                         | pxfr_k3R-nX9vLm7Qp2Yw5Jd8F                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchPoolTransferResponse](../../models/operations/fetchpooltransferresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.ForbiddenError      | 403                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |