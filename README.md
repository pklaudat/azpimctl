# AzPimCtl

AzPimCtl is a lightweight Go CLI tool designed to manage resource eligible roles in Azure using the Privileged Identity Management (PIM) API. It provides a simple interface for activating and listing roles, streamlining development activities that require elevated permissions in Azure. This is designed to be used with Azure CLI, as it leverages the existing authentication context.

## Features

- List eligible roles for your Azure resources
- Activate roles on-demand via the PIM API
- Simplifies role management for developers

## Installation

```sh
go install github.com/your-org/azpimctl@latest
```

## Usage

```sh
azpimctl list
azpimctl activate --role <roleName> --resource <resourceId>
```

## Requirements

- Go 1.18+
- Azure CLI authenticated session

## Configuration

AzPimCtl uses your current Azure CLI authentication context. Ensure you are logged in with:

```sh
az login
```

## License

MIT

## Disclaimer

This tool is not an official Microsoft product. Use at your own risk.