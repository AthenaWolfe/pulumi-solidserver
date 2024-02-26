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

// IP Subnet resource allows to create and manage IPAM networks that are key to organize the IP space
// Subnet can be blocks or subnets. Blocks reflect the assigned IP ranges (RFC1918 or public prefixes).
// Subnets reflect the internal sub-division of your network.
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
//			myFirstIPBlock, err := solidserver.NewIpSubnet(ctx, "myFirstIPBlock", &solidserver.IpSubnetArgs{
//				PrefixSize: pulumi.Int(8),
//				RequestIp:  pulumi.String("10.0.0.0"),
//				Space:      pulumi.Any(solidserver_ip_space.MyFirstSpace.Name),
//				Terminal:   pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = solidserver.NewIpSubnet(ctx, "myFirstIPSubnet", &solidserver.IpSubnetArgs{
//				Block: myFirstIPBlock.Name,
//				Class: pulumi.String("VIRTUAL"),
//				ClassParameters: pulumi.StringMap{
//					"vnid": pulumi.String("12666"),
//				},
//				GatewayOffset: -1,
//				PrefixSize:    pulumi.Int(24),
//				Space:         pulumi.Any(solidserver_ip_space.MyFirstSpace.Name),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type IpSubnet struct {
	pulumi.CustomResourceState

	// The provisionned IP network address.
	Address pulumi.StringOutput `pulumi:"address"`
	// The name of the parent IP block/subnet into which creating the IP subnet.
	Block pulumi.StringPtrOutput `pulumi:"block"`
	// The class associated to the IP subnet.
	Class pulumi.StringPtrOutput `pulumi:"class"`
	// The class parameters associated to the IP subnet.
	ClassParameters pulumi.StringMapOutput `pulumi:"classParameters"`
	// The subnet's computed gateway.
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// Offset for creating the gateway. Default is 0 (No gateway).
	GatewayOffset pulumi.IntPtrOutput `pulumi:"gatewayOffset"`
	// The name of the IP subnet to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisionned IP address netmask.
	Netmask pulumi.StringOutput `pulumi:"netmask"`
	// The provisionned IP prefix.
	Prefix pulumi.StringOutput `pulumi:"prefix"`
	// The expected IP subnet's prefix length (ex: 24 for a '/24').
	PrefixSize pulumi.IntOutput `pulumi:"prefixSize"`
	// The optionally requested subnet IP address.
	RequestIp pulumi.StringPtrOutput `pulumi:"requestIp"`
	// The name of the space into which creating the subnet.
	Space pulumi.StringOutput `pulumi:"space"`
	// The terminal property of the IP subnet.
	Terminal pulumi.BoolPtrOutput `pulumi:"terminal"`
}

// NewIpSubnet registers a new resource with the given unique name, arguments, and options.
func NewIpSubnet(ctx *pulumi.Context,
	name string, args *IpSubnetArgs, opts ...pulumi.ResourceOption) (*IpSubnet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrefixSize == nil {
		return nil, errors.New("invalid value for required argument 'PrefixSize'")
	}
	if args.Space == nil {
		return nil, errors.New("invalid value for required argument 'Space'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpSubnet
	err := ctx.RegisterResource("solidserver:index/ipSubnet:IpSubnet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpSubnet gets an existing IpSubnet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpSubnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpSubnetState, opts ...pulumi.ResourceOption) (*IpSubnet, error) {
	var resource IpSubnet
	err := ctx.ReadResource("solidserver:index/ipSubnet:IpSubnet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpSubnet resources.
type ipSubnetState struct {
	// The provisionned IP network address.
	Address *string `pulumi:"address"`
	// The name of the parent IP block/subnet into which creating the IP subnet.
	Block *string `pulumi:"block"`
	// The class associated to the IP subnet.
	Class *string `pulumi:"class"`
	// The class parameters associated to the IP subnet.
	ClassParameters map[string]string `pulumi:"classParameters"`
	// The subnet's computed gateway.
	Gateway *string `pulumi:"gateway"`
	// Offset for creating the gateway. Default is 0 (No gateway).
	GatewayOffset *int `pulumi:"gatewayOffset"`
	// The name of the IP subnet to create.
	Name *string `pulumi:"name"`
	// The provisionned IP address netmask.
	Netmask *string `pulumi:"netmask"`
	// The provisionned IP prefix.
	Prefix *string `pulumi:"prefix"`
	// The expected IP subnet's prefix length (ex: 24 for a '/24').
	PrefixSize *int `pulumi:"prefixSize"`
	// The optionally requested subnet IP address.
	RequestIp *string `pulumi:"requestIp"`
	// The name of the space into which creating the subnet.
	Space *string `pulumi:"space"`
	// The terminal property of the IP subnet.
	Terminal *bool `pulumi:"terminal"`
}

type IpSubnetState struct {
	// The provisionned IP network address.
	Address pulumi.StringPtrInput
	// The name of the parent IP block/subnet into which creating the IP subnet.
	Block pulumi.StringPtrInput
	// The class associated to the IP subnet.
	Class pulumi.StringPtrInput
	// The class parameters associated to the IP subnet.
	ClassParameters pulumi.StringMapInput
	// The subnet's computed gateway.
	Gateway pulumi.StringPtrInput
	// Offset for creating the gateway. Default is 0 (No gateway).
	GatewayOffset pulumi.IntPtrInput
	// The name of the IP subnet to create.
	Name pulumi.StringPtrInput
	// The provisionned IP address netmask.
	Netmask pulumi.StringPtrInput
	// The provisionned IP prefix.
	Prefix pulumi.StringPtrInput
	// The expected IP subnet's prefix length (ex: 24 for a '/24').
	PrefixSize pulumi.IntPtrInput
	// The optionally requested subnet IP address.
	RequestIp pulumi.StringPtrInput
	// The name of the space into which creating the subnet.
	Space pulumi.StringPtrInput
	// The terminal property of the IP subnet.
	Terminal pulumi.BoolPtrInput
}

func (IpSubnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipSubnetState)(nil)).Elem()
}

type ipSubnetArgs struct {
	// The name of the parent IP block/subnet into which creating the IP subnet.
	Block *string `pulumi:"block"`
	// The class associated to the IP subnet.
	Class *string `pulumi:"class"`
	// The class parameters associated to the IP subnet.
	ClassParameters map[string]string `pulumi:"classParameters"`
	// Offset for creating the gateway. Default is 0 (No gateway).
	GatewayOffset *int `pulumi:"gatewayOffset"`
	// The name of the IP subnet to create.
	Name *string `pulumi:"name"`
	// The expected IP subnet's prefix length (ex: 24 for a '/24').
	PrefixSize int `pulumi:"prefixSize"`
	// The optionally requested subnet IP address.
	RequestIp *string `pulumi:"requestIp"`
	// The name of the space into which creating the subnet.
	Space string `pulumi:"space"`
	// The terminal property of the IP subnet.
	Terminal *bool `pulumi:"terminal"`
}

// The set of arguments for constructing a IpSubnet resource.
type IpSubnetArgs struct {
	// The name of the parent IP block/subnet into which creating the IP subnet.
	Block pulumi.StringPtrInput
	// The class associated to the IP subnet.
	Class pulumi.StringPtrInput
	// The class parameters associated to the IP subnet.
	ClassParameters pulumi.StringMapInput
	// Offset for creating the gateway. Default is 0 (No gateway).
	GatewayOffset pulumi.IntPtrInput
	// The name of the IP subnet to create.
	Name pulumi.StringPtrInput
	// The expected IP subnet's prefix length (ex: 24 for a '/24').
	PrefixSize pulumi.IntInput
	// The optionally requested subnet IP address.
	RequestIp pulumi.StringPtrInput
	// The name of the space into which creating the subnet.
	Space pulumi.StringInput
	// The terminal property of the IP subnet.
	Terminal pulumi.BoolPtrInput
}

func (IpSubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipSubnetArgs)(nil)).Elem()
}

type IpSubnetInput interface {
	pulumi.Input

	ToIpSubnetOutput() IpSubnetOutput
	ToIpSubnetOutputWithContext(ctx context.Context) IpSubnetOutput
}

func (*IpSubnet) ElementType() reflect.Type {
	return reflect.TypeOf((**IpSubnet)(nil)).Elem()
}

func (i *IpSubnet) ToIpSubnetOutput() IpSubnetOutput {
	return i.ToIpSubnetOutputWithContext(context.Background())
}

func (i *IpSubnet) ToIpSubnetOutputWithContext(ctx context.Context) IpSubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSubnetOutput)
}

// IpSubnetArrayInput is an input type that accepts IpSubnetArray and IpSubnetArrayOutput values.
// You can construct a concrete instance of `IpSubnetArrayInput` via:
//
//	IpSubnetArray{ IpSubnetArgs{...} }
type IpSubnetArrayInput interface {
	pulumi.Input

	ToIpSubnetArrayOutput() IpSubnetArrayOutput
	ToIpSubnetArrayOutputWithContext(context.Context) IpSubnetArrayOutput
}

type IpSubnetArray []IpSubnetInput

func (IpSubnetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpSubnet)(nil)).Elem()
}

func (i IpSubnetArray) ToIpSubnetArrayOutput() IpSubnetArrayOutput {
	return i.ToIpSubnetArrayOutputWithContext(context.Background())
}

func (i IpSubnetArray) ToIpSubnetArrayOutputWithContext(ctx context.Context) IpSubnetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSubnetArrayOutput)
}

// IpSubnetMapInput is an input type that accepts IpSubnetMap and IpSubnetMapOutput values.
// You can construct a concrete instance of `IpSubnetMapInput` via:
//
//	IpSubnetMap{ "key": IpSubnetArgs{...} }
type IpSubnetMapInput interface {
	pulumi.Input

	ToIpSubnetMapOutput() IpSubnetMapOutput
	ToIpSubnetMapOutputWithContext(context.Context) IpSubnetMapOutput
}

type IpSubnetMap map[string]IpSubnetInput

func (IpSubnetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpSubnet)(nil)).Elem()
}

func (i IpSubnetMap) ToIpSubnetMapOutput() IpSubnetMapOutput {
	return i.ToIpSubnetMapOutputWithContext(context.Background())
}

func (i IpSubnetMap) ToIpSubnetMapOutputWithContext(ctx context.Context) IpSubnetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSubnetMapOutput)
}

type IpSubnetOutput struct{ *pulumi.OutputState }

func (IpSubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpSubnet)(nil)).Elem()
}

func (o IpSubnetOutput) ToIpSubnetOutput() IpSubnetOutput {
	return o
}

func (o IpSubnetOutput) ToIpSubnetOutputWithContext(ctx context.Context) IpSubnetOutput {
	return o
}

// The provisionned IP network address.
func (o IpSubnetOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSubnet) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The name of the parent IP block/subnet into which creating the IP subnet.
func (o IpSubnetOutput) Block() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpSubnet) pulumi.StringPtrOutput { return v.Block }).(pulumi.StringPtrOutput)
}

// The class associated to the IP subnet.
func (o IpSubnetOutput) Class() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpSubnet) pulumi.StringPtrOutput { return v.Class }).(pulumi.StringPtrOutput)
}

// The class parameters associated to the IP subnet.
func (o IpSubnetOutput) ClassParameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IpSubnet) pulumi.StringMapOutput { return v.ClassParameters }).(pulumi.StringMapOutput)
}

// The subnet's computed gateway.
func (o IpSubnetOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSubnet) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// Offset for creating the gateway. Default is 0 (No gateway).
func (o IpSubnetOutput) GatewayOffset() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IpSubnet) pulumi.IntPtrOutput { return v.GatewayOffset }).(pulumi.IntPtrOutput)
}

// The name of the IP subnet to create.
func (o IpSubnetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSubnet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisionned IP address netmask.
func (o IpSubnetOutput) Netmask() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSubnet) pulumi.StringOutput { return v.Netmask }).(pulumi.StringOutput)
}

// The provisionned IP prefix.
func (o IpSubnetOutput) Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSubnet) pulumi.StringOutput { return v.Prefix }).(pulumi.StringOutput)
}

// The expected IP subnet's prefix length (ex: 24 for a '/24').
func (o IpSubnetOutput) PrefixSize() pulumi.IntOutput {
	return o.ApplyT(func(v *IpSubnet) pulumi.IntOutput { return v.PrefixSize }).(pulumi.IntOutput)
}

// The optionally requested subnet IP address.
func (o IpSubnetOutput) RequestIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpSubnet) pulumi.StringPtrOutput { return v.RequestIp }).(pulumi.StringPtrOutput)
}

// The name of the space into which creating the subnet.
func (o IpSubnetOutput) Space() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSubnet) pulumi.StringOutput { return v.Space }).(pulumi.StringOutput)
}

// The terminal property of the IP subnet.
func (o IpSubnetOutput) Terminal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IpSubnet) pulumi.BoolPtrOutput { return v.Terminal }).(pulumi.BoolPtrOutput)
}

type IpSubnetArrayOutput struct{ *pulumi.OutputState }

func (IpSubnetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpSubnet)(nil)).Elem()
}

func (o IpSubnetArrayOutput) ToIpSubnetArrayOutput() IpSubnetArrayOutput {
	return o
}

func (o IpSubnetArrayOutput) ToIpSubnetArrayOutputWithContext(ctx context.Context) IpSubnetArrayOutput {
	return o
}

func (o IpSubnetArrayOutput) Index(i pulumi.IntInput) IpSubnetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpSubnet {
		return vs[0].([]*IpSubnet)[vs[1].(int)]
	}).(IpSubnetOutput)
}

type IpSubnetMapOutput struct{ *pulumi.OutputState }

func (IpSubnetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpSubnet)(nil)).Elem()
}

func (o IpSubnetMapOutput) ToIpSubnetMapOutput() IpSubnetMapOutput {
	return o
}

func (o IpSubnetMapOutput) ToIpSubnetMapOutputWithContext(ctx context.Context) IpSubnetMapOutput {
	return o
}

func (o IpSubnetMapOutput) MapIndex(k pulumi.StringInput) IpSubnetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpSubnet {
		return vs[0].(map[string]*IpSubnet)[vs[1].(string)]
	}).(IpSubnetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpSubnetInput)(nil)).Elem(), &IpSubnet{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpSubnetArrayInput)(nil)).Elem(), IpSubnetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpSubnetMapInput)(nil)).Elem(), IpSubnetMap{})
	pulumi.RegisterOutputType(IpSubnetOutput{})
	pulumi.RegisterOutputType(IpSubnetArrayOutput{})
	pulumi.RegisterOutputType(IpSubnetMapOutput{})
}