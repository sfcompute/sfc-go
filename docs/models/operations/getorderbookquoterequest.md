# GetOrderbookQuoteRequest


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Requirements`                                     | `*string`                                          | :heavy_minus_sign:                                 | URL-safe `field[:op]:value` triples joined by `;`. | accelerator:H100                                   |
| `StartAt`                                          | `int64`                                            | :heavy_check_mark:                                 | Start of the delivery window.                      | 1738972800                                         |
| `EndAt`                                            | `int64`                                            | :heavy_check_mark:                                 | End of the delivery window.                        | 1738972800                                         |