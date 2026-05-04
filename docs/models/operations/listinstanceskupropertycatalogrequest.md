# ListInstanceSkuPropertyCatalogRequest


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Limit`                                                                    | `*int64`                                                                   | :heavy_minus_sign:                                                         | Maximum number of results to return (1-200, default 50).                   |
| `StartingAfter`                                                            | `*string`                                                                  | :heavy_minus_sign:                                                         | Cursor for forward pagination (from a previous response's `cursor` field). |
| `EndingBefore`                                                             | `*string`                                                                  | :heavy_minus_sign:                                                         | Cursor for backward pagination.                                            |