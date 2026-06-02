# Instances

## Overview

Spin up instances in a capacity to use your available compute.

### Available Operations

* [List](#list) - List instances
* [Create](#create) - Create instance
* [BatchPatchInstances](#batchpatchinstances) - Update multiple instances
* [Fetch](#fetch) - Get instance
* [Delete](#delete) - Delete instance
* [Update](#update) - Update instance
* [GetLogsForInstance](#getlogsforinstance) - Get instance logs
* [GetSSHInfoForInstance](#getsshinfoforinstance) - Get instance SSH info
* [TerminateInstance](#terminateinstance) - Terminate instance

## List

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

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
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Create

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

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
        InstanceSku: "isku_k3R-nX9vLm7Qp2Yw5Jd8F",
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
        // handle response
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

## BatchPatchInstances

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

Update one or more instances atomically. All listed instances must be in the same workspace; mixed-workspace batches are rejected with 422.

### Example Usage

<!-- UsageSnippet language="go" operationID="batch_patch_instances" method="patch" path="/preview/v2/instances" -->
```go
package main

import(
	"context"
	sfc "github.com/sfcompute/sfc-go"
	"github.com/sfcompute/sfc-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sfc.New(
        sfc.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Instances.BatchPatchInstances(ctx, components.BatchPatchInstancesRequest{
        Data: []components.BatchPatchInstanceEntry{
            components.BatchPatchInstanceEntry{
                ID: "inst_k3R-nX9vLm7Qp2Yw5Jd8F",
                Priority: 479768,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BatchPatchInstancesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.BatchPatchInstancesRequest](../../models/components/batchpatchinstancesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.BatchPatchInstancesResponse](../../models/operations/batchpatchinstancesresponse.md), error**

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

Retrieve an instance by ID or name.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch_instance" method="get" path="/preview/v2/instances/{id}" -->
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

    res, err := s.Instances.Fetch(ctx, "inst_k3R-nX9vLm7Qp2Yw5Jd8F")
    if err != nil {
        log.Fatal(err)
    }
    if res.InstanceResponse != nil {
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

**[*operations.FetchInstanceResponse](../../models/operations/fetchinstanceresponse.md), error**

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
| apierrors.ForbiddenError      | 403                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.ConflictError       | 409                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

Update an instance. Omitted fields are left unchanged.

### Example Usage

<!-- UsageSnippet language="go" operationID="update_instance" method="patch" path="/preview/v2/instances/{id}" -->
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

    res, err := s.Instances.Update(ctx, "inst_k3R-nX9vLm7Qp2Yw5Jd8F", components.PatchInstanceRequest{
        Name: optionalnullable.From(sfc.Pointer("my-resource-name")),
        Tags: optionalnullable.From(sfc.Pointer(map[string]string{
            "env": "prod",
            "team": "infra",
        })),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.InstanceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `id`                                                                               | `string`                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                | inst_k3R-nX9vLm7Qp2Yw5Jd8F                                                         |
| `body`                                                                             | [components.PatchInstanceRequest](../../models/components/patchinstancerequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.UpdateInstanceResponse](../../models/operations/updateinstanceresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.BadRequestError     | 400                           | application/json              |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.ForbiddenError      | 403                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetLogsForInstance

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

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
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## GetSSHInfoForInstance

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

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
| apierrors.ForbiddenError      | 403                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## TerminateInstance

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

Terminates a running instance. Terminated instances can not be restarted.

### Example Usage

<!-- UsageSnippet language="go" operationID="terminate_instance" method="post" path="/preview/v2/instances/{id}/terminate" -->
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

    res, err := s.Instances.TerminateInstance(ctx, "inst_k3R-nX9vLm7Qp2Yw5Jd8F")
    if err != nil {
        log.Fatal(err)
    }
    if res.InstanceResponse != nil {
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

**[*operations.TerminateInstanceResponse](../../models/operations/terminateinstanceresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.ForbiddenError      | 403                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |