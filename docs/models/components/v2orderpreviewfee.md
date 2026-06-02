# V2OrderPreviewFee

Fee charged on fills for the targeted instance SKU.

Total fee = `flat_dollars_per_node_hour × node_count × duration_hours` + `percentage_bps / 10000 × execution_total_dollars`.

Recomputed against the realized price at fill time.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `FlatDollarsPerNodeHour`                                       | `string`                                                       | :heavy_check_mark:                                             | Price rate in dollars per node-hour.                           | 2.500000                                                       |
| `PercentageBps`                                                | `int`                                                          | :heavy_check_mark:                                             | Percentage of execution total, in basis points (10000 = 100%). |                                                                |