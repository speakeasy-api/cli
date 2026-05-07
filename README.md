# gf

Public Gram function tooling.

## Layout

- `go/cli`: Go CLI entrypoint. This will depend on the generated Go SDK in `go/sdk`, not on Gram server packages.
- `go/sdk`: Placeholder for the OAS-generated Go SDK.
- `ts/functions`: TypeScript functions framework package. Generated TypeScript client code for framework API calls can live under this package when needed.

