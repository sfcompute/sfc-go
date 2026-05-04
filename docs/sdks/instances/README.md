# Instances

## Overview

Spin up instances in a capacity to use your available compute.

### Available Operations

* [List](#list) - List instances
* [Create](#create) - Create instance
* [Fetch](#fetch) - Get instance
* [Delete](#delete) - Delete instance
* [GetLogsForInstance](#getlogsforinstance) - Get instance logs
* [GetSSHInfoForInstance](#getsshinfoforinstance) - Get instance SSH info
* [TerminateInstance](#terminateinstance) - Terminate instance

## List

List all instances.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_instances" method="get" path="/preview/v2/instances" -->
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

    res, err := s.Instances.List(ctx, operations.ListInstancesRequest{
        Workspace: "wksp_k3R-nX9vLm7Qp2Yw5Jd8F",
        ID: []string{
            "i",
            "n",
            "s",
            "t",
            "_",
            "k",
            "3",
            "R",
            "-",
            "n",
            "X",
            "9",
            "v",
            "L",
            "m",
            "7",
            "Q",
            "p",
            "2",
            "Y",
            "w",
            "5",
            "J",
            "d",
            "8",
            "F",
        },
        Capacity: sfc.Pointer("cap_k3R-nX9vLm7Qp2Yw5Jd8F"),
        StartingAfter: sfc.Pointer("nodec_gqXR7s0Kj5mHvE2wNpLc4Q"),
        EndingBefore: sfc.Pointer("nodec_gqXR7s0Kj5mHvE2wNpLc4Q"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListInstancesResponse != nil {
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListInstancesRequest](../../models/operations/listinstancesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListInstancesResponse](../../models/operations/listinstancesresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Create

Create an instance.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_instance" method="post" path="/preview/v2/instances" -->
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

    res, err := s.Instances.Create(ctx, components.CreateInstanceRequest{
        Name: optionalnullable.From(sfc.Pointer("my-resource-name")),
        Capacity: "cap_k3R-nX9vLm7Qp2Yw5Jd8F",
        Image: "image_k3R-nX9vLm7Qp2Yw5Jd8F",
        CloudInitUserData: sfc.Pointer("IyEvYmluL2Jhc2gKZWNobyBoZWxsbyB3b3JsZAo="),
        Tags: optionalnullable.From(sfc.Pointer(map[string]string{
            "env": "prod",
            "team": "infra",
        })),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.InstanceResponse != nil {
        switch res.InstanceResponse.Capacity.Type {
            case components.ExpandableCapacityIDCapacitySummaryUnionTypeStr:
                // res.InstanceResponse.Capacity.Str is populated
            case components.ExpandableCapacityIDCapacitySummaryUnionTypeExpandableCapacityIDCapacitySummary:
                // res.InstanceResponse.Capacity.ExpandableCapacityIDCapacitySummary is populated
            default:
                // Unknown type - use res.InstanceResponse.Capacity.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.CreateInstanceRequest](../../models/components/createinstancerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateInstanceResponse](../../models/operations/createinstanceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Fetch

Retrieve an instance by ID or name.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch_instance" method="get" path="/preview/v2/instances/{id}" -->
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

    res, err := s.Instances.Fetch(ctx, "inst_k3R-nX9vLm7Qp2Yw5Jd8F", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.InstanceResponse != nil {
        switch res.InstanceResponse.Capacity.Type {
            case components.ExpandableCapacityIDCapacitySummaryUnionTypeStr:
                // res.InstanceResponse.Capacity.Str is populated
            case components.ExpandableCapacityIDCapacitySummaryUnionTypeExpandableCapacityIDCapacitySummary:
                // res.InstanceResponse.Capacity.ExpandableCapacityIDCapacitySummary is populated
            default:
                // Unknown type - use res.InstanceResponse.Capacity.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `id`                                                                               | `string`                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                | inst_k3R-nX9vLm7Qp2Yw5Jd8F                                                         |
| `expand`                                                                           | [][operations.FetchInstanceExpand](../../models/operations/fetchinstanceexpand.md) | :heavy_minus_sign:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.FetchInstanceResponse](../../models/operations/fetchinstanceresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete an instance.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_instance" method="delete" path="/preview/v2/instances/{id}" -->
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

    res, err := s.Instances.Delete(ctx, "inst_k3R-nX9vLm7Qp2Yw5Jd8F")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | inst_k3R-nX9vLm7Qp2Yw5Jd8F                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteInstanceResponse](../../models/operations/deleteinstanceresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.ConflictError       | 409                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetLogsForInstance

Retrieve logs for an instance.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_instance_logs" method="get" path="/preview/v2/instances/{id}/logs" -->
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

    res, err := s.Instances.GetLogsForInstance(ctx, operations.GetInstanceLogsRequest{
        ID: "inst_k3R-nX9vLm7Qp2Yw5Jd8F",
        RealtimeTimestampBefore: sfc.Pointer[int64](1738972800),
        RealtimeTimestampAfter: sfc.Pointer[int64](1738972800),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.InstanceLogsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetInstanceLogsRequest](../../models/operations/getinstancelogsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetInstanceLogsResponse](../../models/operations/getinstancelogsresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## GetSSHInfoForInstance

Retrieve SSH connection details for an instance.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_instance_ssh" method="get" path="/preview/v2/instances/{id}/ssh" -->
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

    res, err := s.Instances.GetSSHInfoForInstance(ctx, "inst_k3R-nX9vLm7Qp2Yw5Jd8F")
    if err != nil {
        log.Fatal(err)
    }
    if res.InstanceSSHInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | inst_k3R-nX9vLm7Qp2Yw5Jd8F                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetInstanceSSHResponse](../../models/operations/getinstancesshresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## TerminateInstance

Terminates a running instance. Terminated instances can not be restarted.

### Example Usage

<!-- UsageSnippet language="go" operationID="terminate_instance" method="post" path="/preview/v2/instances/{id}/terminate" -->
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

    res, err := s.Instances.TerminateInstance(ctx, "inst_k3R-nX9vLm7Qp2Yw5Jd8F")
    if err != nil {
        log.Fatal(err)
    }
    if res.InstanceResponse != nil {
        switch res.InstanceResponse.Capacity.Type {
            case components.ExpandableCapacityIDCapacitySummaryUnionTypeStr:
                // res.InstanceResponse.Capacity.Str is populated
            case components.ExpandableCapacityIDCapacitySummaryUnionTypeExpandableCapacityIDCapacitySummary:
                // res.InstanceResponse.Capacity.ExpandableCapacityIDCapacitySummary is populated
            default:
                // Unknown type - use res.InstanceResponse.Capacity.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | inst_k3R-nX9vLm7Qp2Yw5Jd8F                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.TerminateInstanceResponse](../../models/operations/terminateinstanceresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |