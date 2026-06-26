# FeatureFlags

## Overview

### Available Operations

* [ListUserFeatureFlags](#listuserfeatureflags) - List feature flags scoped to the caller
* [SetFeatureFlagEnrollment](#setfeatureflagenrollment) - Enroll or unenroll the caller in a feature flag

## ListUserFeatureFlags

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

List every user-visible feature flag whose audience matches the caller, along with whether the caller is currently enrolled and whether the flag is self-enrollable via `POST /v2/feature_flags/{feature_flag}`. Flags with `user_enroll = false` are read-only here. Flags not marked `user_visible` are hidden from this listing entirely.

### Example Usage

<!-- UsageSnippet language="go" operationID="listUserFeatureFlags" method="get" path="/preview/v2/feature_flags" -->
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

    res, err := s.FeatureFlags.ListUserFeatureFlags(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListUserFeatureFlagsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListUserFeatureFlagsResponse](../../models/operations/listuserfeatureflagsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SetFeatureFlagEnrollment

> ⚠️ This endpoint is in [public preview](/preview/roadmap).

Set the caller's enrollment state for a self-enrollable feature flag.

### Example Usage

<!-- UsageSnippet language="go" operationID="setFeatureFlagEnrollment" method="post" path="/preview/v2/feature_flags/{feature_flag}" -->
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

    res, err := s.FeatureFlags.SetFeatureFlagEnrollment(ctx, "<value>", components.SetEnrollmentRequest{
        Enrolled: false,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SetEnrollmentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `featureFlag`                                                                      | `string`                                                                           | :heavy_check_mark:                                                                 | Feature flag name                                                                  |
| `body`                                                                             | [components.SetEnrollmentRequest](../../models/components/setenrollmentrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.SetFeatureFlagEnrollmentResponse](../../models/operations/setfeatureflagenrollmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |