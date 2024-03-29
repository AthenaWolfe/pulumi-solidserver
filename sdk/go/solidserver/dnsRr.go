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

// DNS RR resource allows to create and manage DNS resource records of type A, AAAA, PTR, CNAME, DNAME, NS.
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
//			_, err := solidserver.GetIpPtr(ctx, &solidserver.GetIpPtrArgs{
//				Address: solidserver_ip_address.MyFirstIPAddress.Address,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = solidserver.NewDnsRr(ctx, "aaRecord", &solidserver.DnsRrArgs{
//				Dnsserver: pulumi.String("ns.mycompany.priv"),
//				Dnsview:   pulumi.String("Internal"),
//				Dnszone:   pulumi.String("mycompany.priv"),
//				Type:      pulumi.String("PTR"),
//				Value:     pulumi.String("myapp.mycompany.priv"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DnsRr struct {
	pulumi.CustomResourceState

	// The class associated to the DNS view.
	Class pulumi.StringPtrOutput `pulumi:"class"`
	// The class parameters associated to the view.
	ClassParameters pulumi.StringMapOutput `pulumi:"classParameters"`
	// The managed SMART DNS server name, or DNS server name hosting the RR's zone.
	Dnsserver pulumi.StringOutput `pulumi:"dnsserver"`
	// The View name of the RR to create.
	Dnsview pulumi.StringPtrOutput `pulumi:"dnsview"`
	// The Zone name of the RR to create.
	Dnszone pulumi.StringPtrOutput `pulumi:"dnszone"`
	// The Fully Qualified Domain Name of the RR to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// The DNS Time To Live of the RR to create.
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
	// The type of the RR to create (Supported: A, AAAA, PTR, CNAME, DNAME and NS).
	Type pulumi.StringOutput `pulumi:"type"`
	// The value od the RR to create.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewDnsRr registers a new resource with the given unique name, arguments, and options.
func NewDnsRr(ctx *pulumi.Context,
	name string, args *DnsRrArgs, opts ...pulumi.ResourceOption) (*DnsRr, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dnsserver == nil {
		return nil, errors.New("invalid value for required argument 'Dnsserver'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DnsRr
	err := ctx.RegisterResource("solidserver:index/dnsRr:DnsRr", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsRr gets an existing DnsRr resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsRr(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsRrState, opts ...pulumi.ResourceOption) (*DnsRr, error) {
	var resource DnsRr
	err := ctx.ReadResource("solidserver:index/dnsRr:DnsRr", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsRr resources.
type dnsRrState struct {
	// The class associated to the DNS view.
	Class *string `pulumi:"class"`
	// The class parameters associated to the view.
	ClassParameters map[string]string `pulumi:"classParameters"`
	// The managed SMART DNS server name, or DNS server name hosting the RR's zone.
	Dnsserver *string `pulumi:"dnsserver"`
	// The View name of the RR to create.
	Dnsview *string `pulumi:"dnsview"`
	// The Zone name of the RR to create.
	Dnszone *string `pulumi:"dnszone"`
	// The Fully Qualified Domain Name of the RR to create.
	Name *string `pulumi:"name"`
	// The DNS Time To Live of the RR to create.
	Ttl *int `pulumi:"ttl"`
	// The type of the RR to create (Supported: A, AAAA, PTR, CNAME, DNAME and NS).
	Type *string `pulumi:"type"`
	// The value od the RR to create.
	Value *string `pulumi:"value"`
}

type DnsRrState struct {
	// The class associated to the DNS view.
	Class pulumi.StringPtrInput
	// The class parameters associated to the view.
	ClassParameters pulumi.StringMapInput
	// The managed SMART DNS server name, or DNS server name hosting the RR's zone.
	Dnsserver pulumi.StringPtrInput
	// The View name of the RR to create.
	Dnsview pulumi.StringPtrInput
	// The Zone name of the RR to create.
	Dnszone pulumi.StringPtrInput
	// The Fully Qualified Domain Name of the RR to create.
	Name pulumi.StringPtrInput
	// The DNS Time To Live of the RR to create.
	Ttl pulumi.IntPtrInput
	// The type of the RR to create (Supported: A, AAAA, PTR, CNAME, DNAME and NS).
	Type pulumi.StringPtrInput
	// The value od the RR to create.
	Value pulumi.StringPtrInput
}

func (DnsRrState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsRrState)(nil)).Elem()
}

type dnsRrArgs struct {
	// The class associated to the DNS view.
	Class *string `pulumi:"class"`
	// The class parameters associated to the view.
	ClassParameters map[string]string `pulumi:"classParameters"`
	// The managed SMART DNS server name, or DNS server name hosting the RR's zone.
	Dnsserver string `pulumi:"dnsserver"`
	// The View name of the RR to create.
	Dnsview *string `pulumi:"dnsview"`
	// The Zone name of the RR to create.
	Dnszone *string `pulumi:"dnszone"`
	// The Fully Qualified Domain Name of the RR to create.
	Name *string `pulumi:"name"`
	// The DNS Time To Live of the RR to create.
	Ttl *int `pulumi:"ttl"`
	// The type of the RR to create (Supported: A, AAAA, PTR, CNAME, DNAME and NS).
	Type string `pulumi:"type"`
	// The value od the RR to create.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a DnsRr resource.
type DnsRrArgs struct {
	// The class associated to the DNS view.
	Class pulumi.StringPtrInput
	// The class parameters associated to the view.
	ClassParameters pulumi.StringMapInput
	// The managed SMART DNS server name, or DNS server name hosting the RR's zone.
	Dnsserver pulumi.StringInput
	// The View name of the RR to create.
	Dnsview pulumi.StringPtrInput
	// The Zone name of the RR to create.
	Dnszone pulumi.StringPtrInput
	// The Fully Qualified Domain Name of the RR to create.
	Name pulumi.StringPtrInput
	// The DNS Time To Live of the RR to create.
	Ttl pulumi.IntPtrInput
	// The type of the RR to create (Supported: A, AAAA, PTR, CNAME, DNAME and NS).
	Type pulumi.StringInput
	// The value od the RR to create.
	Value pulumi.StringInput
}

func (DnsRrArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsRrArgs)(nil)).Elem()
}

type DnsRrInput interface {
	pulumi.Input

	ToDnsRrOutput() DnsRrOutput
	ToDnsRrOutputWithContext(ctx context.Context) DnsRrOutput
}

func (*DnsRr) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsRr)(nil)).Elem()
}

func (i *DnsRr) ToDnsRrOutput() DnsRrOutput {
	return i.ToDnsRrOutputWithContext(context.Background())
}

func (i *DnsRr) ToDnsRrOutputWithContext(ctx context.Context) DnsRrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRrOutput)
}

// DnsRrArrayInput is an input type that accepts DnsRrArray and DnsRrArrayOutput values.
// You can construct a concrete instance of `DnsRrArrayInput` via:
//
//	DnsRrArray{ DnsRrArgs{...} }
type DnsRrArrayInput interface {
	pulumi.Input

	ToDnsRrArrayOutput() DnsRrArrayOutput
	ToDnsRrArrayOutputWithContext(context.Context) DnsRrArrayOutput
}

type DnsRrArray []DnsRrInput

func (DnsRrArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsRr)(nil)).Elem()
}

func (i DnsRrArray) ToDnsRrArrayOutput() DnsRrArrayOutput {
	return i.ToDnsRrArrayOutputWithContext(context.Background())
}

func (i DnsRrArray) ToDnsRrArrayOutputWithContext(ctx context.Context) DnsRrArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRrArrayOutput)
}

// DnsRrMapInput is an input type that accepts DnsRrMap and DnsRrMapOutput values.
// You can construct a concrete instance of `DnsRrMapInput` via:
//
//	DnsRrMap{ "key": DnsRrArgs{...} }
type DnsRrMapInput interface {
	pulumi.Input

	ToDnsRrMapOutput() DnsRrMapOutput
	ToDnsRrMapOutputWithContext(context.Context) DnsRrMapOutput
}

type DnsRrMap map[string]DnsRrInput

func (DnsRrMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsRr)(nil)).Elem()
}

func (i DnsRrMap) ToDnsRrMapOutput() DnsRrMapOutput {
	return i.ToDnsRrMapOutputWithContext(context.Background())
}

func (i DnsRrMap) ToDnsRrMapOutputWithContext(ctx context.Context) DnsRrMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRrMapOutput)
}

type DnsRrOutput struct{ *pulumi.OutputState }

func (DnsRrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsRr)(nil)).Elem()
}

func (o DnsRrOutput) ToDnsRrOutput() DnsRrOutput {
	return o
}

func (o DnsRrOutput) ToDnsRrOutputWithContext(ctx context.Context) DnsRrOutput {
	return o
}

// The class associated to the DNS view.
func (o DnsRrOutput) Class() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRr) pulumi.StringPtrOutput { return v.Class }).(pulumi.StringPtrOutput)
}

// The class parameters associated to the view.
func (o DnsRrOutput) ClassParameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DnsRr) pulumi.StringMapOutput { return v.ClassParameters }).(pulumi.StringMapOutput)
}

// The managed SMART DNS server name, or DNS server name hosting the RR's zone.
func (o DnsRrOutput) Dnsserver() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsRr) pulumi.StringOutput { return v.Dnsserver }).(pulumi.StringOutput)
}

// The View name of the RR to create.
func (o DnsRrOutput) Dnsview() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRr) pulumi.StringPtrOutput { return v.Dnsview }).(pulumi.StringPtrOutput)
}

// The Zone name of the RR to create.
func (o DnsRrOutput) Dnszone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRr) pulumi.StringPtrOutput { return v.Dnszone }).(pulumi.StringPtrOutput)
}

// The Fully Qualified Domain Name of the RR to create.
func (o DnsRrOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsRr) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The DNS Time To Live of the RR to create.
func (o DnsRrOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DnsRr) pulumi.IntPtrOutput { return v.Ttl }).(pulumi.IntPtrOutput)
}

// The type of the RR to create (Supported: A, AAAA, PTR, CNAME, DNAME and NS).
func (o DnsRrOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsRr) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The value od the RR to create.
func (o DnsRrOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsRr) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type DnsRrArrayOutput struct{ *pulumi.OutputState }

func (DnsRrArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsRr)(nil)).Elem()
}

func (o DnsRrArrayOutput) ToDnsRrArrayOutput() DnsRrArrayOutput {
	return o
}

func (o DnsRrArrayOutput) ToDnsRrArrayOutputWithContext(ctx context.Context) DnsRrArrayOutput {
	return o
}

func (o DnsRrArrayOutput) Index(i pulumi.IntInput) DnsRrOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsRr {
		return vs[0].([]*DnsRr)[vs[1].(int)]
	}).(DnsRrOutput)
}

type DnsRrMapOutput struct{ *pulumi.OutputState }

func (DnsRrMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsRr)(nil)).Elem()
}

func (o DnsRrMapOutput) ToDnsRrMapOutput() DnsRrMapOutput {
	return o
}

func (o DnsRrMapOutput) ToDnsRrMapOutputWithContext(ctx context.Context) DnsRrMapOutput {
	return o
}

func (o DnsRrMapOutput) MapIndex(k pulumi.StringInput) DnsRrOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsRr {
		return vs[0].(map[string]*DnsRr)[vs[1].(string)]
	}).(DnsRrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRrInput)(nil)).Elem(), &DnsRr{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRrArrayInput)(nil)).Elem(), DnsRrArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRrMapInput)(nil)).Elem(), DnsRrMap{})
	pulumi.RegisterOutputType(DnsRrOutput{})
	pulumi.RegisterOutputType(DnsRrArrayOutput{})
	pulumi.RegisterOutputType(DnsRrMapOutput{})
}
