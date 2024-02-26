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

// IPv6 MAC resource allows to map an IP address with a MAC address.
// It does not reflect any object within SOLIDserver, it is useful when provisioning
// IP addresses for VM(s) for which the MAC address is unknown until deployed.
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
//			_, err := solidserver.NewIp6Mac(ctx, "myFirstIP6MacAassoc", &solidserver.Ip6MacArgs{
//				Address: pulumi.Any(solidserver_ip6_address.MyFirstIP6Address.Address),
//				Mac:     pulumi.String("06:16:26:36:46:56"),
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
type Ip6Mac struct {
	pulumi.CustomResourceState

	// The IPv6 address to map with the MAC address.
	Address pulumi.StringOutput `pulumi:"address"`
	// The MAC Address o map with the IPv6 address.
	Mac pulumi.StringOutput `pulumi:"mac"`
	// The name of the space into which mapping the IP and the MAC address.
	Space pulumi.StringOutput `pulumi:"space"`
}

// NewIp6Mac registers a new resource with the given unique name, arguments, and options.
func NewIp6Mac(ctx *pulumi.Context,
	name string, args *Ip6MacArgs, opts ...pulumi.ResourceOption) (*Ip6Mac, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.Mac == nil {
		return nil, errors.New("invalid value for required argument 'Mac'")
	}
	if args.Space == nil {
		return nil, errors.New("invalid value for required argument 'Space'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ip6Mac
	err := ctx.RegisterResource("solidserver:index/ip6Mac:Ip6Mac", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIp6Mac gets an existing Ip6Mac resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIp6Mac(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ip6MacState, opts ...pulumi.ResourceOption) (*Ip6Mac, error) {
	var resource Ip6Mac
	err := ctx.ReadResource("solidserver:index/ip6Mac:Ip6Mac", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ip6Mac resources.
type ip6MacState struct {
	// The IPv6 address to map with the MAC address.
	Address *string `pulumi:"address"`
	// The MAC Address o map with the IPv6 address.
	Mac *string `pulumi:"mac"`
	// The name of the space into which mapping the IP and the MAC address.
	Space *string `pulumi:"space"`
}

type Ip6MacState struct {
	// The IPv6 address to map with the MAC address.
	Address pulumi.StringPtrInput
	// The MAC Address o map with the IPv6 address.
	Mac pulumi.StringPtrInput
	// The name of the space into which mapping the IP and the MAC address.
	Space pulumi.StringPtrInput
}

func (Ip6MacState) ElementType() reflect.Type {
	return reflect.TypeOf((*ip6MacState)(nil)).Elem()
}

type ip6MacArgs struct {
	// The IPv6 address to map with the MAC address.
	Address string `pulumi:"address"`
	// The MAC Address o map with the IPv6 address.
	Mac string `pulumi:"mac"`
	// The name of the space into which mapping the IP and the MAC address.
	Space string `pulumi:"space"`
}

// The set of arguments for constructing a Ip6Mac resource.
type Ip6MacArgs struct {
	// The IPv6 address to map with the MAC address.
	Address pulumi.StringInput
	// The MAC Address o map with the IPv6 address.
	Mac pulumi.StringInput
	// The name of the space into which mapping the IP and the MAC address.
	Space pulumi.StringInput
}

func (Ip6MacArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ip6MacArgs)(nil)).Elem()
}

type Ip6MacInput interface {
	pulumi.Input

	ToIp6MacOutput() Ip6MacOutput
	ToIp6MacOutputWithContext(ctx context.Context) Ip6MacOutput
}

func (*Ip6Mac) ElementType() reflect.Type {
	return reflect.TypeOf((**Ip6Mac)(nil)).Elem()
}

func (i *Ip6Mac) ToIp6MacOutput() Ip6MacOutput {
	return i.ToIp6MacOutputWithContext(context.Background())
}

func (i *Ip6Mac) ToIp6MacOutputWithContext(ctx context.Context) Ip6MacOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ip6MacOutput)
}

// Ip6MacArrayInput is an input type that accepts Ip6MacArray and Ip6MacArrayOutput values.
// You can construct a concrete instance of `Ip6MacArrayInput` via:
//
//	Ip6MacArray{ Ip6MacArgs{...} }
type Ip6MacArrayInput interface {
	pulumi.Input

	ToIp6MacArrayOutput() Ip6MacArrayOutput
	ToIp6MacArrayOutputWithContext(context.Context) Ip6MacArrayOutput
}

type Ip6MacArray []Ip6MacInput

func (Ip6MacArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ip6Mac)(nil)).Elem()
}

func (i Ip6MacArray) ToIp6MacArrayOutput() Ip6MacArrayOutput {
	return i.ToIp6MacArrayOutputWithContext(context.Background())
}

func (i Ip6MacArray) ToIp6MacArrayOutputWithContext(ctx context.Context) Ip6MacArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ip6MacArrayOutput)
}

// Ip6MacMapInput is an input type that accepts Ip6MacMap and Ip6MacMapOutput values.
// You can construct a concrete instance of `Ip6MacMapInput` via:
//
//	Ip6MacMap{ "key": Ip6MacArgs{...} }
type Ip6MacMapInput interface {
	pulumi.Input

	ToIp6MacMapOutput() Ip6MacMapOutput
	ToIp6MacMapOutputWithContext(context.Context) Ip6MacMapOutput
}

type Ip6MacMap map[string]Ip6MacInput

func (Ip6MacMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ip6Mac)(nil)).Elem()
}

func (i Ip6MacMap) ToIp6MacMapOutput() Ip6MacMapOutput {
	return i.ToIp6MacMapOutputWithContext(context.Background())
}

func (i Ip6MacMap) ToIp6MacMapOutputWithContext(ctx context.Context) Ip6MacMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ip6MacMapOutput)
}

type Ip6MacOutput struct{ *pulumi.OutputState }

func (Ip6MacOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ip6Mac)(nil)).Elem()
}

func (o Ip6MacOutput) ToIp6MacOutput() Ip6MacOutput {
	return o
}

func (o Ip6MacOutput) ToIp6MacOutputWithContext(ctx context.Context) Ip6MacOutput {
	return o
}

// The IPv6 address to map with the MAC address.
func (o Ip6MacOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Ip6Mac) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The MAC Address o map with the IPv6 address.
func (o Ip6MacOutput) Mac() pulumi.StringOutput {
	return o.ApplyT(func(v *Ip6Mac) pulumi.StringOutput { return v.Mac }).(pulumi.StringOutput)
}

// The name of the space into which mapping the IP and the MAC address.
func (o Ip6MacOutput) Space() pulumi.StringOutput {
	return o.ApplyT(func(v *Ip6Mac) pulumi.StringOutput { return v.Space }).(pulumi.StringOutput)
}

type Ip6MacArrayOutput struct{ *pulumi.OutputState }

func (Ip6MacArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ip6Mac)(nil)).Elem()
}

func (o Ip6MacArrayOutput) ToIp6MacArrayOutput() Ip6MacArrayOutput {
	return o
}

func (o Ip6MacArrayOutput) ToIp6MacArrayOutputWithContext(ctx context.Context) Ip6MacArrayOutput {
	return o
}

func (o Ip6MacArrayOutput) Index(i pulumi.IntInput) Ip6MacOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ip6Mac {
		return vs[0].([]*Ip6Mac)[vs[1].(int)]
	}).(Ip6MacOutput)
}

type Ip6MacMapOutput struct{ *pulumi.OutputState }

func (Ip6MacMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ip6Mac)(nil)).Elem()
}

func (o Ip6MacMapOutput) ToIp6MacMapOutput() Ip6MacMapOutput {
	return o
}

func (o Ip6MacMapOutput) ToIp6MacMapOutputWithContext(ctx context.Context) Ip6MacMapOutput {
	return o
}

func (o Ip6MacMapOutput) MapIndex(k pulumi.StringInput) Ip6MacOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ip6Mac {
		return vs[0].(map[string]*Ip6Mac)[vs[1].(string)]
	}).(Ip6MacOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Ip6MacInput)(nil)).Elem(), &Ip6Mac{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ip6MacArrayInput)(nil)).Elem(), Ip6MacArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ip6MacMapInput)(nil)).Elem(), Ip6MacMap{})
	pulumi.RegisterOutputType(Ip6MacOutput{})
	pulumi.RegisterOutputType(Ip6MacArrayOutput{})
	pulumi.RegisterOutputType(Ip6MacMapOutput{})
}
