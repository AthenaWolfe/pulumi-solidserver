// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package solidserver

import (
	"context"
	"reflect"

	"errors"
	"github.com/AthenaWolfe/pulumi-solidserver/sdk/go/solidserver/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DNS Zone resource allows to create and configure DNS zones.
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
//			_, err := solidserver.NewDnsZone(ctx, "myFirstZone", &solidserver.DnsZoneArgs{
//				Createptr: pulumi.Bool(false),
//				Dnsserver: pulumi.String("ns.priv"),
//				Space:     pulumi.Any(solidserver_ip_space.MyFirstSpace.Name),
//				Type:      pulumi.String("master"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DnsZone struct {
	pulumi.CustomResourceState

	// The list of IP addresses (Format \n\n:\n\n) that will receive zone change notifications in addition to the NS listed in the SOA
	AlsoNotifies pulumi.StringArrayOutput `pulumi:"alsoNotifies"`
	// The class associated to the zone.
	Class pulumi.StringPtrOutput `pulumi:"class"`
	// The class parameters associated to the zone.
	ClassParameters pulumi.StringMapOutput `pulumi:"classParameters"`
	// Automaticaly create PTR records for the zone.
	Createptr pulumi.BoolPtrOutput `pulumi:"createptr"`
	// The name of DNS server or DNS SMART hosting the DNS zone to create.
	Dnsserver pulumi.StringOutput `pulumi:"dnsserver"`
	// The name of DNS view hosting the DNS zone to create.
	Dnsview pulumi.StringPtrOutput `pulumi:"dnsview"`
	// The Domain Name to be hosted by the zone.
	Name pulumi.StringOutput `pulumi:"name"`
	// The expected notify behavior (Supported: empty (Inherited), Yes, No, Explicit; Default: empty (Inherited).
	Notify pulumi.StringPtrOutput `pulumi:"notify"`
	// The name of a space associated to the zone.
	Space pulumi.StringPtrOutput `pulumi:"space"`
	// The type of the zone to create (Supported: Master).
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewDnsZone registers a new resource with the given unique name, arguments, and options.
func NewDnsZone(ctx *pulumi.Context,
	name string, args *DnsZoneArgs, opts ...pulumi.ResourceOption) (*DnsZone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dnsserver == nil {
		return nil, errors.New("invalid value for required argument 'Dnsserver'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DnsZone
	err := ctx.RegisterResource("solidserver:index/dnsZone:DnsZone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsZone gets an existing DnsZone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsZoneState, opts ...pulumi.ResourceOption) (*DnsZone, error) {
	var resource DnsZone
	err := ctx.ReadResource("solidserver:index/dnsZone:DnsZone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsZone resources.
type dnsZoneState struct {
	// The list of IP addresses (Format \n\n:\n\n) that will receive zone change notifications in addition to the NS listed in the SOA
	AlsoNotifies []string `pulumi:"alsoNotifies"`
	// The class associated to the zone.
	Class *string `pulumi:"class"`
	// The class parameters associated to the zone.
	ClassParameters map[string]string `pulumi:"classParameters"`
	// Automaticaly create PTR records for the zone.
	Createptr *bool `pulumi:"createptr"`
	// The name of DNS server or DNS SMART hosting the DNS zone to create.
	Dnsserver *string `pulumi:"dnsserver"`
	// The name of DNS view hosting the DNS zone to create.
	Dnsview *string `pulumi:"dnsview"`
	// The Domain Name to be hosted by the zone.
	Name *string `pulumi:"name"`
	// The expected notify behavior (Supported: empty (Inherited), Yes, No, Explicit; Default: empty (Inherited).
	Notify *string `pulumi:"notify"`
	// The name of a space associated to the zone.
	Space *string `pulumi:"space"`
	// The type of the zone to create (Supported: Master).
	Type *string `pulumi:"type"`
}

type DnsZoneState struct {
	// The list of IP addresses (Format \n\n:\n\n) that will receive zone change notifications in addition to the NS listed in the SOA
	AlsoNotifies pulumi.StringArrayInput
	// The class associated to the zone.
	Class pulumi.StringPtrInput
	// The class parameters associated to the zone.
	ClassParameters pulumi.StringMapInput
	// Automaticaly create PTR records for the zone.
	Createptr pulumi.BoolPtrInput
	// The name of DNS server or DNS SMART hosting the DNS zone to create.
	Dnsserver pulumi.StringPtrInput
	// The name of DNS view hosting the DNS zone to create.
	Dnsview pulumi.StringPtrInput
	// The Domain Name to be hosted by the zone.
	Name pulumi.StringPtrInput
	// The expected notify behavior (Supported: empty (Inherited), Yes, No, Explicit; Default: empty (Inherited).
	Notify pulumi.StringPtrInput
	// The name of a space associated to the zone.
	Space pulumi.StringPtrInput
	// The type of the zone to create (Supported: Master).
	Type pulumi.StringPtrInput
}

func (DnsZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsZoneState)(nil)).Elem()
}

type dnsZoneArgs struct {
	// The list of IP addresses (Format \n\n:\n\n) that will receive zone change notifications in addition to the NS listed in the SOA
	AlsoNotifies []string `pulumi:"alsoNotifies"`
	// The class associated to the zone.
	Class *string `pulumi:"class"`
	// The class parameters associated to the zone.
	ClassParameters map[string]string `pulumi:"classParameters"`
	// Automaticaly create PTR records for the zone.
	Createptr *bool `pulumi:"createptr"`
	// The name of DNS server or DNS SMART hosting the DNS zone to create.
	Dnsserver string `pulumi:"dnsserver"`
	// The name of DNS view hosting the DNS zone to create.
	Dnsview *string `pulumi:"dnsview"`
	// The Domain Name to be hosted by the zone.
	Name *string `pulumi:"name"`
	// The expected notify behavior (Supported: empty (Inherited), Yes, No, Explicit; Default: empty (Inherited).
	Notify *string `pulumi:"notify"`
	// The name of a space associated to the zone.
	Space *string `pulumi:"space"`
	// The type of the zone to create (Supported: Master).
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a DnsZone resource.
type DnsZoneArgs struct {
	// The list of IP addresses (Format \n\n:\n\n) that will receive zone change notifications in addition to the NS listed in the SOA
	AlsoNotifies pulumi.StringArrayInput
	// The class associated to the zone.
	Class pulumi.StringPtrInput
	// The class parameters associated to the zone.
	ClassParameters pulumi.StringMapInput
	// Automaticaly create PTR records for the zone.
	Createptr pulumi.BoolPtrInput
	// The name of DNS server or DNS SMART hosting the DNS zone to create.
	Dnsserver pulumi.StringInput
	// The name of DNS view hosting the DNS zone to create.
	Dnsview pulumi.StringPtrInput
	// The Domain Name to be hosted by the zone.
	Name pulumi.StringPtrInput
	// The expected notify behavior (Supported: empty (Inherited), Yes, No, Explicit; Default: empty (Inherited).
	Notify pulumi.StringPtrInput
	// The name of a space associated to the zone.
	Space pulumi.StringPtrInput
	// The type of the zone to create (Supported: Master).
	Type pulumi.StringPtrInput
}

func (DnsZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsZoneArgs)(nil)).Elem()
}

type DnsZoneInput interface {
	pulumi.Input

	ToDnsZoneOutput() DnsZoneOutput
	ToDnsZoneOutputWithContext(ctx context.Context) DnsZoneOutput
}

func (*DnsZone) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsZone)(nil)).Elem()
}

func (i *DnsZone) ToDnsZoneOutput() DnsZoneOutput {
	return i.ToDnsZoneOutputWithContext(context.Background())
}

func (i *DnsZone) ToDnsZoneOutputWithContext(ctx context.Context) DnsZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsZoneOutput)
}

// DnsZoneArrayInput is an input type that accepts DnsZoneArray and DnsZoneArrayOutput values.
// You can construct a concrete instance of `DnsZoneArrayInput` via:
//
//	DnsZoneArray{ DnsZoneArgs{...} }
type DnsZoneArrayInput interface {
	pulumi.Input

	ToDnsZoneArrayOutput() DnsZoneArrayOutput
	ToDnsZoneArrayOutputWithContext(context.Context) DnsZoneArrayOutput
}

type DnsZoneArray []DnsZoneInput

func (DnsZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsZone)(nil)).Elem()
}

func (i DnsZoneArray) ToDnsZoneArrayOutput() DnsZoneArrayOutput {
	return i.ToDnsZoneArrayOutputWithContext(context.Background())
}

func (i DnsZoneArray) ToDnsZoneArrayOutputWithContext(ctx context.Context) DnsZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsZoneArrayOutput)
}

// DnsZoneMapInput is an input type that accepts DnsZoneMap and DnsZoneMapOutput values.
// You can construct a concrete instance of `DnsZoneMapInput` via:
//
//	DnsZoneMap{ "key": DnsZoneArgs{...} }
type DnsZoneMapInput interface {
	pulumi.Input

	ToDnsZoneMapOutput() DnsZoneMapOutput
	ToDnsZoneMapOutputWithContext(context.Context) DnsZoneMapOutput
}

type DnsZoneMap map[string]DnsZoneInput

func (DnsZoneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsZone)(nil)).Elem()
}

func (i DnsZoneMap) ToDnsZoneMapOutput() DnsZoneMapOutput {
	return i.ToDnsZoneMapOutputWithContext(context.Background())
}

func (i DnsZoneMap) ToDnsZoneMapOutputWithContext(ctx context.Context) DnsZoneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsZoneMapOutput)
}

type DnsZoneOutput struct{ *pulumi.OutputState }

func (DnsZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsZone)(nil)).Elem()
}

func (o DnsZoneOutput) ToDnsZoneOutput() DnsZoneOutput {
	return o
}

func (o DnsZoneOutput) ToDnsZoneOutputWithContext(ctx context.Context) DnsZoneOutput {
	return o
}

// The list of IP addresses (Format \n\n:\n\n) that will receive zone change notifications in addition to the NS listed in the SOA
func (o DnsZoneOutput) AlsoNotifies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringArrayOutput { return v.AlsoNotifies }).(pulumi.StringArrayOutput)
}

// The class associated to the zone.
func (o DnsZoneOutput) Class() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringPtrOutput { return v.Class }).(pulumi.StringPtrOutput)
}

// The class parameters associated to the zone.
func (o DnsZoneOutput) ClassParameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringMapOutput { return v.ClassParameters }).(pulumi.StringMapOutput)
}

// Automaticaly create PTR records for the zone.
func (o DnsZoneOutput) Createptr() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.BoolPtrOutput { return v.Createptr }).(pulumi.BoolPtrOutput)
}

// The name of DNS server or DNS SMART hosting the DNS zone to create.
func (o DnsZoneOutput) Dnsserver() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringOutput { return v.Dnsserver }).(pulumi.StringOutput)
}

// The name of DNS view hosting the DNS zone to create.
func (o DnsZoneOutput) Dnsview() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringPtrOutput { return v.Dnsview }).(pulumi.StringPtrOutput)
}

// The Domain Name to be hosted by the zone.
func (o DnsZoneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The expected notify behavior (Supported: empty (Inherited), Yes, No, Explicit; Default: empty (Inherited).
func (o DnsZoneOutput) Notify() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringPtrOutput { return v.Notify }).(pulumi.StringPtrOutput)
}

// The name of a space associated to the zone.
func (o DnsZoneOutput) Space() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringPtrOutput { return v.Space }).(pulumi.StringPtrOutput)
}

// The type of the zone to create (Supported: Master).
func (o DnsZoneOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type DnsZoneArrayOutput struct{ *pulumi.OutputState }

func (DnsZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsZone)(nil)).Elem()
}

func (o DnsZoneArrayOutput) ToDnsZoneArrayOutput() DnsZoneArrayOutput {
	return o
}

func (o DnsZoneArrayOutput) ToDnsZoneArrayOutputWithContext(ctx context.Context) DnsZoneArrayOutput {
	return o
}

func (o DnsZoneArrayOutput) Index(i pulumi.IntInput) DnsZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsZone {
		return vs[0].([]*DnsZone)[vs[1].(int)]
	}).(DnsZoneOutput)
}

type DnsZoneMapOutput struct{ *pulumi.OutputState }

func (DnsZoneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsZone)(nil)).Elem()
}

func (o DnsZoneMapOutput) ToDnsZoneMapOutput() DnsZoneMapOutput {
	return o
}

func (o DnsZoneMapOutput) ToDnsZoneMapOutputWithContext(ctx context.Context) DnsZoneMapOutput {
	return o
}

func (o DnsZoneMapOutput) MapIndex(k pulumi.StringInput) DnsZoneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsZone {
		return vs[0].(map[string]*DnsZone)[vs[1].(string)]
	}).(DnsZoneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsZoneInput)(nil)).Elem(), &DnsZone{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsZoneArrayInput)(nil)).Elem(), DnsZoneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsZoneMapInput)(nil)).Elem(), DnsZoneMap{})
	pulumi.RegisterOutputType(DnsZoneOutput{})
	pulumi.RegisterOutputType(DnsZoneArrayOutput{})
	pulumi.RegisterOutputType(DnsZoneMapOutput{})
}