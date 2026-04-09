# ErrorDetail


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Field`                                                 | optionalnullable.OptionalNullable[`string`]             | :heavy_minus_sign:                                      | The field that caused the error (for validation errors) |
| `Code`                                                  | `string`                                                | :heavy_check_mark:                                      | Specific error code for this detail                     |
| `Message`                                               | `string`                                                | :heavy_check_mark:                                      | Human-readable error message                            |