# InstanceSKUs

## Overview

### Available Operations

* [ListAvailability](#listavailability) - List grouped instance-SKU availability
* [ListInstanceSKUs](#listinstanceskus) - List instance SKUs
* [GetInstanceSku](#getinstancesku) - Get instance SKU

## ListAvailability

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

Aggregate availability across instance SKUs that match `requirements`, grouped by the given property keys. Each group exposes a summed `total` schedule plus a per-SKU breakdown.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_availability" method="get" path="/preview/v2/availability" -->
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

    res, err := s.InstanceSKUs.ListAvailability(ctx, nil, []string{
        "my-resource-name",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAvailabilityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |
| `requirements`                                                                                                                                                         | `*string`                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                     | Filter SKUs before grouping. Semicolon-separated `key:value[,value...]` pairs (e.g. `accelerator:H100`). Use keys and values from `/v2/instance_sku_property_catalog`. |
| `groupBy`                                                                                                                                                              | []`string`                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                     | Property keys to group by. Repeatable: `?group_by=accelerator`. Each key must be a public registry key. Empty `group_by` → a single aggregate group.                   |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |

### Response

**[*operations.ListAvailabilityResponse](../../models/operations/listavailabilityresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## ListInstanceSKUs

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

List the instance SKUs available on the market with their properties.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_instance_skus" method="get" path="/preview/v2/instance_skus" -->
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

    res, err := s.InstanceSKUs.ListInstanceSKUs(ctx, sfc.Pointer[int64](50), sfc.Pointer("iskuc_gqXR7s0Kj5mHvE2wNpLc4Q"), sfc.Pointer("iskuc_gqXR7s0Kj5mHvE2wNpLc4Q"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListInstanceSkusResponse != nil {
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |                                                                            |
| `limit`                                                                    | `*int64`                                                                   | :heavy_minus_sign:                                                         | Maximum number of results to return (1-200, default 50).                   |                                                                            |
| `startingAfter`                                                            | `*string`                                                                  | :heavy_minus_sign:                                                         | Cursor for forward pagination (from a previous response's `cursor` field). | iskuc_gqXR7s0Kj5mHvE2wNpLc4Q                                               |
| `endingBefore`                                                             | `*string`                                                                  | :heavy_minus_sign:                                                         | Cursor for backward pagination.                                            | iskuc_gqXR7s0Kj5mHvE2wNpLc4Q                                               |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |                                                                            |

### Response

**[*operations.ListInstanceSkusResponse](../../models/operations/listinstanceskusresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## GetInstanceSku

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

Retrieve an instance SKU by ID, including its properties.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_instance_sku" method="get" path="/preview/v2/instance_skus/{id}" -->
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

    res, err := s.InstanceSKUs.GetInstanceSku(ctx, "isku_k3R-nX9vLm7Qp2Yw5Jd8F")
    if err != nil {
        log.Fatal(err)
    }
    if res.InstanceSku != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | Instance SKU ID                                          | isku_k3R-nX9vLm7Qp2Yw5Jd8F                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetInstanceSkuResponse](../../models/operations/getinstanceskuresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |