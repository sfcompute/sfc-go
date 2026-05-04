# Images

## Overview

Custom machine images for instances.

### Available Operations

* [List](#list) - List images
* [StartUpload](#startupload) - Create image
* [Fetch](#fetch) - Get image
* [Delete](#delete) - Delete image
* [CompleteUpload](#completeupload) - Complete image upload
* [Download](#download) - Download image
* [UploadPart](#uploadpart) - Get upload part URL

## List

List all images owned by the authenticated user.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_images" method="get" path="/preview/v2/images" -->
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

    res, err := s.Images.List(ctx, operations.ListImagesRequest{
        Workspace: "wksp_k3R-nX9vLm7Qp2Yw5Jd8F",
        ID: []string{
            "image_k3R-nX9vLm7Qp2Yw5Jd8F",
        },
        StartingAfter: sfc.Pointer("imagec_gqXR7s0Kj5mHvE2wNpLc4Q"),
        EndingBefore: sfc.Pointer("imagec_gqXR7s0Kj5mHvE2wNpLc4Q"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListImagesResponse != nil {
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
| `request`                                                                    | [operations.ListImagesRequest](../../models/operations/listimagesrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.ListImagesResponse](../../models/operations/listimagesresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.BadRequestError     | 400                           | application/json              |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## StartUpload

Create an image and start a multipart upload.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_image" method="post" path="/preview/v2/images" -->
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

    res, err := s.Images.StartUpload(ctx, components.StartUploadRequest{
        Name: optionalnullable.From(sfc.Pointer("my-resource-name")),
        Workspace: "wksp_k3R-nX9vLm7Qp2Yw5Jd8F",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ImageUploadResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.StartUploadRequest](../../models/components/startuploadrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CreateImageResponse](../../models/operations/createimageresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.BadRequestError     | 400                           | application/json              |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.ConflictError       | 409                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Fetch

Retrieve an image by ID. Returns both user-owned and public images.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch_image" method="get" path="/preview/v2/images/{id}" -->
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

    res, err := s.Images.Fetch(ctx, "image_k3R-nX9vLm7Qp2Yw5Jd8F")
    if err != nil {
        log.Fatal(err)
    }
    if res.ImageListEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | Image ID or resource path                                | image_k3R-nX9vLm7Qp2Yw5Jd8F                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchImageResponse](../../models/operations/fetchimageresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete an image.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_image" method="delete" path="/preview/v2/images/{id}" -->
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

    res, err := s.Images.Delete(ctx, "image_k3R-nX9vLm7Qp2Yw5Jd8F")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteImageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | Image ID, name, or resource path                         | image_k3R-nX9vLm7Qp2Yw5Jd8F                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteImageResponse](../../models/operations/deleteimageresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.ConflictError       | 409                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## CompleteUpload

Finalize a multipart image upload.

### Example Usage

<!-- UsageSnippet language="go" operationID="complete_image_upload" method="post" path="/preview/v2/images/{id}/complete" -->
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

    res, err := s.Images.CompleteUpload(ctx, "image_k3R-nX9vLm7Qp2Yw5Jd8F", components.CompleteUploadRequest{
        Sha256: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ImageUploadResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          | Example                                                                              |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |                                                                                      |
| `id`                                                                                 | `string`                                                                             | :heavy_check_mark:                                                                   | Image ID, name, or resource path                                                     | image_k3R-nX9vLm7Qp2Yw5Jd8F                                                          |
| `body`                                                                               | [components.CompleteUploadRequest](../../models/components/completeuploadrequest.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |                                                                                      |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |

### Response

**[*operations.CompleteImageUploadResponse](../../models/operations/completeimageuploadresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.BadRequestError     | 400                           | application/json              |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Download

Get a presigned URL to download an image.

### Example Usage

<!-- UsageSnippet language="go" operationID="download_image" method="get" path="/preview/v2/images/{id}/download" -->
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

    res, err := s.Images.Download(ctx, "image_k3R-nX9vLm7Qp2Yw5Jd8F")
    if err != nil {
        log.Fatal(err)
    }
    if res.ImageDownloadResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | Image ID, name, or resource path                         | image_k3R-nX9vLm7Qp2Yw5Jd8F                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DownloadImageResponse](../../models/operations/downloadimageresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.ConflictError       | 409                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## UploadPart

Get a presigned URL to upload one part of a multipart image upload.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_image_upload_part_url" method="post" path="/preview/v2/images/{id}/parts" -->
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

    res, err := s.Images.UploadPart(ctx, "image_k3R-nX9vLm7Qp2Yw5Jd8F", components.UploadPartRequest{
        PartID: 797058,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UploadPartResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |                                                                              |
| `id`                                                                         | `string`                                                                     | :heavy_check_mark:                                                           | Image ID, name, or resource path                                             | image_k3R-nX9vLm7Qp2Yw5Jd8F                                                  |
| `body`                                                                       | [components.UploadPartRequest](../../models/components/uploadpartrequest.md) | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |                                                                              |

### Response

**[*operations.CreateImageUploadPartURLResponse](../../models/operations/createimageuploadparturlresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.BadRequestError     | 400                           | application/json              |
| apierrors.UnauthorizedError   | 401                           | application/json              |
| apierrors.NotFoundError       | 404                           | application/json              |
| apierrors.InternalServerError | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |