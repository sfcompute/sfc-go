# InstanceSkuEnumerationValue

One allowed value for an enumeration-typed property key.


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Name`                                                 | `string`                                               | :heavy_check_mark:                                     | N/A                                                    | my-resource-name                                       |
| `DisplayName`                                          | `string`                                               | :heavy_check_mark:                                     | Human-readable display name.                           |                                                        |
| `Description`                                          | `string`                                               | :heavy_check_mark:                                     | N/A                                                    |                                                        |
| `DocumentationLink`                                    | optionalnullable.OptionalNullable[`string`]            | :heavy_minus_sign:                                     | N/A                                                    |                                                        |
| `StableAt`                                             | `int64`                                                | :heavy_check_mark:                                     | Unix timestamp.                                        | 1738972800                                             |
| `DeprecatedAt`                                         | optionalnullable.OptionalNullable[`int64`]             | :heavy_minus_sign:                                     | N/A                                                    | 1738972800                                             |
| `DeprecationInfo`                                      | optionalnullable.OptionalNullable[`string`]            | :heavy_minus_sign:                                     | Migration guidance shown when the value is deprecated. |                                                        |