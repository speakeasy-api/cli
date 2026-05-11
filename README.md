<a href="https://www.speakeasy.com/product/mcp-gateway" target="_blank">
   <picture>
       <source media="(prefers-color-scheme: light)" srcset="https://github.com/user-attachments/assets/1812f171-1650-4045-ac35-21bdd7b103a6">
       <source media="(prefers-color-scheme: dark)" srcset="https://github.com/user-attachments/assets/3f14e446-0dec-4b8a-b36e-fd92efc25751">
       <img src="https://github.com/user-attachments/assets/3f14e446-0dec-4b8a-b36e-fd92efc25751#gh-dark-mode-only" alt="Speakeasy AI Control Plane">
   </picture>
</a>

<h3 align="center">Speakeasy AI Control Plane</h3>
<p align="center">
    <br />
    <a href="https://www.speakeasy.com/product/mcp-gateway"><strong>Learn more »</strong></a>
    <br />
    <br />
    <a href="https://speakeasy.com/"><img alt="Built by Speakeasy" src="https://www.speakeasy.com/assets/badges/built-by-speakeasy.svg" />
    <br />
  </a>
    <a href="#speakeasy-functions"><strong>Speakeasy Functions</strong></a> ·
    <a href="#cli"><strong>CLI</strong></a> ·
    <a href="#support"><strong>Support</strong></a> ·
    <a href="#contributing"><strong>Contributing</strong></a>
</p>

<hr />

# Speakeasy CLI

Public function tooling for the Speakeasy AI Control Plane.

This repository contains the TypeScript framework for authoring hosted MCP
tools, the scaffolding package for new function projects, and the Go CLI used
to authenticate, stage, push, and inspect deployments.

## Layout

- `go/cli`: Go CLI entrypoint. This depends on the generated Go SDK in `go/sdk`, not on server packages.
- `go/sdk`: Generated Go SDK for the Speakeasy Control Plane API.
- `ts/create-function`: Project scaffolding for new function projects.
- `ts/functions`: TypeScript functions framework package.

## Speakeasy Functions

Create agentic tools from simple TypeScript code using Speakeasy Functions, the
functions framework for the Speakeasy AI Control Plane. Functions can be deployed
to hosted MCP servers or run locally during development.

The fastest way to get started is with the Speakeasy function scaffolder,
which creates a complete TypeScript project with a working function.

```bash
# Scaffold a new function project and follow the prompts
npm create @gram-ai/function@latest

# Move into your newly created function directory
cd my_function

# Build and deploy
npm run build
npm run push
```

A default function is created for you.

```typescript
import { Gram } from "@gram-ai/functions";
import * as z from "zod/mini";

const functions = new Gram().tool({
  name: "add",
  description: "Add two numbers together",
  inputSchema: { a: z.number(), b: z.number() },
  async execute(ctx, input) {
    return ctx.json({ sum: input.a + input.b });
  },
});

export default functions;
```

The generated project also includes a local MCP server entrypoint so you can
test tools with MCP Inspector using `pnpm run dev`.

Common function use cases include:

- Powering in-application chat by exposing context from internal APIs or third-party APIs through tools.
- Adding data and actions to AI workflows in platforms such as Zapier and n8n.
- Creating higher-order tools that compose lower-level API calls.
- Managing and securing MCP servers for an organization through a unified control plane.

See the [TypeScript functions README](./ts/functions/README.md) for framework
details.

## CLI

The Speakeasy CLI provides scripted access to the Speakeasy AI Control
Plane. Use it to authenticate, stage function artifacts, push deployments, and
check deployment status.

Install the CLI on macOS or Linux:

```bash
curl -fsSL https://raw.githubusercontent.com/speakeasy-api/cli/main/install-cli.sh | bash
```

Install the CLI on Windows:

```powershell
iwr -useb https://raw.githubusercontent.com/speakeasy-api/cli/main/install-cli.ps1 | iex
```

Authenticate before pushing deployments:

```bash
gram auth
```

Stage and push a function artifact:

```bash
gram stage function --slug my-function --location ./dist/function.zip
gram push
```

See the [CLI README](./go/cli/README.md) for local development and testing.

## Support

- Slack: [Join our Slack](https://join.slack.com/t/speakeasy-dev/shared_invite/zt-3hudfoj4y-9EPqMmHIFhNiTtannqiV3Q) for support and discussions.
- In-app: When using the Speakeasy application, you can engage with the core maintainers of the product.
- GitHub: Contribute or report issues [on this repository](https://github.com/speakeasy-api/cli/issues/new).
- Documentation: View the [Speakeasy docs](https://www.speakeasy.com/docs) for function framework guidance.

## Contributing

Contributions are welcome. Please open an issue or discussion for questions or
suggestions before starting significant work.

For package-specific development notes, see the README files in
[`ts/functions`](./ts/functions/README.md), [`go/cli`](./go/cli/README.md), and
[`go/sdk`](./go/sdk/README.md).
