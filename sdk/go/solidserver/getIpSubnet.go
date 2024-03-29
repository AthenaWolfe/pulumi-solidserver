// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package solidserver

import (
	"context"
	"reflect"

	"github.com/AthenaWolfe/pulumi-solidserver/sdk/go/solidserver/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// IP subnet data-source allows to retrieve information about reserved IPv4 subnets, including meta-data.
// IP Subnet are key to organize the IP space, they can be blocks or subnets. Blocks reflect assigned IP
// ranges (RFC1918 or public prefixes). Subnets reflect the internal sub-division of your network.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/AthenaWolfe/pulumi-solidserver/sdk/go/solidserver"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := solidserver.LookupIpSubnet(ctx, &solidserver.LookupIpSubnetArgs{
//				Name:  solidserver_ip_subnet.MyFirstIPSubnet.Name,
//				Space: solidserver_ip_subnet.MyFirstIPSubnet.Space,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIpSubnet(ctx *pulumi.Context, args *LookupIpSubnetArgs, opts ...pulumi.InvokeOption) (*LookupIpSubnetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpSubnetResult
	err := ctx.Invoke("solidserver:index/getIpSubnet:getIpSubnet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpSubnet.
type LookupIpSubnetArgs struct {
	// The name of the IP subnet.
	Name string `pulumi:"name"`
	// The space associated to the IP subnet.
	Space string `pulumi:"space"`
}

// A collection of values returned by getIpSubnet.
type LookupIpSubnetResult struct {
	// The IP subnet address.
	Address string `pulumi:"address"`
	// The class associated to the IP subnet.
	Class string `pulumi:"class"`
	// The class parameters associated to IP subnet.
	ClassParameters map[string]string `pulumi:"classParameters"`
	// The subnet's computed gateway.
	Gateway string `pulumi:"gateway"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the IP subnet.
	Name string `pulumi:"name"`
	// The IP subnet netmask.
	Netmask string `pulumi:"netmask"`
	// The IP subnet prefix.
	Prefix string `pulumi:"prefix"`
	// The IP subnet's prefix length (ex: 24 for a '/24').
	PrefixSize int `pulumi:"prefixSize"`
	// The space associated to the IP subnet.
	Space string `pulumi:"space"`
	// The terminal property of the IP subnet.
	Terminal bool `pulumi:"terminal"`
	// The optional vlan ID associated with the subnet.
	VlanId int `pulumi:"vlanId"`
}

func LookupIpSubnetOutput(ctx *pulumi.Context, args LookupIpSubnetOutputArgs, opts ...pulumi.InvokeOption) LookupIpSubnetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpSubnetResult, error) {
			args := v.(LookupIpSubnetArgs)
			r, err := LookupIpSubnet(ctx, &args, opts...)
			var s LookupIpSubnetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpSubnetResultOutput)
}

// A collection of arguments for invoking getIpSubnet.
type LookupIpSubnetOutputArgs struct {
	// The name of the IP subnet.
	Name pulumi.StringInput `pulumi:"name"`
	// The space associated to the IP subnet.
	Space pulumi.StringInput `pulumi:"space"`
}

func (LookupIpSubnetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpSubnetArgs)(nil)).Elem()
}

// A collection of values returned by getIpSubnet.
type LookupIpSubnetResultOutput struct{ *pulumi.OutputState }

func (LookupIpSubnetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpSubnetResult)(nil)).Elem()
}

func (o LookupIpSubnetResultOutput) ToLookupIpSubnetResultOutput() LookupIpSubnetResultOutput {
	return o
}

func (o LookupIpSubnetResultOutput) ToLookupIpSubnetResultOutputWithContext(ctx context.Context) LookupIpSubnetResultOutput {
	return o
}

// The IP subnet address.
func (o LookupIpSubnetResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpSubnetResult) string { return v.Address }).(pulumi.StringOutput)
}

// The class associated to the IP subnet.
func (o LookupIpSubnetResultOutput) Class() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpSubnetResult) string { return v.Class }).(pulumi.StringOutput)
}

// The class parameters associated to IP subnet.
func (o LookupIpSubnetResultOutput) ClassParameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIpSubnetResult) map[string]string { return v.ClassParameters }).(pulumi.StringMapOutput)
}

// The subnet's computed gateway.
func (o LookupIpSubnetResultOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpSubnetResult) string { return v.Gateway }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIpSubnetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpSubnetResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the IP subnet.
func (o LookupIpSubnetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpSubnetResult) string { return v.Name }).(pulumi.StringOutput)
}

// The IP subnet netmask.
func (o LookupIpSubnetResultOutput) Netmask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpSubnetResult) string { return v.Netmask }).(pulumi.StringOutput)
}

// The IP subnet prefix.
func (o LookupIpSubnetResultOutput) Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpSubnetResult) string { return v.Prefix }).(pulumi.StringOutput)
}

// The IP subnet's prefix length (ex: 24 for a '/24').
func (o LookupIpSubnetResultOutput) PrefixSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIpSubnetResult) int { return v.PrefixSize }).(pulumi.IntOutput)
}

// The space associated to the IP subnet.
func (o LookupIpSubnetResultOutput) Space() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpSubnetResult) string { return v.Space }).(pulumi.StringOutput)
}

// The terminal property of the IP subnet.
func (o LookupIpSubnetResultOutput) Terminal() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpSubnetResult) bool { return v.Terminal }).(pulumi.BoolOutput)
}

// The optional vlan ID associated with the subnet.
func (o LookupIpSubnetResultOutput) VlanId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIpSubnetResult) int { return v.VlanId }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpSubnetResultOutput{})
}
