// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package solidserver

import (
	"context"
	"reflect"

	"github.com/AthenaWolfe/pulumi-solidserver/sdk/go/solidserver/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// IP pool data-source allows to retrieve information about reserved IPv4 pools including meta-data.
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
//			_, err := solidserver.LookupIpPool(ctx, &solidserver.LookupIpPoolArgs{
//				Name:   solidserver_ip_subnet.MyFirstIPPool.Name,
//				Subnet: solidserver_ip_subnet.MyFirstIPPool.Subnet,
//				Space:  solidserver_ip_subnet.MyFirstIPPool.Space,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIpPool(ctx *pulumi.Context, args *LookupIpPoolArgs, opts ...pulumi.InvokeOption) (*LookupIpPoolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpPoolResult
	err := ctx.Invoke("solidserver:index/getIpPool:getIpPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpPool.
type LookupIpPoolArgs struct {
	// The name of the IP pool.
	Name string `pulumi:"name"`
	// The space associated to the IP pool.
	Space string `pulumi:"space"`
	// The parent subnet of the IP pool.
	Subnet string `pulumi:"subnet"`
}

// A collection of values returned by getIpPool.
type LookupIpPoolResult struct {
	// The class associated to the IP pool.
	Class string `pulumi:"class"`
	// The class parameters associated to the IP pool.
	ClassParameters map[string]interface{} `pulumi:"classParameters"`
	// The last address of the IP pool.
	End string `pulumi:"end"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the IP pool.
	Name string `pulumi:"name"`
	// The prefix of the parent subnet of the IP pool.
	Prefix string `pulumi:"prefix"`
	// The size prefix of the parent subnet of the IP pool.
	PrefixSize int `pulumi:"prefixSize"`
	// The size of the IP pool.
	Size string `pulumi:"size"`
	// The space associated to the IP pool.
	Space string `pulumi:"space"`
	// The fisrt address of the IP pool.
	Start string `pulumi:"start"`
	// The parent subnet of the IP pool.
	Subnet string `pulumi:"subnet"`
}

func LookupIpPoolOutput(ctx *pulumi.Context, args LookupIpPoolOutputArgs, opts ...pulumi.InvokeOption) LookupIpPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpPoolResult, error) {
			args := v.(LookupIpPoolArgs)
			r, err := LookupIpPool(ctx, &args, opts...)
			var s LookupIpPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpPoolResultOutput)
}

// A collection of arguments for invoking getIpPool.
type LookupIpPoolOutputArgs struct {
	// The name of the IP pool.
	Name pulumi.StringInput `pulumi:"name"`
	// The space associated to the IP pool.
	Space pulumi.StringInput `pulumi:"space"`
	// The parent subnet of the IP pool.
	Subnet pulumi.StringInput `pulumi:"subnet"`
}

func (LookupIpPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpPoolArgs)(nil)).Elem()
}

// A collection of values returned by getIpPool.
type LookupIpPoolResultOutput struct{ *pulumi.OutputState }

func (LookupIpPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpPoolResult)(nil)).Elem()
}

func (o LookupIpPoolResultOutput) ToLookupIpPoolResultOutput() LookupIpPoolResultOutput {
	return o
}

func (o LookupIpPoolResultOutput) ToLookupIpPoolResultOutputWithContext(ctx context.Context) LookupIpPoolResultOutput {
	return o
}

// The class associated to the IP pool.
func (o LookupIpPoolResultOutput) Class() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpPoolResult) string { return v.Class }).(pulumi.StringOutput)
}

// The class parameters associated to the IP pool.
func (o LookupIpPoolResultOutput) ClassParameters() pulumi.MapOutput {
	return o.ApplyT(func(v LookupIpPoolResult) map[string]interface{} { return v.ClassParameters }).(pulumi.MapOutput)
}

// The last address of the IP pool.
func (o LookupIpPoolResultOutput) End() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpPoolResult) string { return v.End }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIpPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the IP pool.
func (o LookupIpPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// The prefix of the parent subnet of the IP pool.
func (o LookupIpPoolResultOutput) Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpPoolResult) string { return v.Prefix }).(pulumi.StringOutput)
}

// The size prefix of the parent subnet of the IP pool.
func (o LookupIpPoolResultOutput) PrefixSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIpPoolResult) int { return v.PrefixSize }).(pulumi.IntOutput)
}

// The size of the IP pool.
func (o LookupIpPoolResultOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpPoolResult) string { return v.Size }).(pulumi.StringOutput)
}

// The space associated to the IP pool.
func (o LookupIpPoolResultOutput) Space() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpPoolResult) string { return v.Space }).(pulumi.StringOutput)
}

// The fisrt address of the IP pool.
func (o LookupIpPoolResultOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpPoolResult) string { return v.Start }).(pulumi.StringOutput)
}

// The parent subnet of the IP pool.
func (o LookupIpPoolResultOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpPoolResult) string { return v.Subnet }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpPoolResultOutput{})
}
