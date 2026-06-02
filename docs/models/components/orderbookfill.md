# OrderbookFill

A single trade: execution rate, node count, and the time it was recorded.


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    | Example                                        |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `DollarsPerNodeHour`                           | `string`                                       | :heavy_check_mark:                             | Price rate in dollars per node-hour.           | 2.500000                                       |
| `NodeCount`                                    | `int64`                                        | :heavy_check_mark:                             | Number of nodes filled at this execution rate. | 3                                              |
| `FilledAt`                                     | `int64`                                        | :heavy_check_mark:                             | Unix timestamp.                                | 1738972800                                     |