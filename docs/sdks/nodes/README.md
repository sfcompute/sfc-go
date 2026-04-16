# Nodes

## Overview

Spin up nodes in a capacity to use your available compute.

### Available Operations

* [List](#list) - List nodes
* [Create](#create) - Create node
* [Fetch](#fetch) - Get node
* [Delete](#delete) - Delete node
* [GetLogsForNode](#getlogsfornode) - Get node logs
* [GetSSHInfoForNode](#getsshinfofornode) - Get node SSH info
* [TerminateNode](#terminatenode) - Terminate node

## List

List all nodes.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_nodes" method="get" path="/v2/nodes" -->
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

    res, err := s.Nodes.List(ctx, operations.ListNodesRequest{
        Workspace: "wksp_k3R-nX9vLm7Qp2Yw5Jd8F",
        ID: []string{
            "n",
            "o",
            "d",
            "e",
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
    if res.ListNodesResponse != nil {
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
| `request`                                                                  | [operations.ListNodesRequest](../../models/operations/listnodesrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.ListNodesResponse](../../models/operations/listnodesresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Create

Create a node.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_node" method="post" path="/v2/nodes" -->
```go
package main

import(
	"context"
	"os"
	sfc "github.com/sfcompute/sfc-go"
	"github.com/sfcompute/sfc-go/optionalnullable"
	"github.com/sfcompute/sfc-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sfc.New(
        sfc.WithSecurity(os.Getenv("SFC_BEARER_AUTH")),
    )

    res, err := s.Nodes.Create(ctx, components.CreateNodeRequest{
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
    if res.NodeResponse != nil {
        switch res.NodeResponse.Capacity.Type {
            case components.ExpandableCapacityIDCapacitySummaryUnionTypeStr:
                // res.NodeResponse.Capacity.Str is populated
            case components.ExpandableCapacityIDCapacitySummaryUnionTypeExpandableCapacityIDCapacitySummary:
                // res.NodeResponse.Capacity.ExpandableCapacityIDCapacitySummary is populated
            default:
                // Unknown type - use res.NodeResponse.Capacity.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.CreateNodeRequest](../../models/components/createnoderequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.CreateNodeResponse](../../models/operations/createnoderesponse.md), error**

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

Retrieve a node by ID or name.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch_node" method="get" path="/v2/nodes/{id}" -->
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

    res, err := s.Nodes.Fetch(ctx, "node_k3R-nX9vLm7Qp2Yw5Jd8F", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.NodeResponse != nil {
        switch res.NodeResponse.Capacity.Type {
            case components.ExpandableCapacityIDCapacitySummaryUnionTypeStr:
                // res.NodeResponse.Capacity.Str is populated
            case components.ExpandableCapacityIDCapacitySummaryUnionTypeExpandableCapacityIDCapacitySummary:
                // res.NodeResponse.Capacity.ExpandableCapacityIDCapacitySummary is populated
            default:
                // Unknown type - use res.NodeResponse.Capacity.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |                                                                            |
| `id`                                                                       | `string`                                                                   | :heavy_check_mark:                                                         | N/A                                                                        | node_k3R-nX9vLm7Qp2Yw5Jd8F                                                 |
| `expand`                                                                   | [][operations.FetchNodeExpand](../../models/operations/fetchnodeexpand.md) | :heavy_minus_sign:                                                         | N/A                                                                        |                                                                            |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |                                                                            |

### Response

**[*operations.FetchNodeResponse](../../models/operations/fetchnoderesponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete a node.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_node" method="delete" path="/v2/nodes/{id}" -->
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

    res, err := s.Nodes.Delete(ctx, "node_k3R-nX9vLm7Qp2Yw5Jd8F")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | node_k3R-nX9vLm7Qp2Yw5Jd8F                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteNodeResponse](../../models/operations/deletenoderesponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.ConflictError       | 409                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetLogsForNode

Retrieve logs for a node.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_node_logs" method="get" path="/v2/nodes/{id}/logs" -->
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

    res, err := s.Nodes.GetLogsForNode(ctx, operations.GetNodeLogsRequest{
        ID: "node_k3R-nX9vLm7Qp2Yw5Jd8F",
        RealtimeTimestampBefore: sfc.Pointer[int64](1738972800),
        RealtimeTimestampAfter: sfc.Pointer[int64](1738972800),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.NodeLogsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetNodeLogsRequest](../../models/operations/getnodelogsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GetNodeLogsResponse](../../models/operations/getnodelogsresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## GetSSHInfoForNode

Retrieve SSH connection details for a node.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_node_ssh" method="get" path="/v2/nodes/{id}/ssh" -->
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

    res, err := s.Nodes.GetSSHInfoForNode(ctx, "node_k3R-nX9vLm7Qp2Yw5Jd8F")
    if err != nil {
        log.Fatal(err)
    }
    if res.NodeSSHInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | node_k3R-nX9vLm7Qp2Yw5Jd8F                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetNodeSSHResponse](../../models/operations/getnodesshresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## TerminateNode

Terminates a running node. Terminated nodes can not be restarted.

### Example Usage

<!-- UsageSnippet language="go" operationID="terminate_node" method="post" path="/v2/nodes/{id}/terminate" -->
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

    res, err := s.Nodes.TerminateNode(ctx, "node_k3R-nX9vLm7Qp2Yw5Jd8F")
    if err != nil {
        log.Fatal(err)
    }
    if res.NodeResponse != nil {
        switch res.NodeResponse.Capacity.Type {
            case components.ExpandableCapacityIDCapacitySummaryUnionTypeStr:
                // res.NodeResponse.Capacity.Str is populated
            case components.ExpandableCapacityIDCapacitySummaryUnionTypeExpandableCapacityIDCapacitySummary:
                // res.NodeResponse.Capacity.ExpandableCapacityIDCapacitySummary is populated
            default:
                // Unknown type - use res.NodeResponse.Capacity.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | node_k3R-nX9vLm7Qp2Yw5Jd8F                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.TerminateNodeResponse](../../models/operations/terminatenoderesponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |