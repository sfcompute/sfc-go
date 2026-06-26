# V2OrderPrincipal

The principal that placed an order. `type` distinguishes the two cases: a user carries `email` (when an email identity exists); a token carries `name`. The inapplicable field is omitted.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ID`                                                                               | `string`                                                                           | :heavy_check_mark:                                                                 | `users.id` for a user principal, `api_tokens.id` for a token principal.            |
| `Type`                                                                             | [components.V2OrderPrincipalKind](../../models/components/v2orderprincipalkind.md) | :heavy_check_mark:                                                                 | Whether an order's `created_by` principal is a human user or an API token.         |
| `Email`                                                                            | optionalnullable.OptionalNullable[`string`]                                        | :heavy_minus_sign:                                                                 | Email of the user, when the principal is a user with a recorded email.             |
| `Name`                                                                             | optionalnullable.OptionalNullable[`string`]                                        | :heavy_minus_sign:                                                                 | Name of the API token, when the principal is a token.                              |