# ListInstanceSkusResponse

Paginated list of instance SKUs returned by `GET /v2/instance_skus`.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Object`                                                           | `*string`                                                          | :heavy_minus_sign:                                                 | N/A                                                                |                                                                    |
| `Cursor`                                                           | optionalnullable.OptionalNullable[`string`]                        | :heavy_minus_sign:                                                 | N/A                                                                | iskuc_gqXR7s0Kj5mHvE2wNpLc4Q                                       |
| `HasMore`                                                          | `bool`                                                             | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |
| `Data`                                                             | [][components.InstanceSku](../../models/components/instancesku.md) | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |