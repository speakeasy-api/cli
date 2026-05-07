# ExternalMCPHeaderDefinition


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Description`                                                  | `*string`                                                      | :heavy_minus_sign:                                             | Description of the header                                      |
| `HeaderName`                                                   | `string`                                                       | :heavy_check_mark:                                             | The actual HTTP header name to send (e.g., X-Api-Key)          |
| `Name`                                                         | `string`                                                       | :heavy_check_mark:                                             | The prefixed environment variable name (e.g., SLACK_X_API_KEY) |
| `Placeholder`                                                  | `*string`                                                      | :heavy_minus_sign:                                             | Placeholder value for the header                               |
| `Required`                                                     | `bool`                                                         | :heavy_check_mark:                                             | Whether the header is required                                 |
| `Secret`                                                       | `bool`                                                         | :heavy_check_mark:                                             | Whether the header value is secret                             |