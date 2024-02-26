// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package solidserver

import (
	"context"
	"reflect"

	"github.com/AthenaWolfe/pulumi-solidserver/sdk/go/solidserver/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// IPv6 pool data-source allows to retrieve information about reserved IPv6 pools including meta-data.
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
//			_, err := solidserver.LookupIp6Pool(ctx, &solidserver.LookupIp6PoolArgs{
//				Name:   solidserver_ip6_subnet.MyFirstIPv6Pool.Name,
//				Subnet: solidserver_ip6_subnet.MyFirstIPv6Pool.Subnet,
//				Space:  solidserver_ip6_subnet.MyFirstIPv6Pool.Space,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIp6Pool(ctx *pulumi.Context, args *LookupIp6PoolArgs, opts ...pulumi.InvokeOption) (*LookupIp6PoolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIp6PoolResult
	err := ctx.Invoke("solidserver:index/getIp6Pool:getIp6Pool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIp6Pool.
type LookupIp6PoolArgs struct {
	// The name of the IPv6 pool.
	Name string `pulumi:"name"`
	// The space associated to the IPv6 pool.
	Space string `pulumi:"space"`
	// The parent subnet of the IPv6 pool.
	Subnet string `pulumi:"subnet"`
}

// A collection of values returned by getIp6Pool.
type LookupIp6PoolResult struct {
	// The class associated to the IPv6 pool.
	Class string `pulumi:"class"`
	// The class parameters associated to the IPv6 pool.
	ClassParameters map[string]interface{} `pulumi:"classParameters"`
	// The last address of the IPv6 pool.
	End string `pulumi:"end"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the IPv6 pool.
	Name string `pulumi:"name"`
	// The prefix of the parent subnet of the IPv6 pool.
	Prefix string `pulumi:"prefix"`
	// The size prefix of the parent subnet of the IPv6 pool.
	PrefixSize int `pulumi:"prefixSize"`
	// The space associated to the IPv6 pool.
	Space string `pulumi:"space"`
	// The fisrt address of the IPv6 pool.
	Start string `pulumi:"start"`
	// The parent subnet of the IPv6 pool.
	Subnet string `pulumi:"subnet"`
}

func LookupIp6PoolOutput(ctx *pulumi.Context, args LookupIp6PoolOutputArgs, opts ...pulumi.InvokeOption) LookupIp6PoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIp6PoolResult, error) {
			args := v.(LookupIp6PoolArgs)
			r, err := LookupIp6Pool(ctx, &args, opts...)
			var s LookupIp6PoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIp6PoolResultOutput)
}

// A collection of arguments for invoking getIp6Pool.
type LookupIp6PoolOutputArgs struct {
	// The name of the IPv6 pool.
	Name pulumi.StringInput `pulumi:"name"`
	// The space associated to the IPv6 pool.
	Space pulumi.StringInput `pulumi:"space"`
	// The parent subnet of the IPv6 pool.
	Subnet pulumi.StringInput `pulumi:"subnet"`
}

func (LookupIp6PoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIp6PoolArgs)(nil)).Elem()
}

// A collection of values returned by getIp6Pool.
type LookupIp6PoolResultOutput struct{ *pulumi.OutputState }

func (LookupIp6PoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIp6PoolResult)(nil)).Elem()
}

func (o LookupIp6PoolResultOutput) ToLookupIp6PoolResultOutput() LookupIp6PoolResultOutput {
	return o
}

func (o LookupIp6PoolResultOutput) ToLookupIp6PoolResultOutputWithContext(ctx context.Context) LookupIp6PoolResultOutput {
	return o
}

// The class associated to the IPv6 pool.
func (o LookupIp6PoolResultOutput) Class() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIp6PoolResult) string { return v.Class }).(pulumi.StringOutput)
}

// The class parameters associated to the IPv6 pool.
func (o LookupIp6PoolResultOutput) ClassParameters() pulumi.MapOutput {
	return o.ApplyT(func(v LookupIp6PoolResult) map[string]interface{} { return v.ClassParameters }).(pulumi.MapOutput)
}

// The last address of the IPv6 pool.
func (o LookupIp6PoolResultOutput) End() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIp6PoolResult) string { return v.End }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIp6PoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIp6PoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the IPv6 pool.
func (o LookupIp6PoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIp6PoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// The prefix of the parent subnet of the IPv6 pool.
func (o LookupIp6PoolResultOutput) Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIp6PoolResult) string { return v.Prefix }).(pulumi.StringOutput)
}

// The size prefix of the parent subnet of the IPv6 pool.
func (o LookupIp6PoolResultOutput) PrefixSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIp6PoolResult) int { return v.PrefixSize }).(pulumi.IntOutput)
}

// The space associated to the IPv6 pool.
func (o LookupIp6PoolResultOutput) Space() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIp6PoolResult) string { return v.Space }).(pulumi.StringOutput)
}

// The fisrt address of the IPv6 pool.
func (o LookupIp6PoolResultOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIp6PoolResult) string { return v.Start }).(pulumi.StringOutput)
}

// The parent subnet of the IPv6 pool.
func (o LookupIp6PoolResultOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIp6PoolResult) string { return v.Subnet }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIp6PoolResultOutput{})
}
