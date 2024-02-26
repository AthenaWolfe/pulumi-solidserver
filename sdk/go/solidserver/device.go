// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package solidserver

import (
	"context"
	"reflect"

	"github.com/AthenaWolfe/pulumi-solidserver/sdk/go/solidserver/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Device resource allows to create and manage network devices and link them with IP addresses.
// It is pretty useful to link together the many IP addresses a device can rely on.
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
//			_, err := solidserver.NewDevice(ctx, "myFirstDevice", &solidserver.DeviceArgs{
//				Class: pulumi.String("AWS_EC2_INSTANCE"),
//				ClassParameters: pulumi.StringMap{
//					"cloudaz":      pulumi.String("eu-west-1a"),
//					"instanceid":   pulumi.String("i-03d4bd36f915b0322"),
//					"instancetype": pulumi.String("t2.micro"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Device struct {
	pulumi.CustomResourceState

	// The class associated to the device.
	Class pulumi.StringPtrOutput `pulumi:"class"`
	// The class parameters associated to device.
	ClassParameters pulumi.StringMapOutput `pulumi:"classParameters"`
	// The name of the device to create.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewDevice registers a new resource with the given unique name, arguments, and options.
func NewDevice(ctx *pulumi.Context,
	name string, args *DeviceArgs, opts ...pulumi.ResourceOption) (*Device, error) {
	if args == nil {
		args = &DeviceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Device
	err := ctx.RegisterResource("solidserver:index/device:Device", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevice gets an existing Device resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceState, opts ...pulumi.ResourceOption) (*Device, error) {
	var resource Device
	err := ctx.ReadResource("solidserver:index/device:Device", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Device resources.
type deviceState struct {
	// The class associated to the device.
	Class *string `pulumi:"class"`
	// The class parameters associated to device.
	ClassParameters map[string]string `pulumi:"classParameters"`
	// The name of the device to create.
	Name *string `pulumi:"name"`
}

type DeviceState struct {
	// The class associated to the device.
	Class pulumi.StringPtrInput
	// The class parameters associated to device.
	ClassParameters pulumi.StringMapInput
	// The name of the device to create.
	Name pulumi.StringPtrInput
}

func (DeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceState)(nil)).Elem()
}

type deviceArgs struct {
	// The class associated to the device.
	Class *string `pulumi:"class"`
	// The class parameters associated to device.
	ClassParameters map[string]string `pulumi:"classParameters"`
	// The name of the device to create.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Device resource.
type DeviceArgs struct {
	// The class associated to the device.
	Class pulumi.StringPtrInput
	// The class parameters associated to device.
	ClassParameters pulumi.StringMapInput
	// The name of the device to create.
	Name pulumi.StringPtrInput
}

func (DeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceArgs)(nil)).Elem()
}

type DeviceInput interface {
	pulumi.Input

	ToDeviceOutput() DeviceOutput
	ToDeviceOutputWithContext(ctx context.Context) DeviceOutput
}

func (*Device) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil)).Elem()
}

func (i *Device) ToDeviceOutput() DeviceOutput {
	return i.ToDeviceOutputWithContext(context.Background())
}

func (i *Device) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceOutput)
}

// DeviceArrayInput is an input type that accepts DeviceArray and DeviceArrayOutput values.
// You can construct a concrete instance of `DeviceArrayInput` via:
//
//	DeviceArray{ DeviceArgs{...} }
type DeviceArrayInput interface {
	pulumi.Input

	ToDeviceArrayOutput() DeviceArrayOutput
	ToDeviceArrayOutputWithContext(context.Context) DeviceArrayOutput
}

type DeviceArray []DeviceInput

func (DeviceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Device)(nil)).Elem()
}

func (i DeviceArray) ToDeviceArrayOutput() DeviceArrayOutput {
	return i.ToDeviceArrayOutputWithContext(context.Background())
}

func (i DeviceArray) ToDeviceArrayOutputWithContext(ctx context.Context) DeviceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceArrayOutput)
}

// DeviceMapInput is an input type that accepts DeviceMap and DeviceMapOutput values.
// You can construct a concrete instance of `DeviceMapInput` via:
//
//	DeviceMap{ "key": DeviceArgs{...} }
type DeviceMapInput interface {
	pulumi.Input

	ToDeviceMapOutput() DeviceMapOutput
	ToDeviceMapOutputWithContext(context.Context) DeviceMapOutput
}

type DeviceMap map[string]DeviceInput

func (DeviceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Device)(nil)).Elem()
}

func (i DeviceMap) ToDeviceMapOutput() DeviceMapOutput {
	return i.ToDeviceMapOutputWithContext(context.Background())
}

func (i DeviceMap) ToDeviceMapOutputWithContext(ctx context.Context) DeviceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceMapOutput)
}

type DeviceOutput struct{ *pulumi.OutputState }

func (DeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil)).Elem()
}

func (o DeviceOutput) ToDeviceOutput() DeviceOutput {
	return o
}

func (o DeviceOutput) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return o
}

// The class associated to the device.
func (o DeviceOutput) Class() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.StringPtrOutput { return v.Class }).(pulumi.StringPtrOutput)
}

// The class parameters associated to device.
func (o DeviceOutput) ClassParameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Device) pulumi.StringMapOutput { return v.ClassParameters }).(pulumi.StringMapOutput)
}

// The name of the device to create.
func (o DeviceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type DeviceArrayOutput struct{ *pulumi.OutputState }

func (DeviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Device)(nil)).Elem()
}

func (o DeviceArrayOutput) ToDeviceArrayOutput() DeviceArrayOutput {
	return o
}

func (o DeviceArrayOutput) ToDeviceArrayOutputWithContext(ctx context.Context) DeviceArrayOutput {
	return o
}

func (o DeviceArrayOutput) Index(i pulumi.IntInput) DeviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Device {
		return vs[0].([]*Device)[vs[1].(int)]
	}).(DeviceOutput)
}

type DeviceMapOutput struct{ *pulumi.OutputState }

func (DeviceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Device)(nil)).Elem()
}

func (o DeviceMapOutput) ToDeviceMapOutput() DeviceMapOutput {
	return o
}

func (o DeviceMapOutput) ToDeviceMapOutputWithContext(ctx context.Context) DeviceMapOutput {
	return o
}

func (o DeviceMapOutput) MapIndex(k pulumi.StringInput) DeviceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Device {
		return vs[0].(map[string]*Device)[vs[1].(string)]
	}).(DeviceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceInput)(nil)).Elem(), &Device{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceArrayInput)(nil)).Elem(), DeviceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceMapInput)(nil)).Elem(), DeviceMap{})
	pulumi.RegisterOutputType(DeviceOutput{})
	pulumi.RegisterOutputType(DeviceArrayOutput{})
	pulumi.RegisterOutputType(DeviceMapOutput{})
}
