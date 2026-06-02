# BatchPatchInstancesRequest

Request body for `PATCH /v2/instances` (batch). Each entry in `data` applies a partial patch to one instance; instances not mentioned are untouched. All entries must succeed or none — a single failure rolls back every other entry's writes (422).

Duplicate `id` entries are not deduplicated by serde; the handler runs the patches in order, so the last write wins.


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Data`                                                                                     | [][components.BatchPatchInstanceEntry](../../models/components/batchpatchinstanceentry.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |