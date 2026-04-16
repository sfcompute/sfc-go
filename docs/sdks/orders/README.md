# Orders

## Overview

Place orders targeting a capacity to increase your reserved compute balance during some time period.

### Available Operations

* [List](#list) - List orders
* [Create](#create) - Create order
* [Fetch](#fetch) - Get order
* [Cancel](#cancel) - Cancel order

## List

List all orders.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_orders" method="get" path="/v2/orders" -->
```go
package main

import(
	"context"
	"os"
	sfc "github.com/sfcompute/sfc-go"
	"github.com/sfcompute/sfc-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sfc.New(
        sfc.WithSecurity(os.Getenv("SFC_BEARER_AUTH")),
    )

    res, err := s.Orders.List(ctx, operations.ListOrdersRequest{
        ID: []string{
            "ordr_k3R-nX9vLm7Qp2Yw5Jd8F",
        },
        Capacity: sfc.Pointer("cap_k3R-nX9vLm7Qp2Yw5Jd8F"),
        CreatedAfter: sfc.Pointer[int64](1738972800),
        CreatedBefore: sfc.Pointer[int64](1738972800),
        StartingAfter: sfc.Pointer("ordrc_gqXR7s0Kj5mHvE2wNpLc4Q"),
        EndingBefore: sfc.Pointer("ordrc_gqXR7s0Kj5mHvE2wNpLc4Q"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V2ListOrdersResponse != nil {
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListOrdersRequest](../../models/operations/listordersrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.ListOrdersResponse](../../models/operations/listordersresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Create

Place a buy or sell order. Orders fill completely or not at all. All nodes fill in a single zone. Order filling is asynchronous; poll `GET /v2/orders/{id}` to check status.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_order" method="post" path="/v2/orders" -->
```go
package main

import(
	"context"
	"os"
	sfc "github.com/sfcompute/sfc-go"
	"github.com/sfcompute/sfc-go/models/components"
	"github.com/sfcompute/sfc-go/optionalnullable"
	"log"
)

func main() {
    ctx := context.Background()

    s := sfc.New(
        sfc.WithSecurity(os.Getenv("SFC_BEARER_AUTH")),
    )

    res, err := s.Orders.Create(ctx, components.V2CreateOrderRequest{
        Capacity: "cap_k3R-nX9vLm7Qp2Yw5Jd8F",
        Side: components.SideSell,
        Delta: components.CreateV2CreateOrderRequestDeltaUnionV2CreateOrderRequestDeltaRectangle(
            components.V2CreateOrderRequestDeltaRectangle{
                NodeCount: 387888,
                StartAt: 1738972800,
                EndAt: 1738972800,
            },
        ),
        LimitPriceDollarsPerNodeHour: "2.500000",
    }, optionalnullable.From[string](nil))
    if err != nil {
        log.Fatal(err)
    }
    if res.V2OrderResponse != nil {
        switch res.V2OrderResponse.Capacity.Type {
            case components.ExpandableCapacityIDCapacitySummaryUnionTypeStr:
                // res.V2OrderResponse.Capacity.Str is populated
            case components.ExpandableCapacityIDCapacitySummaryUnionTypeExpandableCapacityIDCapacitySummary:
                // res.V2OrderResponse.Capacity.ExpandableCapacityIDCapacitySummary is populated
            default:
                // Unknown type - use res.V2OrderResponse.Capacity.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                                                                                                                   | Type                                                                                                                                                        | Required                                                                                                                                                    | Description                                                                                                                                                 |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                                                       | :heavy_check_mark:                                                                                                                                          | The context to use for the request.                                                                                                                         |
| `body`                                                                                                                                                      | [components.V2CreateOrderRequest](../../models/components/v2createorderrequest.md)                                                                          | :heavy_check_mark:                                                                                                                                          | N/A                                                                                                                                                         |
| `idempotencyKey`                                                                                                                                            | optionalnullable.OptionalNullable[`string`]                                                                                                                 | :heavy_minus_sign:                                                                                                                                          | Unique key to ensure idempotent order creation. If provided, duplicate requests with the same key will not place a new order and return the original order. |
| `opts`                                                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                                                    | :heavy_minus_sign:                                                                                                                                          | The options for this request.                                                                                                                               |

### Response

**[*operations.CreateOrderResponse](../../models/operations/createorderresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.PaymentRequiredError     | 402                                | application/json                   |
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.ServiceUnavailableError  | 503                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Fetch

Retrieve an order by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch_order" method="get" path="/v2/orders/{id}" -->
```go
package main

import(
	"context"
	"os"
	sfc "github.com/sfcompute/sfc-go"
	"log"
	"github.com/sfcompute/sfc-go/models/components"
)

func main() {
    ctx := context.Background()

    s := sfc.New(
        sfc.WithSecurity(os.Getenv("SFC_BEARER_AUTH")),
    )

    res, err := s.Orders.Fetch(ctx, "ordr_xyz789", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.V2OrderResponse != nil {
        switch res.V2OrderResponse.Capacity.Type {
            case components.ExpandableCapacityIDCapacitySummaryUnionTypeStr:
                // res.V2OrderResponse.Capacity.Str is populated
            case components.ExpandableCapacityIDCapacitySummaryUnionTypeExpandableCapacityIDCapacitySummary:
                // res.V2OrderResponse.Capacity.ExpandableCapacityIDCapacitySummary is populated
            default:
                // Unknown type - use res.V2OrderResponse.Capacity.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | Order ID                                                 | ordr_k3R-nX9vLm7Qp2Yw5Jd8F                               |
| `expand`                                                 | []`string`                                               | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchOrderResponse](../../models/operations/fetchorderresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Cancel

Request cancellation of an order. This is asynchronous — poll `GET /v2/orders/{id}` to confirm the status has changed to cancelled.

### Example Usage

<!-- UsageSnippet language="go" operationID="cancel_order" method="delete" path="/v2/orders/{id}" -->
```go
package main

import(
	"context"
	"os"
	sfc "github.com/sfcompute/sfc-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sfc.New(
        sfc.WithSecurity(os.Getenv("SFC_BEARER_AUTH")),
    )

    res, err := s.Orders.Cancel(ctx, "ordr_xyz789")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | Order ID                                                 | ordr_k3R-nX9vLm7Qp2Yw5Jd8F                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.CancelOrderResponse](../../models/operations/cancelorderresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |