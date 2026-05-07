# FunctionResourceDefinition

A function resource


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `CreatedAt`                                        | [time.Time](https://pkg.go.dev/time#Time)          | :heavy_check_mark:                                 | The creation date of the resource.                 |
| `DeploymentID`                                     | `string`                                           | :heavy_check_mark:                                 | The ID of the deployment                           |
| `Description`                                      | `string`                                           | :heavy_check_mark:                                 | Description of the resource                        |
| `FunctionID`                                       | `string`                                           | :heavy_check_mark:                                 | The ID of the function                             |
| `ID`                                               | `string`                                           | :heavy_check_mark:                                 | The ID of the resource                             |
| `Meta`                                             | map[string]`any`                                   | :heavy_minus_sign:                                 | Meta tags for the tool                             |
| `MimeType`                                         | `*string`                                          | :heavy_minus_sign:                                 | Optional MIME type of the resource                 |
| `Name`                                             | `string`                                           | :heavy_check_mark:                                 | The name of the resource                           |
| `ProjectID`                                        | `string`                                           | :heavy_check_mark:                                 | The ID of the project                              |
| `ResourceUrn`                                      | `string`                                           | :heavy_check_mark:                                 | The URN of this resource                           |
| `Runtime`                                          | `string`                                           | :heavy_check_mark:                                 | Runtime environment (e.g., nodejs:24, python:3.12) |
| `Title`                                            | `*string`                                          | :heavy_minus_sign:                                 | Optional title for the resource                    |
| `UpdatedAt`                                        | [time.Time](https://pkg.go.dev/time#Time)          | :heavy_check_mark:                                 | The last update date of the resource.              |
| `URI`                                              | `string`                                           | :heavy_check_mark:                                 | The URI of the resource                            |
| `Variables`                                        | `any`                                              | :heavy_minus_sign:                                 | Variables configuration for the resource           |