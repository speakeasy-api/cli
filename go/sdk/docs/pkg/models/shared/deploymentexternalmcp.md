# DeploymentExternalMCP


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ID`                                                                  | `string`                                                              | :heavy_check_mark:                                                    | The ID of the deployment external MCP record.                         |
| `Name`                                                                | `string`                                                              | :heavy_check_mark:                                                    | The display name for the external MCP server.                         |
| `OrganizationMcpCollectionRegistryID`                                 | `*string`                                                             | :heavy_minus_sign:                                                    | The ID of the internal collection registry the server is from.        |
| `RegistryID`                                                          | `*string`                                                             | :heavy_minus_sign:                                                    | The ID of the external MCP registry the server is from.               |
| `RegistryServerSpecifier`                                             | `string`                                                              | :heavy_check_mark:                                                    | The canonical server name used to look up the server in the registry. |
| `Slug`                                                                | `string`                                                              | :heavy_check_mark:                                                    | A short url-friendly label that uniquely identifies a resource.       |