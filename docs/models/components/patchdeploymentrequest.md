# PatchDeploymentRequest


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 | Example                                     |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `Name`                                      | optionalnullable.OptionalNullable[`string`] | :heavy_minus_sign:                          | N/A                                         | my-resource-name                            |
| `InstanceTemplate`                          | optionalnullable.OptionalNullable[`string`] | :heavy_minus_sign:                          | N/A                                         | ntmpl_k3R-nX9vLm7Qp2Yw5Jd8F                 |
| `TargetInstanceCount`                       | optionalnullable.OptionalNullable[`int`]    | :heavy_minus_sign:                          | N/A                                         |                                             |
| `InstanceNameTemplate`                      | optionalnullable.OptionalNullable[`string`] | :heavy_minus_sign:                          | N/A                                         | my-fleet-{{nanoid(9)}}                      |