# BatchPatchInstancesResponse

Response shape for `PATCH /v2/instances` (batch). Mirrors the input list — one `InstanceResponse` per *unique* `id` in the request body, reflecting the post-write state. Unlike the paginated list response, there's no `cursor` or `has_more`: the response is exactly the instances the caller mentioned, no pagination involved.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Object`                                                                     | `*string`                                                                    | :heavy_minus_sign:                                                           | N/A                                                                          |
| `Data`                                                                       | [][components.InstanceResponse](../../models/components/instanceresponse.md) | :heavy_check_mark:                                                           | N/A                                                                          |