<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	sfc "github.com/sfcompute/sfc-go"
	"github.com/sfcompute/sfc-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sfc.New(
		sfc.WithSecurity(os.Getenv("SFC_BEARER_AUTH")),
	)

	res, err := s.Capacities.List(ctx, operations.ListCapacitiesRequest{
		Workspace: "wksp_k3R-nX9vLm7Qp2Yw5Jd8F",
		ID: []string{
			"cap_k3R-nX9vLm7Qp2Yw5Jd8F",
		},
		StartingAfter: sfc.Pointer("capc_gqXR7s0Kj5mHvE2wNpLc4Q"),
		EndingBefore:  sfc.Pointer("capc_gqXR7s0Kj5mHvE2wNpLc4Q"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ListCapacitiesResponse != nil {
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
<!-- End SDK Example Usage [usage] -->