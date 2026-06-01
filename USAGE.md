<!-- Start SDK Example Usage [usage] -->
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