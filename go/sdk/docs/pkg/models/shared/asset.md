# Asset


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `ContentLength`                                   | `int64`                                           | :heavy_check_mark:                                | The content length of the asset                   |
| `ContentType`                                     | `string`                                          | :heavy_check_mark:                                | The content type of the asset                     |
| `CreatedAt`                                       | [time.Time](https://pkg.go.dev/time#Time)         | :heavy_check_mark:                                | The creation date of the asset.                   |
| `ID`                                              | `string`                                          | :heavy_check_mark:                                | The ID of the asset                               |
| `Kind`                                            | [shared.Kind](../../../pkg/models/shared/kind.md) | :heavy_check_mark:                                | N/A                                               |
| `Sha256`                                          | `string`                                          | :heavy_check_mark:                                | The SHA256 hash of the asset                      |
| `UpdatedAt`                                       | [time.Time](https://pkg.go.dev/time#Time)         | :heavy_check_mark:                                | The last update date of the asset.                |