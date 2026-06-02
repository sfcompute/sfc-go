# InstanceSKUCatalog

## Overview

Browse available instance SKU property definitions.

### Available Operations

* [List](#list) - List instance SKU property catalog

## List

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

List every property key and its allowed values. Use the keys and values here when filling in `requirements` on orders and procurements.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_instance_sku_property_catalog" method="get" path="/preview/v2/instance_sku_property_catalog" -->
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

    res, err := s.InstanceSKUCatalog.List(ctx, sfc.Pointer[int64](50), nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListInstanceSkuPropertyCatalogResponse != nil {
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
| `limit`                                                                    | `*int64`                                                                   | :heavy_minus_sign:                                                         | Maximum number of results to return (1-200, default 50).                   |
| `startingAfter`                                                            | `*string`                                                                  | :heavy_minus_sign:                                                         | Cursor for forward pagination (from a previous response's `cursor` field). |
| `endingBefore`                                                             | `*string`                                                                  | :heavy_minus_sign:                                                         | Cursor for backward pagination.                                            |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.ListInstanceSkuPropertyCatalogResponse](../../models/operations/listinstanceskupropertycatalogresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |