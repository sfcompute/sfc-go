# ImageDownloadResponse


## Fields

| Field                       | Type                        | Required                    | Description                 | Example                     |
| --------------------------- | --------------------------- | --------------------------- | --------------------------- | --------------------------- |
| `URL`                       | `string`                    | :heavy_check_mark:          | Presigned download URL.     |                             |
| `ExpiresAt`                 | `int64`                     | :heavy_check_mark:          | Unix timestamp.             | 1738972800                  |
| `Sha256`                    | `string`                    | :heavy_check_mark:          | For integrity verification. | e3b0c44298fc1c149af...      |
| `Size`                      | `int64`                     | :heavy_check_mark:          | Image size in bytes.        |                             |