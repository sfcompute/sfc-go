# PriceLevel

One price level in the order book — orders sharing the same per-node-hour rate are summed into a single entry.


## Fields

| Field                                | Type                                 | Required                             | Description                          | Example                              |
| ------------------------------------ | ------------------------------------ | ------------------------------------ | ------------------------------------ | ------------------------------------ |
| `DollarsPerNodeHour`                 | `string`                             | :heavy_check_mark:                   | Price rate in dollars per node-hour. | 2.500000                             |
| `NodeCount`                          | `int64`                              | :heavy_check_mark:                   | Total nodes resting at this rate.    | 4                                    |