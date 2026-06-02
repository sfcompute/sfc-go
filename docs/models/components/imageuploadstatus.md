# ImageUploadStatus

## Example Usage

```go
import (
	"github.com/sfcompute/sfc-go/models/components"
)

value := components.ImageUploadStatusStarted

// Open enum: custom values can be created with a direct type cast
custom := components.ImageUploadStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `ImageUploadStatusStarted`   | started                      |
| `ImageUploadStatusUploading` | uploading                    |
| `ImageUploadStatusCompleted` | completed                    |
| `ImageUploadStatusFailed`    | failed                       |
| `ImageUploadStatusRevoked`   | revoked                      |