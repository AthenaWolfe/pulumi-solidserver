// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package solidserver

import (
	"context"
	"reflect"

	"github.com/AthenaWolfe/pulumi-solidserver/sdk/go/solidserver/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DNS Zone data-source allows to retrieve information about DNS zones.
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
//			_, err := solidserver.LookupDnsZone(ctx, &solidserver.LookupDnsZoneArgs{
//				Name: solidserver_dns_zone.MyFirstZone.Name,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDnsZone(ctx *pulumi.Context, args *LookupDnsZoneArgs, opts ...pulumi.InvokeOption) (*LookupDnsZoneResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDnsZoneResult
	err := ctx.Invoke("solidserver:index/getDnsZone:getDnsZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDnsZone.
type LookupDnsZoneArgs struct {
	// The name of DNS view hosting the DNS zone.
	Dnsview *string `pulumi:"dnsview"`
	// The Domain Name served by the DNS zone.
	Name string `pulumi:"name"`
}

// A collection of values returned by getDnsZone.
type LookupDnsZoneResult struct {
	// The class associated to the DNS zone.
	Class string `pulumi:"class"`
	// The class parameters associated to IP space.
	ClassParameters map[string]interface{} `pulumi:"classParameters"`
	// Automaticaly create PTR records for the DNS zone.
	Createptr bool `pulumi:"createptr"`
	// The name of DNS server or DNS SMART hosting the DNS zone.
	Dnsserver string `pulumi:"dnsserver"`
	// The name of DNS view hosting the DNS zone.
	Dnsview *string `pulumi:"dnsview"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Domain Name served by the DNS zone.
	Name string `pulumi:"name"`
	// The name of a space associated to the DNS zone.
	Space string `pulumi:"space"`
	// The Type of the DNS zone.
	Type string `pulumi:"type"`
}

func LookupDnsZoneOutput(ctx *pulumi.Context, args LookupDnsZoneOutputArgs, opts ...pulumi.InvokeOption) LookupDnsZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDnsZoneResult, error) {
			args := v.(LookupDnsZoneArgs)
			r, err := LookupDnsZone(ctx, &args, opts...)
			var s LookupDnsZoneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDnsZoneResultOutput)
}

// A collection of arguments for invoking getDnsZone.
type LookupDnsZoneOutputArgs struct {
	// The name of DNS view hosting the DNS zone.
	Dnsview pulumi.StringPtrInput `pulumi:"dnsview"`
	// The Domain Name served by the DNS zone.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupDnsZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsZoneArgs)(nil)).Elem()
}

// A collection of values returned by getDnsZone.
type LookupDnsZoneResultOutput struct{ *pulumi.OutputState }

func (LookupDnsZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsZoneResult)(nil)).Elem()
}

func (o LookupDnsZoneResultOutput) ToLookupDnsZoneResultOutput() LookupDnsZoneResultOutput {
	return o
}

func (o LookupDnsZoneResultOutput) ToLookupDnsZoneResultOutputWithContext(ctx context.Context) LookupDnsZoneResultOutput {
	return o
}

// The class associated to the DNS zone.
func (o LookupDnsZoneResultOutput) Class() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) string { return v.Class }).(pulumi.StringOutput)
}

// The class parameters associated to IP space.
func (o LookupDnsZoneResultOutput) ClassParameters() pulumi.MapOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) map[string]interface{} { return v.ClassParameters }).(pulumi.MapOutput)
}

// Automaticaly create PTR records for the DNS zone.
func (o LookupDnsZoneResultOutput) Createptr() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) bool { return v.Createptr }).(pulumi.BoolOutput)
}

// The name of DNS server or DNS SMART hosting the DNS zone.
func (o LookupDnsZoneResultOutput) Dnsserver() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) string { return v.Dnsserver }).(pulumi.StringOutput)
}

// The name of DNS view hosting the DNS zone.
func (o LookupDnsZoneResultOutput) Dnsview() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) *string { return v.Dnsview }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDnsZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Domain Name served by the DNS zone.
func (o LookupDnsZoneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) string { return v.Name }).(pulumi.StringOutput)
}

// The name of a space associated to the DNS zone.
func (o LookupDnsZoneResultOutput) Space() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) string { return v.Space }).(pulumi.StringOutput)
}

// The Type of the DNS zone.
func (o LookupDnsZoneResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDnsZoneResultOutput{})
}