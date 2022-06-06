# Speakeasy CLI 

This is a shell repo that houses the CLI used for interacting with Speakeasy's orchestration capabilities for the API lifecycle. Under the hood it calls lanaguage specific code parsing libraries.   

## Brew Tap
This CLI is available using brew: 

```bash
brew install speakeasy-api/taps/speakeasy
```

## Creating a release
We use [goreleaser](https://goreleaser.com/) to manage the versioned releases of the CLI. To release a new version: 

* Ensure you have the right version of `speakeasy-core` libraries pinned in [dependencies](https://github.com/speakeasy-api/cli/blob/main/cmd/speakeasy/go.mod#L6) 
* Create a new tag of the CLI with `git tag vX.X.X` and `git push origin tag vX.X.X` . Add any useful descriptions or associated GH tickets that have been closed in the description
* Assuming you have `goreleaser` installed run `goreleaser release`. This will take the latest tag of the cli and update our homebrew tap to point to the latest version as well as upload built binaries for different architectures 
* (optional) If you want to sandbox the above step you may run `goreleaser release --snapshot --rm-dist` locally to check the build. If you want to check syntax in the `goreleaser` config file run `goreleaser check`.  

## Documentation
Usage documentation for the CLI powered through the `parser` repo and hosted at [docs.speakeasyapi.dev](docs.speakeasyapi.dev) 
