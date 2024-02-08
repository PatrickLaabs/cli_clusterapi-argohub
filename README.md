# ArgoHub - GitOps with batteries included

[![Go Reference](https://pkg.go.dev/badge/github.com/PatrickLaabs/cli_clusterapi-argohub.svg)](https://pkg.go.dev/github.com/PatrickLaabs/cli_clusterapi-argohub)
[![Go Report Card](https://goreportcard.com/badge/github.com/PatrickLaabs/cli_clusterapi-argohub)](https://goreportcard.com/badge/github.com/PatrickLaabs/cli_clusterapi-argohub)

## License
[![Licence](https://img.shields.io/github/license/Ileriayo/markdown-badges?style=for-the-badge)](./LICENSE)

## Goals

- Enable a user to easily provision the ArgoHub instance on kind
- Bootstrap the GitOps repository, with the given credentials
- CLI Tool
- Built-in Server to bootstrap everything

## Getting started

- docker
- kind

```

Usage:
  argohub bootstrap [flags]
  argohub bootstrap [command]

Available Commands:
  capd        capd
  capv        capv
  capz        capz

```

## Supported Providers

- vCluster
- CAPD

### Under Development
- vSphere
- Azure
- Google
