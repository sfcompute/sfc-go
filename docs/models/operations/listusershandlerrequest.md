# ListUsersHandlerRequest


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ID`                                                       | []`string`                                                 | :heavy_minus_sign:                                         | Filter by user ID or resource path (repeatable).           |                                                            |
| `Limit`                                                    | `*int64`                                                   | :heavy_minus_sign:                                         | N/A                                                        |                                                            |
| `StartingAfter`                                            | `*string`                                                  | :heavy_minus_sign:                                         | Set to the response's `cursor` to fetch the next page.     | usrc_gqXR7s0Kj5mHvE2wNpLc4Q                                |
| `EndingBefore`                                             | `*string`                                                  | :heavy_minus_sign:                                         | Set to the response's `cursor` to fetch the previous page. | usrc_gqXR7s0Kj5mHvE2wNpLc4Q                                |