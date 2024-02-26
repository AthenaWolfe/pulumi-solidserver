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

// The provider type for the solidserver package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// PEM formatted file with additional certificates to trust for TLS connection
	AdditionalTrustCertsFile pulumi.StringPtrOutput `pulumi:"additionalTrustCertsFile"`
	// SOLIDServer Hostname or IP address
	Host pulumi.StringOutput `pulumi:"host"`
	// SOLIDServer API user's password
	Password pulumi.StringOutput `pulumi:"password"`
	// SOLIDServer Version in case API user does not have admin permissions
	Solidserverversion pulumi.StringPtrOutput `pulumi:"solidserverversion"`
	// SOLIDServer API user's ID
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Host == nil {
		return nil, errors.New("invalid value for required argument 'Host'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:solidserver", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// PEM formatted file with additional certificates to trust for TLS connection
	AdditionalTrustCertsFile *string `pulumi:"additionalTrustCertsFile"`
	// SOLIDServer Hostname or IP address
	Host string `pulumi:"host"`
	// SOLIDServer API user's password
	Password string `pulumi:"password"`
	// SOLIDServer Version in case API user does not have admin permissions
	Solidserverversion *string `pulumi:"solidserverversion"`
	// Enable/Disable ssl verify (Default : enabled)
	Sslverify *bool `pulumi:"sslverify"`
	// SOLIDServer API user's ID
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// PEM formatted file with additional certificates to trust for TLS connection
	AdditionalTrustCertsFile pulumi.StringPtrInput
	// SOLIDServer Hostname or IP address
	Host pulumi.StringInput
	// SOLIDServer API user's password
	Password pulumi.StringInput
	// SOLIDServer Version in case API user does not have admin permissions
	Solidserverversion pulumi.StringPtrInput
	// Enable/Disable ssl verify (Default : enabled)
	Sslverify pulumi.BoolPtrInput
	// SOLIDServer API user's ID
	Username pulumi.StringInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// PEM formatted file with additional certificates to trust for TLS connection
func (o ProviderOutput) AdditionalTrustCertsFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.AdditionalTrustCertsFile }).(pulumi.StringPtrOutput)
}

// SOLIDServer Hostname or IP address
func (o ProviderOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// SOLIDServer API user's password
func (o ProviderOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// SOLIDServer Version in case API user does not have admin permissions
func (o ProviderOutput) Solidserverversion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Solidserverversion }).(pulumi.StringPtrOutput)
}

// SOLIDServer API user's ID
func (o ProviderOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}