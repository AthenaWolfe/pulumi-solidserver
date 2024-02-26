// Copyright 2016-2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package solidserver

import (
	"fmt"
	"path"
	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	// Replace this provider with the provider you are bridging.
	//solidserver "github.com/iwahbe/terraform-provider-solidserver/provider"
	"github.com/EfficientIP-Labs/terraform-provider-solidserver/solidserver"

	"github.com/AthenaWolfe/pulumi-solidserver/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "solidserver"
	// modules:
	mainMod = "index" // the solidserver module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(resource.PropertyMap, shim.ResourceConfig) error {
	return nil
}

//go:embed cmd/pulumi-resource-solidserver/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		// Instantiate the Terraform provider
		P:    shimv2.NewProvider(solidserver.Provider()),
		Name: "solidserver",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "AthenaWolfe",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing solidserver cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "solidserver", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/AthenaWolfe/pulumi-solidserver",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg:    "EfficientIP-Labs",
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
		Config:       map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
			//	"host": {
			//		Type: tfbridge.MakeResource(mainPkg, mainMod, "Host"),
			//		Default: &tfbridge.DefaultInfo{
			//			EnvVars: []string{"SDS_HOST"},
			//		},
			//	},
			//	"username": {
			//		Type: tfbridge.MakeResource(mainPkg, mainMod, "Username"),
			//		Default: &tfbridge.DefaultInfo{
			//			EnvVars: []string{"SDS_USERNAME"},
			//		},
			//	},
			//	"password": {
			//		Type: tfbridge.MakeResource(mainPkg, mainMod, "Password"),
			//		Default: &tfbridge.DefaultInfo{
			//			EnvVars: []string{"SDS_PASSWORD"},
			//		},
			//	},
			//	"solidserverversion": {
			//		Type: tfbridge.MakeResource(mainPkg, mainMod, "Solidserverversion"),
			//		Default: &tfbridge.DefaultInfo{
			//			EnvVars: []string{"SDS_SOLIDSERVERVERION"},
			//		},
			//	},
			//	"sslverify": {
			//		Type: tfbridge.MakeResource(mainPkg, mainMod, "Sslverify"),
			//		Default: &tfbridge.DefaultInfo{
			//			EnvVars: []string{"SDS_SSLVERIFY"},
			//		},
			//	},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags")},
			// 	},
			// },
			"solidserver_app_application":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AppApplication")},
			"solidserver_app_node":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AppNode")},
			"solidserver_app_pool":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AppPool")},
			"solidserver_cdb":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Cdb")},
			"solidserver_cdb_data":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CdbData")},
			"solidserver_device":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Device")},
			"solidserver_dns_forward_zone": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DnsForwardZone")},
			"solidserver_dns_rr":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DnsRr")},
			"solidserver_dns_server":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DnsServer")},
			"solidserver_dns_smart":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DnsSmart")},
			"solidserver_dns_view":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DnsView")},
			"solidserver_dns_zone":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DnsZone")},
			"solidserver_ip6_address":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Ip6Address")},
			"solidserver_ip6_alias":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Ip6Alias")},
			"solidserver_ip6_mac":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Ip6Mac")},
			"solidserver_ip6_pool":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Ip6Pool")},
			"solidserver_ip6_subnet":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Ip6Subnet")},
			"solidserver_ip_address":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IpAddress")},
			"solidserver_ip_alias":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IpAlias")},
			"solidserver_ip_mac":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IpMac")},
			"solidserver_ip_pool":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IpPool")},
			"solidserver_ip_space":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IpSpace")},
			"solidserver_ip_subnet":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IpSubnet")},
			"solidserver_user":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "User")},
			"solidserver_usergroup":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Usergroup")},
			"solidserver_vlan":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Vlan")},
			"solidserver_vlan_domain":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VlanDomain")},
			"solidserver_vlan_range":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VlanRange")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi")},
			"solidserver_cdb":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCdb")},
			"solidserver_cdb_data":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCdbData")},
			"solidserver_dns_server":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDnsServer")},
			"solidserver_dns_smart":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDnsSmart")},
			"solidserver_dns_view":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDnsView")},
			"solidserver_dns_zone":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDnsZone")},
			"solidserver_ip6_address":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIp6Address")},
			"solidserver_ip6_pool":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIp6Pool")},
			"solidserver_ip6_subnet":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIp6Subnet")},
			"solidserver_ip6_ptr":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIp6Ptr")},
			"solidserver_ip6_subnet_query": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIp6SubnetQuery")},
			"solidserver_ip_address":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIpAddress")},
			"solidserver_ip_pool":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIpPool")},
			"solidserver_ip_space":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIpSpace")},
			"solidserver_ip_ptr":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIpPtr")},
			"solidserver_ip_subnet":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIpSubnet")},
			"solidserver_ip_subnet_query":  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIpSubnetQuery")},
			"solidserver_usergroup":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUsergroup")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/AthenaWolfe/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// These are new API's that you may opt to use to automatically compute resource
	// tokens, and apply auto aliasing for full backwards compatibility.  For more
	// information, please reference:
	// https://pkg.go.dev/github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge#ProviderInfo.ComputeTokens
	prov.MustComputeTokens(tokens.SingleModule("solidserver_", mainMod,
		tokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
