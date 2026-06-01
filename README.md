# sfc

Developer-friendly & type-safe Go SDK specifically catered to leverage *sfc* API.

[![Built by Speakeasy](https://img.shields.io/badge/Built_by-SPEAKEASY-374151?style=for-the-badge&labelColor=f3f4f6)](https://www.speakeasy.com/?utm_source=sfc&utm_campaign=go)
[![License: MIT](https://img.shields.io/badge/LICENSE_//_MIT-3b5bdb?style=for-the-badge&labelColor=eff6ff)](https://opensource.org/licenses/MIT)


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/the-san-francisco-compute-company/v2). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary


<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [sfc](#sfc)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Pagination](#pagination)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/sfcompute/sfc-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
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
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type | Scheme      |
| ------------ | ---- | ----------- |
| `BearerAuth` | http | HTTP Bearer |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
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
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Capacities](docs/sdks/capacities/README.md)

* [List](docs/sdks/capacities/README.md#list) - List capacities
* [Create](docs/sdks/capacities/README.md#create) - Create capacity
* [Fetch](docs/sdks/capacities/README.md#fetch) - Get capacity
* [Delete](docs/sdks/capacities/README.md#delete) - Delete capacity
* [Update](docs/sdks/capacities/README.md#update) - Update capacity
* [ListCapacityTransfers](docs/sdks/capacities/README.md#listcapacitytransfers) - List capacity transfers
* [CreateCapacityTransfer](docs/sdks/capacities/README.md#createcapacitytransfer) - Create capacity transfer
* [FetchCapacityTransfer](docs/sdks/capacities/README.md#fetchcapacitytransfer) - Get capacity transfer

### [Deployments](docs/sdks/deployments/README.md)

* [List](docs/sdks/deployments/README.md#list) - List deployments
* [Create](docs/sdks/deployments/README.md#create) - Create deployment
* [GetDeployment](docs/sdks/deployments/README.md#getdeployment) - Get deployment
* [Delete](docs/sdks/deployments/README.md#delete) - Delete deployment
* [PatchDeployment](docs/sdks/deployments/README.md#patchdeployment) - Update deployment

### [Images](docs/sdks/images/README.md)

* [List](docs/sdks/images/README.md#list) - List images
* [StartUpload](docs/sdks/images/README.md#startupload) - Create image
* [Fetch](docs/sdks/images/README.md#fetch) - Get image
* [Delete](docs/sdks/images/README.md#delete) - Delete image
* [CompleteUpload](docs/sdks/images/README.md#completeupload) - Complete image upload
* [Download](docs/sdks/images/README.md#download) - Download image
* [UploadPart](docs/sdks/images/README.md#uploadpart) - Get upload part URL

### [InstanceSKUCatalog](docs/sdks/instanceskucatalog/README.md)

* [List](docs/sdks/instanceskucatalog/README.md#list) - List instance SKU property catalog

### [InstanceSKUs](docs/sdks/instanceskus/README.md)

* [ListAvailability](docs/sdks/instanceskus/README.md#listavailability) - List grouped instance-SKU availability
* [ListInstanceSKUs](docs/sdks/instanceskus/README.md#listinstanceskus) - List instance SKUs
* [GetInstanceSku](docs/sdks/instanceskus/README.md#getinstancesku) - Get instance SKU

### [InstanceTemplates](docs/sdks/instancetemplates/README.md)

* [List](docs/sdks/instancetemplates/README.md#list) - List instance templates
* [Create](docs/sdks/instancetemplates/README.md#create) - Create instance template
* [Fetch](docs/sdks/instancetemplates/README.md#fetch) - Get instance template
* [Delete](docs/sdks/instancetemplates/README.md#delete) - Delete instance template

### [Instances](docs/sdks/instances/README.md)

* [List](docs/sdks/instances/README.md#list) - List instances
* [Create](docs/sdks/instances/README.md#create) - Create instance
* [Fetch](docs/sdks/instances/README.md#fetch) - Get instance
* [Delete](docs/sdks/instances/README.md#delete) - Delete instance
* [GetLogsForInstance](docs/sdks/instances/README.md#getlogsforinstance) - Get instance logs
* [GetSSHInfoForInstance](docs/sdks/instances/README.md#getsshinfoforinstance) - Get instance SSH info
* [TerminateInstance](docs/sdks/instances/README.md#terminateinstance) - Terminate instance

### [Orderbook](docs/sdks/orderbook/README.md)

* [GetOrderbookDepth](docs/sdks/orderbook/README.md#getorderbookdepth) - Get market depth
* [GetOrderbookQuote](docs/sdks/orderbook/README.md#getorderbookquote) - Get market quote
* [List](docs/sdks/orderbook/README.md#list) - List market windows

### [Orders](docs/sdks/orders/README.md)

* [GetOrderEstimate](docs/sdks/orders/README.md#getorderestimate) - Estimate an order
* [List](docs/sdks/orders/README.md#list) - List orders
* [Create](docs/sdks/orders/README.md#create) - Create order
* [Fetch](docs/sdks/orders/README.md#fetch) - Get order
* [Cancel](docs/sdks/orders/README.md#cancel) - Cancel order

### [Procurements](docs/sdks/procurements/README.md)

* [List](docs/sdks/procurements/README.md#list) - List procurements
* [Create](docs/sdks/procurements/README.md#create) - Create procurement
* [GetProcurement](docs/sdks/procurements/README.md#getprocurement) - Get procurement
* [Delete](docs/sdks/procurements/README.md#delete) - Delete procurement
* [PatchProcurement](docs/sdks/procurements/README.md#patchprocurement) - Update procurement

### [Workspaces](docs/sdks/workspaces/README.md)

* [List](docs/sdks/workspaces/README.md#list) - List workspaces
* [Create](docs/sdks/workspaces/README.md#create) - Create workspace
* [GetWorkspaceHandler](docs/sdks/workspaces/README.md#getworkspacehandler) - Get workspace
* [Delete](docs/sdks/workspaces/README.md#delete) - Delete workspace
* [PatchWorkspaceHandler](docs/sdks/workspaces/README.md#patchworkspacehandler) - Update workspace

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
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
<!-- End Pagination [pagination] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	sfc "github.com/sfcompute/sfc-go"
	"github.com/sfcompute/sfc-go/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := sfc.New(
		sfc.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.InstanceSKUs.ListAvailability(ctx, nil, []string{
		"my-resource-name",
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.ListAvailabilityResponse != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	sfc "github.com/sfcompute/sfc-go"
	"github.com/sfcompute/sfc-go/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := sfc.New(
		sfc.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `ListAvailability` function may return the following errors:

| Error Type                         | Status Code | Content Type     |
| ---------------------------------- | ----------- | ---------------- |
| apierrors.UnauthorizedError        | 401         | application/json |
| apierrors.UnprocessableEntityError | 422         | application/json |
| apierrors.InternalServerError      | 500         | application/json |
| apierrors.APIError                 | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	sfc "github.com/sfcompute/sfc-go"
	"github.com/sfcompute/sfc-go/models/apierrors"
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

		var e *apierrors.UnauthorizedError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.UnprocessableEntityError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.InternalServerError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	sfc "github.com/sfcompute/sfc-go"
	"log"
)

func main() {
	ctx := context.Background()

	s := sfc.New(
		sfc.WithServerURL("https://api.sfcompute.com"),
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/sfcompute/sfc-go"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sfc.New(sfc.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=sfc&utm_campaign=go)
