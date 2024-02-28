# SOLIDserver Resource Provider

The SOLIDserver Resource Provider lets you manage EfficientIP's SOLIDserver.

This ia a bridged provider is built on EfficientIP's Terraform provider  
https://github.com/EfficientIP-Labs/terraform-provider-solidserver

## Installing

This package is available for Go. Python is coming soon.
Nodejs, .NET and Java are not supported at this time.

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/athenawolfe/pulumi-solidserver/sdk/go/...
```

## Configuration

The following configuration points are available for the `solidserver` provider:

- `solidserver:host` (environment: `SOLIDServer_HOST`)     - SOLIDserver's URL - **required**
- `solidserver:username` (environment: `SOLIDServer_USERNAME`) - Login Username - **required**
- `solidserver:password` (environment: `SOLIDServer_PASSWORD`) - Login Password - **required**
- `solidserver:sslverify` (environment: `SOLIDServer_SSLVERIFY`) -  - True/False

