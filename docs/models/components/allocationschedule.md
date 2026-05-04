# AllocationSchedule


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `Total`                                                                               | [][components.ScheduleEntry](../../models/components/scheduleentry.md)                | :heavy_check_mark:                                                                    | Capacity availability over time.                                                      |
| `ByInstanceSku`                                                                       | map[string][][components.ScheduleEntry](../../models/components/scheduleentry.md)     | :heavy_check_mark:                                                                    | Allocation schedule keyed by instance SKU. Only includes current and future schedule. |