# AllocationSchedule


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `Total`                                                                                 | [][components.ScheduleEntry](../../models/components/scheduleentry.md)                  | :heavy_check_mark:                                                                      | Capacity availability over time.                                                        |
| `ByZone`                                                                                | map[string][][components.ScheduleEntry](../../models/components/scheduleentry.md)       | :heavy_check_mark:                                                                      | Allocation schedule on a zone-by-zone basis. Only includes current and future schedule. |