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

// VLANresource allows to create and manage VLAN(s) and VxLAN(s).
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
//			_, err := solidserver.NewVlan(ctx, "myFirstVxlan", &solidserver.VlanArgs{
//				VlanDomain: pulumi.Any(solidserver_vlan_domain.MyFirstVxlanDomain.Name),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Vlan struct {
	pulumi.CustomResourceState

	// The class associated to the vlan.
	Class pulumi.StringPtrOutput `pulumi:"class"`
	// The class parameters associated to vlan.
	ClassParameters pulumi.StringMapOutput `pulumi:"classParameters"`
	// The name of the vlan to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// The optionally requested vlan ID.
	RequestId pulumi.IntPtrOutput `pulumi:"requestId"`
	// The name of the vlan Domain.
	VlanDomain pulumi.StringOutput `pulumi:"vlanDomain"`
	// The vlan ID.
	VlanId pulumi.IntOutput `pulumi:"vlanId"`
	// The name of the vlan Range.
	VlanRange pulumi.StringPtrOutput `pulumi:"vlanRange"`
}

// NewVlan registers a new resource with the given unique name, arguments, and options.
func NewVlan(ctx *pulumi.Context,
	name string, args *VlanArgs, opts ...pulumi.ResourceOption) (*Vlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VlanDomain == nil {
		return nil, errors.New("invalid value for required argument 'VlanDomain'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Vlan
	err := ctx.RegisterResource("solidserver:index/vlan:Vlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVlan gets an existing Vlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VlanState, opts ...pulumi.ResourceOption) (*Vlan, error) {
	var resource Vlan
	err := ctx.ReadResource("solidserver:index/vlan:Vlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vlan resources.
type vlanState struct {
	// The class associated to the vlan.
	Class *string `pulumi:"class"`
	// The class parameters associated to vlan.
	ClassParameters map[string]string `pulumi:"classParameters"`
	// The name of the vlan to create.
	Name *string `pulumi:"name"`
	// The optionally requested vlan ID.
	RequestId *int `pulumi:"requestId"`
	// The name of the vlan Domain.
	VlanDomain *string `pulumi:"vlanDomain"`
	// The vlan ID.
	VlanId *int `pulumi:"vlanId"`
	// The name of the vlan Range.
	VlanRange *string `pulumi:"vlanRange"`
}

type VlanState struct {
	// The class associated to the vlan.
	Class pulumi.StringPtrInput
	// The class parameters associated to vlan.
	ClassParameters pulumi.StringMapInput
	// The name of the vlan to create.
	Name pulumi.StringPtrInput
	// The optionally requested vlan ID.
	RequestId pulumi.IntPtrInput
	// The name of the vlan Domain.
	VlanDomain pulumi.StringPtrInput
	// The vlan ID.
	VlanId pulumi.IntPtrInput
	// The name of the vlan Range.
	VlanRange pulumi.StringPtrInput
}

func (VlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*vlanState)(nil)).Elem()
}

type vlanArgs struct {
	// The class associated to the vlan.
	Class *string `pulumi:"class"`
	// The class parameters associated to vlan.
	ClassParameters map[string]string `pulumi:"classParameters"`
	// The name of the vlan to create.
	Name *string `pulumi:"name"`
	// The optionally requested vlan ID.
	RequestId *int `pulumi:"requestId"`
	// The name of the vlan Domain.
	VlanDomain string `pulumi:"vlanDomain"`
	// The name of the vlan Range.
	VlanRange *string `pulumi:"vlanRange"`
}

// The set of arguments for constructing a Vlan resource.
type VlanArgs struct {
	// The class associated to the vlan.
	Class pulumi.StringPtrInput
	// The class parameters associated to vlan.
	ClassParameters pulumi.StringMapInput
	// The name of the vlan to create.
	Name pulumi.StringPtrInput
	// The optionally requested vlan ID.
	RequestId pulumi.IntPtrInput
	// The name of the vlan Domain.
	VlanDomain pulumi.StringInput
	// The name of the vlan Range.
	VlanRange pulumi.StringPtrInput
}

func (VlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vlanArgs)(nil)).Elem()
}

type VlanInput interface {
	pulumi.Input

	ToVlanOutput() VlanOutput
	ToVlanOutputWithContext(ctx context.Context) VlanOutput
}

func (*Vlan) ElementType() reflect.Type {
	return reflect.TypeOf((**Vlan)(nil)).Elem()
}

func (i *Vlan) ToVlanOutput() VlanOutput {
	return i.ToVlanOutputWithContext(context.Background())
}

func (i *Vlan) ToVlanOutputWithContext(ctx context.Context) VlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanOutput)
}

// VlanArrayInput is an input type that accepts VlanArray and VlanArrayOutput values.
// You can construct a concrete instance of `VlanArrayInput` via:
//
//	VlanArray{ VlanArgs{...} }
type VlanArrayInput interface {
	pulumi.Input

	ToVlanArrayOutput() VlanArrayOutput
	ToVlanArrayOutputWithContext(context.Context) VlanArrayOutput
}

type VlanArray []VlanInput

func (VlanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vlan)(nil)).Elem()
}

func (i VlanArray) ToVlanArrayOutput() VlanArrayOutput {
	return i.ToVlanArrayOutputWithContext(context.Background())
}

func (i VlanArray) ToVlanArrayOutputWithContext(ctx context.Context) VlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanArrayOutput)
}

// VlanMapInput is an input type that accepts VlanMap and VlanMapOutput values.
// You can construct a concrete instance of `VlanMapInput` via:
//
//	VlanMap{ "key": VlanArgs{...} }
type VlanMapInput interface {
	pulumi.Input

	ToVlanMapOutput() VlanMapOutput
	ToVlanMapOutputWithContext(context.Context) VlanMapOutput
}

type VlanMap map[string]VlanInput

func (VlanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vlan)(nil)).Elem()
}

func (i VlanMap) ToVlanMapOutput() VlanMapOutput {
	return i.ToVlanMapOutputWithContext(context.Background())
}

func (i VlanMap) ToVlanMapOutputWithContext(ctx context.Context) VlanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanMapOutput)
}

type VlanOutput struct{ *pulumi.OutputState }

func (VlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vlan)(nil)).Elem()
}

func (o VlanOutput) ToVlanOutput() VlanOutput {
	return o
}

func (o VlanOutput) ToVlanOutputWithContext(ctx context.Context) VlanOutput {
	return o
}

// The class associated to the vlan.
func (o VlanOutput) Class() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vlan) pulumi.StringPtrOutput { return v.Class }).(pulumi.StringPtrOutput)
}

// The class parameters associated to vlan.
func (o VlanOutput) ClassParameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Vlan) pulumi.StringMapOutput { return v.ClassParameters }).(pulumi.StringMapOutput)
}

// The name of the vlan to create.
func (o VlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Vlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The optionally requested vlan ID.
func (o VlanOutput) RequestId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Vlan) pulumi.IntPtrOutput { return v.RequestId }).(pulumi.IntPtrOutput)
}

// The name of the vlan Domain.
func (o VlanOutput) VlanDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *Vlan) pulumi.StringOutput { return v.VlanDomain }).(pulumi.StringOutput)
}

// The vlan ID.
func (o VlanOutput) VlanId() pulumi.IntOutput {
	return o.ApplyT(func(v *Vlan) pulumi.IntOutput { return v.VlanId }).(pulumi.IntOutput)
}

// The name of the vlan Range.
func (o VlanOutput) VlanRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vlan) pulumi.StringPtrOutput { return v.VlanRange }).(pulumi.StringPtrOutput)
}

type VlanArrayOutput struct{ *pulumi.OutputState }

func (VlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vlan)(nil)).Elem()
}

func (o VlanArrayOutput) ToVlanArrayOutput() VlanArrayOutput {
	return o
}

func (o VlanArrayOutput) ToVlanArrayOutputWithContext(ctx context.Context) VlanArrayOutput {
	return o
}

func (o VlanArrayOutput) Index(i pulumi.IntInput) VlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vlan {
		return vs[0].([]*Vlan)[vs[1].(int)]
	}).(VlanOutput)
}

type VlanMapOutput struct{ *pulumi.OutputState }

func (VlanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vlan)(nil)).Elem()
}

func (o VlanMapOutput) ToVlanMapOutput() VlanMapOutput {
	return o
}

func (o VlanMapOutput) ToVlanMapOutputWithContext(ctx context.Context) VlanMapOutput {
	return o
}

func (o VlanMapOutput) MapIndex(k pulumi.StringInput) VlanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vlan {
		return vs[0].(map[string]*Vlan)[vs[1].(string)]
	}).(VlanOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VlanInput)(nil)).Elem(), &Vlan{})
	pulumi.RegisterInputType(reflect.TypeOf((*VlanArrayInput)(nil)).Elem(), VlanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VlanMapInput)(nil)).Elem(), VlanMap{})
	pulumi.RegisterOutputType(VlanOutput{})
	pulumi.RegisterOutputType(VlanArrayOutput{})
	pulumi.RegisterOutputType(VlanMapOutput{})
}
