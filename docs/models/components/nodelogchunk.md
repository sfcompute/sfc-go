# NodeLogChunk


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  | Example                                      |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `TimestampRealtime`                          | `int64`                                      | :heavy_check_mark:                           | Unix timestamp.                              | 1738972800                                   |
| `TimestampMonotonicSecs`                     | `int64`                                      | :heavy_check_mark:                           | Monotonic clock seconds.                     |                                              |
| `TimestampMonotonicNanos`                    | `int64`                                      | :heavy_check_mark:                           | Nanosecond component of the monotonic clock. |                                              |
| `Seqnum`                                     | `int64`                                      | :heavy_check_mark:                           | N/A                                          |                                              |
| `Data`                                       | `string`                                     | :heavy_check_mark:                           | Base-64 encoded raw console output.          | SGVsbG8gV29ybGQK                             |