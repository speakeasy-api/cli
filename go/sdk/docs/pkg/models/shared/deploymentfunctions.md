# DeploymentFunctions


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `AssetID`                                                       | `string`                                                        | :heavy_check_mark:                                              | The ID of the uploaded asset.                                   |
| `ID`                                                            | `string`                                                        | :heavy_check_mark:                                              | The ID of the deployment asset.                                 |
| `MemoryMib`                                                     | `*int`                                                          | :heavy_minus_sign:                                              | The memory limit in MiB of function runner machines.            |
| `Name`                                                          | `string`                                                        | :heavy_check_mark:                                              | The name to give the document as it will be displayed in UIs.   |
| `Runtime`                                                       | `string`                                                        | :heavy_check_mark:                                              | The runtime to use when executing functions.                    |
| `Scale`                                                         | `*int`                                                          | :heavy_minus_sign:                                              | The number of instances to run for the function.                |
| `Slug`                                                          | `string`                                                        | :heavy_check_mark:                                              | A short url-friendly label that uniquely identifies a resource. |