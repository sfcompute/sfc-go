# InstanceSSHHostKey


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `KeyType`                                             | `string`                                              | :heavy_check_mark:                                    | Key algorithm.                                        | **Example 1:** ssh-ed25519<br/>**Example 2:** ssh-rsa |
| `Key`                                                 | `string`                                              | :heavy_check_mark:                                    | Base64-encoded public key.                            | AAAAC3NzaC1lZDI1NTE5AAAAI...                          |