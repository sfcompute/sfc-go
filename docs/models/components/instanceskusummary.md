# InstanceSkuSummary

Summary view of an instance SKU embedded on responses that reference one (orders, procurements, instances, capacity transfers). Carries both the id and the human-readable name. Legacy SKUs whose `name` column hasn't been backfilled use `UNKNOWN_INSTANCE_SKU_NAME` as a placeholder so this field is always populated on the wire.


## Fields

| Field                      | Type                       | Required                   | Description                | Example                    |
| -------------------------- | -------------------------- | -------------------------- | -------------------------- | -------------------------- |
| `ID`                       | `string`                   | :heavy_check_mark:         | N/A                        | isku_k3R-nX9vLm7Qp2Yw5Jd8F |
| `Name`                     | `string`                   | :heavy_check_mark:         | N/A                        | my-resource-name           |