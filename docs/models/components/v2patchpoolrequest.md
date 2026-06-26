# V2PatchPoolRequest

Update a pool. Omitted fields are left unchanged.


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Name`                                                 | optionalnullable.OptionalNullable[`string`]            | :heavy_minus_sign:                                     | N/A                                                    | my-resource-name                                       |
| `Tags`                                                 | optionalnullable.OptionalNullable[map[string]`string`] | :heavy_minus_sign:                                     | N/A                                                    | {<br/>"env": "prod",<br/>"team": "infra"<br/>}         |