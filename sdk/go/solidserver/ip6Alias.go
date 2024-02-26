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

// IPv6 aliases resource allows to create and manage multiple names for a single IP address.
// They are pretty useful to keep IPAM in sync with the DNS handling CNAME(s) from a single repository.
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
//			_, err := solidserver.NewIp6Alias(ctx, "myFirstIP6Alias", &solidserver.Ip6AliasArgs{
//				Address: pulumi.Any(solidserver_ip6_address.MyFirstIP6Address.Address),
//				Space:   pulumi.Any(solidserver_ip_space.MyFirstSpace.Name),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Ip6Alias struct {
	pulumi.CustomResourceState

	// The IPv6 address for which the alias will be associated to.
	Address pulumi.StringOutput `pulumi:"address"`
	// The FQDN of the IPv6 address alias to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the space to which the address belong to.
	Space pulumi.StringOutput `pulumi:"space"`
	// The type of the Alias to create (Supported: A, CNAME; Default: CNAME).
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewIp6Alias registers a new resource with the given unique name, arguments, and options.
func NewIp6Alias(ctx *pulumi.Context,
	name string, args *Ip6AliasArgs, opts ...pulumi.ResourceOption) (*Ip6Alias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.Space == nil {
		return nil, errors.New("invalid value for required argument 'Space'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ip6Alias
	err := ctx.RegisterResource("solidserver:index/ip6Alias:Ip6Alias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIp6Alias gets an existing Ip6Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIp6Alias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ip6AliasState, opts ...pulumi.ResourceOption) (*Ip6Alias, error) {
	var resource Ip6Alias
	err := ctx.ReadResource("solidserver:index/ip6Alias:Ip6Alias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ip6Alias resources.
type ip6AliasState struct {
	// The IPv6 address for which the alias will be associated to.
	Address *string `pulumi:"address"`
	// The FQDN of the IPv6 address alias to create.
	Name *string `pulumi:"name"`
	// The name of the space to which the address belong to.
	Space *string `pulumi:"space"`
	// The type of the Alias to create (Supported: A, CNAME; Default: CNAME).
	Type *string `pulumi:"type"`
}

type Ip6AliasState struct {
	// The IPv6 address for which the alias will be associated to.
	Address pulumi.StringPtrInput
	// The FQDN of the IPv6 address alias to create.
	Name pulumi.StringPtrInput
	// The name of the space to which the address belong to.
	Space pulumi.StringPtrInput
	// The type of the Alias to create (Supported: A, CNAME; Default: CNAME).
	Type pulumi.StringPtrInput
}

func (Ip6AliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*ip6AliasState)(nil)).Elem()
}

type ip6AliasArgs struct {
	// The IPv6 address for which the alias will be associated to.
	Address string `pulumi:"address"`
	// The FQDN of the IPv6 address alias to create.
	Name *string `pulumi:"name"`
	// The name of the space to which the address belong to.
	Space string `pulumi:"space"`
	// The type of the Alias to create (Supported: A, CNAME; Default: CNAME).
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Ip6Alias resource.
type Ip6AliasArgs struct {
	// The IPv6 address for which the alias will be associated to.
	Address pulumi.StringInput
	// The FQDN of the IPv6 address alias to create.
	Name pulumi.StringPtrInput
	// The name of the space to which the address belong to.
	Space pulumi.StringInput
	// The type of the Alias to create (Supported: A, CNAME; Default: CNAME).
	Type pulumi.StringPtrInput
}

func (Ip6AliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ip6AliasArgs)(nil)).Elem()
}

type Ip6AliasInput interface {
	pulumi.Input

	ToIp6AliasOutput() Ip6AliasOutput
	ToIp6AliasOutputWithContext(ctx context.Context) Ip6AliasOutput
}

func (*Ip6Alias) ElementType() reflect.Type {
	return reflect.TypeOf((**Ip6Alias)(nil)).Elem()
}

func (i *Ip6Alias) ToIp6AliasOutput() Ip6AliasOutput {
	return i.ToIp6AliasOutputWithContext(context.Background())
}

func (i *Ip6Alias) ToIp6AliasOutputWithContext(ctx context.Context) Ip6AliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ip6AliasOutput)
}

// Ip6AliasArrayInput is an input type that accepts Ip6AliasArray and Ip6AliasArrayOutput values.
// You can construct a concrete instance of `Ip6AliasArrayInput` via:
//
//	Ip6AliasArray{ Ip6AliasArgs{...} }
type Ip6AliasArrayInput interface {
	pulumi.Input

	ToIp6AliasArrayOutput() Ip6AliasArrayOutput
	ToIp6AliasArrayOutputWithContext(context.Context) Ip6AliasArrayOutput
}

type Ip6AliasArray []Ip6AliasInput

func (Ip6AliasArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ip6Alias)(nil)).Elem()
}

func (i Ip6AliasArray) ToIp6AliasArrayOutput() Ip6AliasArrayOutput {
	return i.ToIp6AliasArrayOutputWithContext(context.Background())
}

func (i Ip6AliasArray) ToIp6AliasArrayOutputWithContext(ctx context.Context) Ip6AliasArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ip6AliasArrayOutput)
}

// Ip6AliasMapInput is an input type that accepts Ip6AliasMap and Ip6AliasMapOutput values.
// You can construct a concrete instance of `Ip6AliasMapInput` via:
//
//	Ip6AliasMap{ "key": Ip6AliasArgs{...} }
type Ip6AliasMapInput interface {
	pulumi.Input

	ToIp6AliasMapOutput() Ip6AliasMapOutput
	ToIp6AliasMapOutputWithContext(context.Context) Ip6AliasMapOutput
}

type Ip6AliasMap map[string]Ip6AliasInput

func (Ip6AliasMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ip6Alias)(nil)).Elem()
}

func (i Ip6AliasMap) ToIp6AliasMapOutput() Ip6AliasMapOutput {
	return i.ToIp6AliasMapOutputWithContext(context.Background())
}

func (i Ip6AliasMap) ToIp6AliasMapOutputWithContext(ctx context.Context) Ip6AliasMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ip6AliasMapOutput)
}

type Ip6AliasOutput struct{ *pulumi.OutputState }

func (Ip6AliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ip6Alias)(nil)).Elem()
}

func (o Ip6AliasOutput) ToIp6AliasOutput() Ip6AliasOutput {
	return o
}

func (o Ip6AliasOutput) ToIp6AliasOutputWithContext(ctx context.Context) Ip6AliasOutput {
	return o
}

// The IPv6 address for which the alias will be associated to.
func (o Ip6AliasOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Ip6Alias) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The FQDN of the IPv6 address alias to create.
func (o Ip6AliasOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ip6Alias) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the space to which the address belong to.
func (o Ip6AliasOutput) Space() pulumi.StringOutput {
	return o.ApplyT(func(v *Ip6Alias) pulumi.StringOutput { return v.Space }).(pulumi.StringOutput)
}

// The type of the Alias to create (Supported: A, CNAME; Default: CNAME).
func (o Ip6AliasOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ip6Alias) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type Ip6AliasArrayOutput struct{ *pulumi.OutputState }

func (Ip6AliasArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ip6Alias)(nil)).Elem()
}

func (o Ip6AliasArrayOutput) ToIp6AliasArrayOutput() Ip6AliasArrayOutput {
	return o
}

func (o Ip6AliasArrayOutput) ToIp6AliasArrayOutputWithContext(ctx context.Context) Ip6AliasArrayOutput {
	return o
}

func (o Ip6AliasArrayOutput) Index(i pulumi.IntInput) Ip6AliasOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ip6Alias {
		return vs[0].([]*Ip6Alias)[vs[1].(int)]
	}).(Ip6AliasOutput)
}

type Ip6AliasMapOutput struct{ *pulumi.OutputState }

func (Ip6AliasMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ip6Alias)(nil)).Elem()
}

func (o Ip6AliasMapOutput) ToIp6AliasMapOutput() Ip6AliasMapOutput {
	return o
}

func (o Ip6AliasMapOutput) ToIp6AliasMapOutputWithContext(ctx context.Context) Ip6AliasMapOutput {
	return o
}

func (o Ip6AliasMapOutput) MapIndex(k pulumi.StringInput) Ip6AliasOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ip6Alias {
		return vs[0].(map[string]*Ip6Alias)[vs[1].(string)]
	}).(Ip6AliasOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Ip6AliasInput)(nil)).Elem(), &Ip6Alias{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ip6AliasArrayInput)(nil)).Elem(), Ip6AliasArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ip6AliasMapInput)(nil)).Elem(), Ip6AliasMap{})
	pulumi.RegisterOutputType(Ip6AliasOutput{})
	pulumi.RegisterOutputType(Ip6AliasArrayOutput{})
	pulumi.RegisterOutputType(Ip6AliasMapOutput{})
}
