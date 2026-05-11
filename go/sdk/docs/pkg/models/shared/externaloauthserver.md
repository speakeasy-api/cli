# ExternalOAuthServer


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `CreatedAt`                                                     | [time.Time](https://pkg.go.dev/time#Time)                       | :heavy_check_mark:                                              | When the external OAuth server was created.                     |
| `ID`                                                            | `string`                                                        | :heavy_check_mark:                                              | The ID of the external OAuth server                             |
| `Metadata`                                                      | `any`                                                           | :heavy_check_mark:                                              | The metadata for the external OAuth server                      |
| `ProjectID`                                                     | `string`                                                        | :heavy_check_mark:                                              | The project ID this external OAuth server belongs to            |
| `Slug`                                                          | `string`                                                        | :heavy_check_mark:                                              | A short url-friendly label that uniquely identifies a resource. |
| `UpdatedAt`                                                     | [time.Time](https://pkg.go.dev/time#Time)                       | :heavy_check_mark:                                              | When the external OAuth server was last updated.                |