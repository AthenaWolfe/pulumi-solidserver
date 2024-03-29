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

// Custom DB Data resource allows to create and manage custom database entries stored within SOLIDserver.
// This custom database entries can be leveraged within object classes and wizards in order to store custom meta-data.
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
//			_, err := solidserver.NewCdbData(ctx, "myFirstCustomData", &solidserver.CdbDataArgs{
//				CustomDb: pulumi.String("myFirstCustomDB"),
//				Value1:   pulumi.String("FR"),
//				Value2:   pulumi.String("France"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type CdbData struct {
	pulumi.CustomResourceState

	// The name of the Custom DB into which creating the data.
	CustomDb pulumi.StringOutput `pulumi:"customDb"`
	// The value 1 (key of the data)
	Value1 pulumi.StringOutput `pulumi:"value1"`
	// The value 10
	Value10 pulumi.StringPtrOutput `pulumi:"value10"`
	// The value 2
	Value2 pulumi.StringPtrOutput `pulumi:"value2"`
	// The value 3
	Value3 pulumi.StringPtrOutput `pulumi:"value3"`
	// The value 4
	Value4 pulumi.StringPtrOutput `pulumi:"value4"`
	// The value 5
	Value5 pulumi.StringPtrOutput `pulumi:"value5"`
	// The value 6
	Value6 pulumi.StringPtrOutput `pulumi:"value6"`
	// The value 7
	Value7 pulumi.StringPtrOutput `pulumi:"value7"`
	// The value 8
	Value8 pulumi.StringPtrOutput `pulumi:"value8"`
	// The value 9
	Value9 pulumi.StringPtrOutput `pulumi:"value9"`
}

// NewCdbData registers a new resource with the given unique name, arguments, and options.
func NewCdbData(ctx *pulumi.Context,
	name string, args *CdbDataArgs, opts ...pulumi.ResourceOption) (*CdbData, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CustomDb == nil {
		return nil, errors.New("invalid value for required argument 'CustomDb'")
	}
	if args.Value1 == nil {
		return nil, errors.New("invalid value for required argument 'Value1'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CdbData
	err := ctx.RegisterResource("solidserver:index/cdbData:CdbData", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCdbData gets an existing CdbData resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCdbData(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CdbDataState, opts ...pulumi.ResourceOption) (*CdbData, error) {
	var resource CdbData
	err := ctx.ReadResource("solidserver:index/cdbData:CdbData", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CdbData resources.
type cdbDataState struct {
	// The name of the Custom DB into which creating the data.
	CustomDb *string `pulumi:"customDb"`
	// The value 1 (key of the data)
	Value1 *string `pulumi:"value1"`
	// The value 10
	Value10 *string `pulumi:"value10"`
	// The value 2
	Value2 *string `pulumi:"value2"`
	// The value 3
	Value3 *string `pulumi:"value3"`
	// The value 4
	Value4 *string `pulumi:"value4"`
	// The value 5
	Value5 *string `pulumi:"value5"`
	// The value 6
	Value6 *string `pulumi:"value6"`
	// The value 7
	Value7 *string `pulumi:"value7"`
	// The value 8
	Value8 *string `pulumi:"value8"`
	// The value 9
	Value9 *string `pulumi:"value9"`
}

type CdbDataState struct {
	// The name of the Custom DB into which creating the data.
	CustomDb pulumi.StringPtrInput
	// The value 1 (key of the data)
	Value1 pulumi.StringPtrInput
	// The value 10
	Value10 pulumi.StringPtrInput
	// The value 2
	Value2 pulumi.StringPtrInput
	// The value 3
	Value3 pulumi.StringPtrInput
	// The value 4
	Value4 pulumi.StringPtrInput
	// The value 5
	Value5 pulumi.StringPtrInput
	// The value 6
	Value6 pulumi.StringPtrInput
	// The value 7
	Value7 pulumi.StringPtrInput
	// The value 8
	Value8 pulumi.StringPtrInput
	// The value 9
	Value9 pulumi.StringPtrInput
}

func (CdbDataState) ElementType() reflect.Type {
	return reflect.TypeOf((*cdbDataState)(nil)).Elem()
}

type cdbDataArgs struct {
	// The name of the Custom DB into which creating the data.
	CustomDb string `pulumi:"customDb"`
	// The value 1 (key of the data)
	Value1 string `pulumi:"value1"`
	// The value 10
	Value10 *string `pulumi:"value10"`
	// The value 2
	Value2 *string `pulumi:"value2"`
	// The value 3
	Value3 *string `pulumi:"value3"`
	// The value 4
	Value4 *string `pulumi:"value4"`
	// The value 5
	Value5 *string `pulumi:"value5"`
	// The value 6
	Value6 *string `pulumi:"value6"`
	// The value 7
	Value7 *string `pulumi:"value7"`
	// The value 8
	Value8 *string `pulumi:"value8"`
	// The value 9
	Value9 *string `pulumi:"value9"`
}

// The set of arguments for constructing a CdbData resource.
type CdbDataArgs struct {
	// The name of the Custom DB into which creating the data.
	CustomDb pulumi.StringInput
	// The value 1 (key of the data)
	Value1 pulumi.StringInput
	// The value 10
	Value10 pulumi.StringPtrInput
	// The value 2
	Value2 pulumi.StringPtrInput
	// The value 3
	Value3 pulumi.StringPtrInput
	// The value 4
	Value4 pulumi.StringPtrInput
	// The value 5
	Value5 pulumi.StringPtrInput
	// The value 6
	Value6 pulumi.StringPtrInput
	// The value 7
	Value7 pulumi.StringPtrInput
	// The value 8
	Value8 pulumi.StringPtrInput
	// The value 9
	Value9 pulumi.StringPtrInput
}

func (CdbDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cdbDataArgs)(nil)).Elem()
}

type CdbDataInput interface {
	pulumi.Input

	ToCdbDataOutput() CdbDataOutput
	ToCdbDataOutputWithContext(ctx context.Context) CdbDataOutput
}

func (*CdbData) ElementType() reflect.Type {
	return reflect.TypeOf((**CdbData)(nil)).Elem()
}

func (i *CdbData) ToCdbDataOutput() CdbDataOutput {
	return i.ToCdbDataOutputWithContext(context.Background())
}

func (i *CdbData) ToCdbDataOutputWithContext(ctx context.Context) CdbDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CdbDataOutput)
}

// CdbDataArrayInput is an input type that accepts CdbDataArray and CdbDataArrayOutput values.
// You can construct a concrete instance of `CdbDataArrayInput` via:
//
//	CdbDataArray{ CdbDataArgs{...} }
type CdbDataArrayInput interface {
	pulumi.Input

	ToCdbDataArrayOutput() CdbDataArrayOutput
	ToCdbDataArrayOutputWithContext(context.Context) CdbDataArrayOutput
}

type CdbDataArray []CdbDataInput

func (CdbDataArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CdbData)(nil)).Elem()
}

func (i CdbDataArray) ToCdbDataArrayOutput() CdbDataArrayOutput {
	return i.ToCdbDataArrayOutputWithContext(context.Background())
}

func (i CdbDataArray) ToCdbDataArrayOutputWithContext(ctx context.Context) CdbDataArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CdbDataArrayOutput)
}

// CdbDataMapInput is an input type that accepts CdbDataMap and CdbDataMapOutput values.
// You can construct a concrete instance of `CdbDataMapInput` via:
//
//	CdbDataMap{ "key": CdbDataArgs{...} }
type CdbDataMapInput interface {
	pulumi.Input

	ToCdbDataMapOutput() CdbDataMapOutput
	ToCdbDataMapOutputWithContext(context.Context) CdbDataMapOutput
}

type CdbDataMap map[string]CdbDataInput

func (CdbDataMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CdbData)(nil)).Elem()
}

func (i CdbDataMap) ToCdbDataMapOutput() CdbDataMapOutput {
	return i.ToCdbDataMapOutputWithContext(context.Background())
}

func (i CdbDataMap) ToCdbDataMapOutputWithContext(ctx context.Context) CdbDataMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CdbDataMapOutput)
}

type CdbDataOutput struct{ *pulumi.OutputState }

func (CdbDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CdbData)(nil)).Elem()
}

func (o CdbDataOutput) ToCdbDataOutput() CdbDataOutput {
	return o
}

func (o CdbDataOutput) ToCdbDataOutputWithContext(ctx context.Context) CdbDataOutput {
	return o
}

// The name of the Custom DB into which creating the data.
func (o CdbDataOutput) CustomDb() pulumi.StringOutput {
	return o.ApplyT(func(v *CdbData) pulumi.StringOutput { return v.CustomDb }).(pulumi.StringOutput)
}

// The value 1 (key of the data)
func (o CdbDataOutput) Value1() pulumi.StringOutput {
	return o.ApplyT(func(v *CdbData) pulumi.StringOutput { return v.Value1 }).(pulumi.StringOutput)
}

// The value 10
func (o CdbDataOutput) Value10() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CdbData) pulumi.StringPtrOutput { return v.Value10 }).(pulumi.StringPtrOutput)
}

// The value 2
func (o CdbDataOutput) Value2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CdbData) pulumi.StringPtrOutput { return v.Value2 }).(pulumi.StringPtrOutput)
}

// The value 3
func (o CdbDataOutput) Value3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CdbData) pulumi.StringPtrOutput { return v.Value3 }).(pulumi.StringPtrOutput)
}

// The value 4
func (o CdbDataOutput) Value4() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CdbData) pulumi.StringPtrOutput { return v.Value4 }).(pulumi.StringPtrOutput)
}

// The value 5
func (o CdbDataOutput) Value5() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CdbData) pulumi.StringPtrOutput { return v.Value5 }).(pulumi.StringPtrOutput)
}

// The value 6
func (o CdbDataOutput) Value6() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CdbData) pulumi.StringPtrOutput { return v.Value6 }).(pulumi.StringPtrOutput)
}

// The value 7
func (o CdbDataOutput) Value7() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CdbData) pulumi.StringPtrOutput { return v.Value7 }).(pulumi.StringPtrOutput)
}

// The value 8
func (o CdbDataOutput) Value8() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CdbData) pulumi.StringPtrOutput { return v.Value8 }).(pulumi.StringPtrOutput)
}

// The value 9
func (o CdbDataOutput) Value9() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CdbData) pulumi.StringPtrOutput { return v.Value9 }).(pulumi.StringPtrOutput)
}

type CdbDataArrayOutput struct{ *pulumi.OutputState }

func (CdbDataArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CdbData)(nil)).Elem()
}

func (o CdbDataArrayOutput) ToCdbDataArrayOutput() CdbDataArrayOutput {
	return o
}

func (o CdbDataArrayOutput) ToCdbDataArrayOutputWithContext(ctx context.Context) CdbDataArrayOutput {
	return o
}

func (o CdbDataArrayOutput) Index(i pulumi.IntInput) CdbDataOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CdbData {
		return vs[0].([]*CdbData)[vs[1].(int)]
	}).(CdbDataOutput)
}

type CdbDataMapOutput struct{ *pulumi.OutputState }

func (CdbDataMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CdbData)(nil)).Elem()
}

func (o CdbDataMapOutput) ToCdbDataMapOutput() CdbDataMapOutput {
	return o
}

func (o CdbDataMapOutput) ToCdbDataMapOutputWithContext(ctx context.Context) CdbDataMapOutput {
	return o
}

func (o CdbDataMapOutput) MapIndex(k pulumi.StringInput) CdbDataOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CdbData {
		return vs[0].(map[string]*CdbData)[vs[1].(string)]
	}).(CdbDataOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CdbDataInput)(nil)).Elem(), &CdbData{})
	pulumi.RegisterInputType(reflect.TypeOf((*CdbDataArrayInput)(nil)).Elem(), CdbDataArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CdbDataMapInput)(nil)).Elem(), CdbDataMap{})
	pulumi.RegisterOutputType(CdbDataOutput{})
	pulumi.RegisterOutputType(CdbDataArrayOutput{})
	pulumi.RegisterOutputType(CdbDataMapOutput{})
}
