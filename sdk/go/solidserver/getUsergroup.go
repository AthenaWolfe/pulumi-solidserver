// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package solidserver

import (
	"context"
	"reflect"

	"github.com/AthenaWolfe/pulumi-solidserver/sdk/go/solidserver/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// User group data-source allows to retrieve information about groups used to manage
// permissions within SOLIDserver.
func LookupUsergroup(ctx *pulumi.Context, args *LookupUsergroupArgs, opts ...pulumi.InvokeOption) (*LookupUsergroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUsergroupResult
	err := ctx.Invoke("solidserver:index/getUsergroup:getUsergroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUsergroup.
type LookupUsergroupArgs struct {
	// The name of the user group.
	Name string `pulumi:"name"`
}

// A collection of values returned by getUsergroup.
type LookupUsergroupResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the user group.
	Name string `pulumi:"name"`
}

func LookupUsergroupOutput(ctx *pulumi.Context, args LookupUsergroupOutputArgs, opts ...pulumi.InvokeOption) LookupUsergroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUsergroupResult, error) {
			args := v.(LookupUsergroupArgs)
			r, err := LookupUsergroup(ctx, &args, opts...)
			var s LookupUsergroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUsergroupResultOutput)
}

// A collection of arguments for invoking getUsergroup.
type LookupUsergroupOutputArgs struct {
	// The name of the user group.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupUsergroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUsergroupArgs)(nil)).Elem()
}

// A collection of values returned by getUsergroup.
type LookupUsergroupResultOutput struct{ *pulumi.OutputState }

func (LookupUsergroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUsergroupResult)(nil)).Elem()
}

func (o LookupUsergroupResultOutput) ToLookupUsergroupResultOutput() LookupUsergroupResultOutput {
	return o
}

func (o LookupUsergroupResultOutput) ToLookupUsergroupResultOutputWithContext(ctx context.Context) LookupUsergroupResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupUsergroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUsergroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the user group.
func (o LookupUsergroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUsergroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUsergroupResultOutput{})
}