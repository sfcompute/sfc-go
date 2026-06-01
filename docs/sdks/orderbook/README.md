# Orderbook

## Overview

Read-only orderbook visibility: bid/ask spread, depth of book, and historical fills, keyed on hardware requirements + delivery window.

### Available Operations

* [GetOrderbookDepth](#getorderbookdepth) - Get market depth
* [GetOrderbookQuote](#getorderbookquote) - Get market quote
* [List](#list) - List market windows

## GetOrderbookDepth

Depth of book for the given requirements and delivery window, aggregated by price level. Individual orders, participants, and matched SKU identities are not exposed.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_orderbook_depth" method="get" path="/preview/v2/orderbook/depth" -->
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

    res, err := s.Orderbook.GetOrderbookDepth(ctx, 1746057600, 1746662400, sfc.Pointer("accelerator:H100"), sfc.Pointer[int64](20))
    if err != nil {
        log.Fatal(err)
    }
    if res.OrderbookDepthResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |                                                                    |
| `startAt`                                                          | `int64`                                                            | :heavy_check_mark:                                                 | Start of the delivery window.                                      | 1738972800                                                         |
| `endAt`                                                            | `int64`                                                            | :heavy_check_mark:                                                 | End of the delivery window.                                        | 1738972800                                                         |
| `requirements`                                                     | `*string`                                                          | :heavy_minus_sign:                                                 | URL-safe `field[:op]:value` triples joined by `;`.                 | accelerator:H100                                                   |
| `depth`                                                            | `*int64`                                                           | :heavy_minus_sign:                                                 | Maximum levels returned per side. Clamped to [1, 100]. Default 20. |                                                                    |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |                                                                    |

### Response

**[*operations.GetOrderbookDepthResponse](../../models/operations/getorderbookdepthresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## GetOrderbookQuote

Top-of-book quote (best bid + best ask) for the given requirements and delivery window. The book is aggregated across every SKU whose orders satisfy the requirements and that the caller is permitted to see.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_orderbook_quote" method="get" path="/preview/v2/orderbook/quote" -->
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

    res, err := s.Orderbook.GetOrderbookQuote(ctx, 1746057600, 1746662400, sfc.Pointer("accelerator:H100"))
    if err != nil {
        log.Fatal(err)
    }
    if res.OrderbookQuoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `startAt`                                                | `int64`                                                  | :heavy_check_mark:                                       | Start of the delivery window.                            | 1738972800                                               |
| `endAt`                                                  | `int64`                                                  | :heavy_check_mark:                                       | End of the delivery window.                              | 1738972800                                               |
| `requirements`                                           | `*string`                                                | :heavy_minus_sign:                                       | URL-safe `field[:op]:value` triples joined by `;`.       | accelerator:H100                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetOrderbookQuoteResponse](../../models/operations/getorderbookquoteresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## List

List every delivery window with resting orders matching the requirements, within the given time range. Each row is a summary; use /quote or /depth for detail on a specific window.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_orderbook_windows" method="get" path="/preview/v2/orderbook/windows" -->
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

    res, err := s.Orderbook.List(ctx, operations.ListOrderbookWindowsRequest{
        Requirements: sfc.Pointer("accelerator:H100"),
        RangeStartAt: 1746057600,
        RangeEndAt: 1748649600,
        StartingAfter: sfc.Pointer("mwin_gqXR7s0Kj5mHvE2wNpLc4Q"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListOrderbookWindowsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListOrderbookWindowsRequest](../../models/operations/listorderbookwindowsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListOrderbookWindowsResponse](../../models/operations/listorderbookwindowsresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |