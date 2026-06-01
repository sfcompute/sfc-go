# ScheduleEntry

A `[start_at, end_at)` time range with a fixed `node_count`. `end_at` is `null` only on the final entry, marking an unbounded tail.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                | Example                                    |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `StartAt`                                  | `int64`                                    | :heavy_check_mark:                         | Unix timestamp.                            | 1738972800                                 |
| `EndAt`                                    | optionalnullable.OptionalNullable[`int64`] | :heavy_minus_sign:                         | N/A                                        | 1738972800                                 |
| `NodeCount`                                | `int`                                      | :heavy_check_mark:                         | N/A                                        |                                            |