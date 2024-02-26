# Foo Resource Provider

The SOLIDserver Resource Provider lets you manage EfficientIP's SOLIDserver.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @athenawolfe/solidserver
```

or `yarn`:

```bash
yarn add @athenawolfe/solidserver
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_solidserver
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/athenawolfe/pulumi-solidserver/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package AthenaWolfe.Solidserver
```

## Configuration


The following configuration points are available for the `solidserver` provider:

- `solidserver:host` (environment: `SDS_HOST`)     - SOLIDserver's URL - **required**
- `solidserver:username` (environment: `SDS_USERNAME`) - Login Username - **required**
- `solidserver:password` (environment: `SDS_PASSWORD`) - Login Password - **required**
- `solidserver:solidserverversion` (environment: `SDS_SOLIDSERVERVERION`) - The version of sds we are connecting too.
- `solidserver:sslverify` (environment: `SDS_SSLVERIFY`) -  - True/False do we need to verify the cert.


## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/solidserver/api-docs/).
