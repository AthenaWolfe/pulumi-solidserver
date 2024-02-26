// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/AthenaWolfe/pulumi-solidserver/sdk/go/solidserver/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// PEM formatted file with additional certificates to trust for TLS connection
func GetAdditionalTrustCertsFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "solidserver:additionalTrustCertsFile")
}

// SOLIDServer Hostname or IP address
func GetHost(ctx *pulumi.Context) string {
	return config.Get(ctx, "solidserver:host")
}

// SOLIDServer API user's password
func GetPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "solidserver:password")
}

// SOLIDServer Version in case API user does not have admin permissions
func GetSolidserverversion(ctx *pulumi.Context) string {
	return config.Get(ctx, "solidserver:solidserverversion")
}

// Enable/Disable ssl verify (Default : enabled)
func GetSslverify(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "solidserver:sslverify")
}

// SOLIDServer API user's ID
func GetUsername(ctx *pulumi.Context) string {
	return config.Get(ctx, "solidserver:username")
}
