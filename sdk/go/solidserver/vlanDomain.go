// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package solidserver

import (
	"context"
	"reflect"

	"github.com/AthenaWolfe/pulumi-solidserver/sdk/go/solidserver/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// VLAN Domain resource allows to create and manage VLAN and VxLAN domains.
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
//			_, err := solidserver.NewVlanDomain(ctx, "myFirstVxlanDomain", &solidserver.VlanDomainArgs{
//				Class: pulumi.String("CUSTOM_VXLAN_DOMAIN"),
//				ClassParameters: pulumi.StringMap{
//					"LOCATION": pulumi.String("PARIS"),
//				},
//				Vxlan: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type VlanDomain struct {
	pulumi.CustomResourceState

	// The class associated to the VLAN Domain.
	Class pulumi.StringPtrOutput `pulumi:"class"`
	// The class parameters associated to VLAN Domain.
	ClassParameters pulumi.StringMapOutput `pulumi:"classParameters"`
	// The name of the VLAN Domain to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specify if the VLAN Domain is a VXLAN Domain.
	Vxlan pulumi.BoolPtrOutput `pulumi:"vxlan"`
}

// NewVlanDomain registers a new resource with the given unique name, arguments, and options.
func NewVlanDomain(ctx *pulumi.Context,
	name string, args *VlanDomainArgs, opts ...pulumi.ResourceOption) (*VlanDomain, error) {
	if args == nil {
		args = &VlanDomainArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VlanDomain
	err := ctx.RegisterResource("solidserver:index/vlanDomain:VlanDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVlanDomain gets an existing VlanDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVlanDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VlanDomainState, opts ...pulumi.ResourceOption) (*VlanDomain, error) {
	var resource VlanDomain
	err := ctx.ReadResource("solidserver:index/vlanDomain:VlanDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VlanDomain resources.
type vlanDomainState struct {
	// The class associated to the VLAN Domain.
	Class *string `pulumi:"class"`
	// The class parameters associated to VLAN Domain.
	ClassParameters map[string]string `pulumi:"classParameters"`
	// The name of the VLAN Domain to create.
	Name *string `pulumi:"name"`
	// Specify if the VLAN Domain is a VXLAN Domain.
	Vxlan *bool `pulumi:"vxlan"`
}

type VlanDomainState struct {
	// The class associated to the VLAN Domain.
	Class pulumi.StringPtrInput
	// The class parameters associated to VLAN Domain.
	ClassParameters pulumi.StringMapInput
	// The name of the VLAN Domain to create.
	Name pulumi.StringPtrInput
	// Specify if the VLAN Domain is a VXLAN Domain.
	Vxlan pulumi.BoolPtrInput
}

func (VlanDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*vlanDomainState)(nil)).Elem()
}

type vlanDomainArgs struct {
	// The class associated to the VLAN Domain.
	Class *string `pulumi:"class"`
	// The class parameters associated to VLAN Domain.
	ClassParameters map[string]string `pulumi:"classParameters"`
	// The name of the VLAN Domain to create.
	Name *string `pulumi:"name"`
	// Specify if the VLAN Domain is a VXLAN Domain.
	Vxlan *bool `pulumi:"vxlan"`
}

// The set of arguments for constructing a VlanDomain resource.
type VlanDomainArgs struct {
	// The class associated to the VLAN Domain.
	Class pulumi.StringPtrInput
	// The class parameters associated to VLAN Domain.
	ClassParameters pulumi.StringMapInput
	// The name of the VLAN Domain to create.
	Name pulumi.StringPtrInput
	// Specify if the VLAN Domain is a VXLAN Domain.
	Vxlan pulumi.BoolPtrInput
}

func (VlanDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vlanDomainArgs)(nil)).Elem()
}

type VlanDomainInput interface {
	pulumi.Input

	ToVlanDomainOutput() VlanDomainOutput
	ToVlanDomainOutputWithContext(ctx context.Context) VlanDomainOutput
}

func (*VlanDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**VlanDomain)(nil)).Elem()
}

func (i *VlanDomain) ToVlanDomainOutput() VlanDomainOutput {
	return i.ToVlanDomainOutputWithContext(context.Background())
}

func (i *VlanDomain) ToVlanDomainOutputWithContext(ctx context.Context) VlanDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanDomainOutput)
}

// VlanDomainArrayInput is an input type that accepts VlanDomainArray and VlanDomainArrayOutput values.
// You can construct a concrete instance of `VlanDomainArrayInput` via:
//
//	VlanDomainArray{ VlanDomainArgs{...} }
type VlanDomainArrayInput interface {
	pulumi.Input

	ToVlanDomainArrayOutput() VlanDomainArrayOutput
	ToVlanDomainArrayOutputWithContext(context.Context) VlanDomainArrayOutput
}

type VlanDomainArray []VlanDomainInput

func (VlanDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VlanDomain)(nil)).Elem()
}

func (i VlanDomainArray) ToVlanDomainArrayOutput() VlanDomainArrayOutput {
	return i.ToVlanDomainArrayOutputWithContext(context.Background())
}

func (i VlanDomainArray) ToVlanDomainArrayOutputWithContext(ctx context.Context) VlanDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanDomainArrayOutput)
}

// VlanDomainMapInput is an input type that accepts VlanDomainMap and VlanDomainMapOutput values.
// You can construct a concrete instance of `VlanDomainMapInput` via:
//
//	VlanDomainMap{ "key": VlanDomainArgs{...} }
type VlanDomainMapInput interface {
	pulumi.Input

	ToVlanDomainMapOutput() VlanDomainMapOutput
	ToVlanDomainMapOutputWithContext(context.Context) VlanDomainMapOutput
}

type VlanDomainMap map[string]VlanDomainInput

func (VlanDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VlanDomain)(nil)).Elem()
}

func (i VlanDomainMap) ToVlanDomainMapOutput() VlanDomainMapOutput {
	return i.ToVlanDomainMapOutputWithContext(context.Background())
}

func (i VlanDomainMap) ToVlanDomainMapOutputWithContext(ctx context.Context) VlanDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanDomainMapOutput)
}

type VlanDomainOutput struct{ *pulumi.OutputState }

func (VlanDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VlanDomain)(nil)).Elem()
}

func (o VlanDomainOutput) ToVlanDomainOutput() VlanDomainOutput {
	return o
}

func (o VlanDomainOutput) ToVlanDomainOutputWithContext(ctx context.Context) VlanDomainOutput {
	return o
}

// The class associated to the VLAN Domain.
func (o VlanDomainOutput) Class() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VlanDomain) pulumi.StringPtrOutput { return v.Class }).(pulumi.StringPtrOutput)
}

// The class parameters associated to VLAN Domain.
func (o VlanDomainOutput) ClassParameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VlanDomain) pulumi.StringMapOutput { return v.ClassParameters }).(pulumi.StringMapOutput)
}

// The name of the VLAN Domain to create.
func (o VlanDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VlanDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specify if the VLAN Domain is a VXLAN Domain.
func (o VlanDomainOutput) Vxlan() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VlanDomain) pulumi.BoolPtrOutput { return v.Vxlan }).(pulumi.BoolPtrOutput)
}

type VlanDomainArrayOutput struct{ *pulumi.OutputState }

func (VlanDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VlanDomain)(nil)).Elem()
}

func (o VlanDomainArrayOutput) ToVlanDomainArrayOutput() VlanDomainArrayOutput {
	return o
}

func (o VlanDomainArrayOutput) ToVlanDomainArrayOutputWithContext(ctx context.Context) VlanDomainArrayOutput {
	return o
}

func (o VlanDomainArrayOutput) Index(i pulumi.IntInput) VlanDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VlanDomain {
		return vs[0].([]*VlanDomain)[vs[1].(int)]
	}).(VlanDomainOutput)
}

type VlanDomainMapOutput struct{ *pulumi.OutputState }

func (VlanDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VlanDomain)(nil)).Elem()
}

func (o VlanDomainMapOutput) ToVlanDomainMapOutput() VlanDomainMapOutput {
	return o
}

func (o VlanDomainMapOutput) ToVlanDomainMapOutputWithContext(ctx context.Context) VlanDomainMapOutput {
	return o
}

func (o VlanDomainMapOutput) MapIndex(k pulumi.StringInput) VlanDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VlanDomain {
		return vs[0].(map[string]*VlanDomain)[vs[1].(string)]
	}).(VlanDomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VlanDomainInput)(nil)).Elem(), &VlanDomain{})
	pulumi.RegisterInputType(reflect.TypeOf((*VlanDomainArrayInput)(nil)).Elem(), VlanDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VlanDomainMapInput)(nil)).Elem(), VlanDomainMap{})
	pulumi.RegisterOutputType(VlanDomainOutput{})
	pulumi.RegisterOutputType(VlanDomainArrayOutput{})
	pulumi.RegisterOutputType(VlanDomainMapOutput{})
}
